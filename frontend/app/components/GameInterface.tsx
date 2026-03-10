'use client';

import { useEffect, useState, useRef } from 'react';
import { useRouter } from 'next/navigation';
import SuccessPopup from './SuccessPopup';
import AppHeader from './AppHeader';
import StreakPopup from './StreakPopup';
import ChallengeOverPopup from './ChallengeOverPopup';
import HowToPlayPopup from './HowToPlayPopup';
import ImageModal from './ImageModal';
import BannerAd from './BannerAd';
import Link from 'next/link';
import { useAuth } from './AuthProvider';

interface Make {
    id: number;
    name: string;
}

interface Model {
    id: number[];
    name: string;
}

interface UserStatus {
    attempts: number;
    max_attempts: number;
    is_completed: boolean;
    is_correct: boolean;
}

interface ChallengeStats {
    players_today: number;
    average_accuracy: number;
    total_bonus_points: number;
}


interface StreakStats {
    attendance_streak: number;
    submission_streak: number;
    max_attendance_streak: number;
    max_submission_streak: number;
    total_days_participated: number;
    total_days_submitted: number;
}


interface Challenge {
    id: number;
    date: string;
    image_url: string;
    user_status?: UserStatus;
    next_challenge_seconds?: number;
    streak_stats?: StreakStats;

    make?: Make;
    model?: {
        id: number;
        name: string;
        year_range?: string;
        generation?: string;
        location?: string;
        codename?: string;
        image_url?: string;
        known_for?: string;
    };
    points_earned?: number;
    bonus_round?: BonusRoundInfo;
    attempt_history?: {
        is_make_correct: boolean;
        is_model_correct: boolean;
        attempt_number: number;
    }[];
}

interface BonusRoundInfo {
    year_range_enabled: boolean;
    generation_enabled: boolean;
    codename_enabled: boolean;
    year_range_attempted: boolean;
    generation_attempted: boolean;
    codename_attempted: boolean;
    year_range_points: number;
    generation_points: number;
    codename_points: number;
    year_range_correct?: boolean;
    generation_correct?: boolean;
    codename_correct?: boolean;
}

interface SubmissionResult {
    correct: boolean;
    message: string;
    image_url: string;
    solution?: {
        make_name: string;
        model_name: string;
        year_range?: string;
        generation?: string;
        codename?: string;
        known_for?: string;
        image_url?: string;
    };
    user_status?: UserStatus;
    bonus_round?: BonusRoundInfo;
    points_earned: number;
    next_challenge_seconds?: number;
    streak_stats?: StreakStats;
    isHistorical?: boolean;
    attempt_history?: {
        is_make_correct: boolean;
        is_model_correct: boolean;
        attempt_number: number;
    }[];
}

export default function GameInterface() {
    const router = useRouter();
    const [challenge, setChallenge] = useState<Challenge | null>(null);
    const [makes, setMakes] = useState<Make[]>([]);
    const [models, setModels] = useState<Model[]>([]);
    const [selectedMake, setSelectedMake] = useState<string>('');
    const [selectedModel, setSelectedModel] = useState<string>('');
    const [userStatus, setUserStatus] = useState<UserStatus | null>(null);
    const [showSuccessPopup, setShowSuccessPopup] = useState(false);
    const [showStreakPopup, setShowStreakPopup] = useState(false);
    const [showChallengeOverPopup, setShowChallengeOverPopup] = useState(false);
    const [showHowToPlayPopup, setShowHowToPlayPopup] = useState(false);
    const [result, setResult] = useState<SubmissionResult | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [showImageModal, setShowImageModal] = useState(false);

    const [secondsUntilReset, setSecondsUntilReset] = useState<number | null>(null);
    const [stats, setStats] = useState<ChallengeStats | null>(null);

    const { managedFetch, isLoading: authLoading, sessionToken } = useAuth();

    useEffect(() => {
        if (challenge?.next_challenge_seconds) {
            setSecondsUntilReset(challenge.next_challenge_seconds);
        }
    }, [challenge]);

    useEffect(() => {
        if (secondsUntilReset === null) return;
        const timer = setInterval(() => {
            setSecondsUntilReset(prev => (prev && prev > 0 ? prev - 1 : 0));
        }, 1000);
        return () => clearInterval(timer);
    }, [secondsUntilReset]);

    const formatCountdown = (totalSeconds: number | null) => {
        if (totalSeconds === null) return '--:--:--';
        const hours = Math.floor(totalSeconds / 3600);
        const minutes = Math.floor((totalSeconds % 3600) / 60);
        const seconds = totalSeconds % 60;
        return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    };

    useEffect(() => {
        const checkFirstVisit = () => {
            const hasSeen = localStorage.getItem('autocorrect_has_seen_how_to_play');
            if (!hasSeen) {
                setShowHowToPlayPopup(true);
            }
        };

        if (!authLoading) {
            fetchChallenge();
            fetchMakes();
            checkFirstVisit();
        }
    }, [authLoading]);

    useEffect(() => {
        if (selectedMake) {
            fetchModels(selectedMake);
        } else {
            setModels([]);
        }
    }, [selectedMake]);

    const fetchChallenge = async () => {
        try {
            const response = await managedFetch('/api/challenge/today');
            const data: Challenge = await response.json();
            setChallenge(data);
            if (data.id) {
                fetchChallengeStats(data.id);
            }
            if (data.user_status) {
                setUserStatus(data.user_status);
                if (data.user_status.is_completed) {
                    // Challenge was already completed in a previous session
                    setResult({
                        correct: data.user_status.is_correct,
                        message: data.user_status.is_correct ? 'Challenge already completed!' : 'Max attempts reached for today.',
                        image_url: data.image_url,
                        user_status: data.user_status,
                        solution: data.model ? {
                            make_name: data.make?.name || '',
                            model_name: data.model.name,
                            year_range: data.model.year_range,
                            generation: data.model.generation,
                            codename: data.model.codename,
                            image_url: data.model.image_url,
                            known_for: data.model.known_for
                        } : undefined,
                        points_earned: data.points_earned || 0,
                        next_challenge_seconds: data.next_challenge_seconds,
                        bonus_round: data.bonus_round,
                        isHistorical: true
                    });
                }
            }
        } catch (error) {
            console.error('Failed to fetch challenge:', error);
            setError('Failed to load today\'s challenge. Please try again.');
        } finally {
            setLoading(false);
        }
    };

    const fetchChallengeStats = async (challengeId: number) => {
        try {
            const response = await managedFetch(`/api/challenge/stats?challenge_id=${challengeId}`);
            const data = await response.json();
            setStats(data);
        } catch (error) {
            console.error('Failed to fetch challenge stats:', error);
        }
    };

    const fetchMakes = async () => {
        try {
            const response = await managedFetch('/api/makes');
            const data = await response.json();
            setMakes(data);
        } catch (error) {
            console.error('Failed to fetch makes:', error);
        }
    };

    const fetchModels = async (makeId: string) => {
        try {
            const response = await managedFetch(`/api/models?make_id=${makeId}`);
            const data = await response.json();
            setModels(data);
        } catch (error) {
            console.error('Failed to fetch models:', error);
        }
    };

    const isChallengeOver = userStatus?.is_completed || result?.correct;

    const handleSubmit = async () => {
        if (!challenge) return;

        // If challenge is already completed, show the appropriate popup
        if (isChallengeOver) {
            if (userStatus?.is_correct || result?.correct) {
                setShowSuccessPopup(true);
            } else {
                setShowChallengeOverPopup(true);
            }
            return;
        }

        if (!selectedMake || !selectedModel) return;

        try {
            const response = await managedFetch('/api/challenge/submit', {
                method: 'POST',
                body: JSON.stringify({
                    challenge_id: challenge.id,
                    make_id: parseInt(selectedMake),
                    model_ids: selectedModel.split(',').map(id => parseInt(id, 10)),
                }),
            });

            // Handle 403 - challenge already completed or max attempts reached
            if (!response.ok) {
                const data = await response.json();
                if (response.status === 403) {
                    setUserStatus(prev => prev ? { ...prev, is_completed: true } : { attempts: 0, max_attempts: 3, is_completed: true, is_correct: false });
                    setShowChallengeOverPopup(true);
                } else {
                    setError(data.error || 'Submission failed');
                }
                return;
            }

            const data: SubmissionResult = await response.json();
            setResult(data);
            console.table(data);
            console.log(`submission data. image = ${data.image_url}`)
            if (data.user_status) {
                setUserStatus(data.user_status);
            }
            // Always refresh stats after an attempt
            fetchChallengeStats(challenge.id);

            if (data.correct) {
                setShowSuccessPopup(true);
            }

            // If the submission result indicates challenge is now over (max attempts used up with wrong answer),
            // show the challenge over popup
            if (data.user_status?.is_completed && !data.correct) {
                setShowChallengeOverPopup(true);
            }
        } catch (error) {
            console.error('Failed to submit guess:', error);
            setError('A network error occurred. Please check your connection.');
        }
    };

    const openLeaderBoard = () => {
        router.push('/leaderboard');
    };

    if (loading) {
        return (
            <div className="flex items-center justify-center min-h-[60vh]">
                <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
            </div>
        );
    }

    return (
        <div className="animate-enter">
            {/* Error Notification */}
            {error && (
                <div className="fixed top-24 left-1/2 -translate-x-1/2 z-[100] w-full max-w-md px-4">
                    <div className="bg-red-500/90 backdrop-blur-md text-white p-4 rounded-xl shadow-2xl flex items-center justify-between border border-white/20">
                        <div className="flex items-center gap-3">
                            <span className="material-symbols-outlined">error</span>
                            <p className="font-bold text-sm tracking-tight">{error}</p>
                        </div>
                        <button onClick={() => setError(null)} className="p-1 hover:bg-white/20 rounded-full transition-colors">
                            <span className="material-symbols-outlined text-lg">close</span>
                        </button>
                    </div>
                </div>
            )}
            {result?.correct && result.solution && showSuccessPopup && (
                <SuccessPopup
                    isOpen={showSuccessPopup}
                    onClose={() => setShowSuccessPopup(false)}
                    challengeId={challenge?.id || 0}
                    data={{
                        make_name: result.solution.make_name,
                        model_name: result.solution.model_name,
                        year_range: result.solution.year_range,
                        generation: result.solution.generation,
                        codename: result.solution.codename,
                        image_url: result.solution.image_url || result.image_url,
                        attempts: userStatus?.attempts || 0,
                        points_earned: result.points_earned,
                        bonus_round: result.bonus_round,
                        attempt_history: result.attempt_history || challenge?.attempt_history,
                        model_known_for: result.solution.known_for,
                        next_challenge_seconds: result.next_challenge_seconds || challenge?.next_challenge_seconds,
                    }}
                    sessionToken={sessionToken || ''}
                    managedFetch={managedFetch}
                    onUpdatePoints={() => challenge && fetchChallengeStats(challenge.id)}
                />
            )}
            {/* Header */}
            <AppHeader
                streakCount={result?.streak_stats?.submission_streak || challenge?.streak_stats?.submission_streak || 0}
                onStreakClick={() => setShowStreakPopup(true)}
                onHelpClick={() => setShowHowToPlayPopup(true)}
            />

            {/* Top Banner Ad — blends naturally below the header */}
            <div className="max-w-[960px] mx-auto w-full px-4 md:px-0 pt-4">
                <BannerAd adSlot={process.env.NEXT_PUBLIC_ADSENSE_SLOT_HOME_TOP || ''} label="" />
            </div>

            {/* Challenge Over Popup */}
            <ChallengeOverPopup
                isOpen={showChallengeOverPopup}
                onClose={() => setShowChallengeOverPopup(false)}
                isCorrect={userStatus?.is_correct || false}
                nextChallengeSeconds={result?.next_challenge_seconds || challenge?.next_challenge_seconds || null}
                solution={result?.solution ? { make_name: result.solution.make_name, model_name: result.solution.model_name } : (challenge?.make && challenge?.model ? { make_name: challenge.make.name, model_name: challenge.model.name } : undefined)}
            />

            {/* Streak Popup */}
            <StreakPopup
                isOpen={showStreakPopup}
                onClose={() => setShowStreakPopup(false)}
                stats={result?.streak_stats || challenge?.streak_stats}
            />

            {/* How to Play Popup */}
            <HowToPlayPopup
                isOpen={showHowToPlayPopup}
                onClose={() => {
                    localStorage.setItem('autocorrect_has_seen_how_to_play', 'true');
                    setShowHowToPlayPopup(false);
                }}
            />

            <ImageModal
                isOpen={showImageModal}
                onClose={() => setShowImageModal(false)}
                imageUrl={challenge?.image_url || ''}
            />

            <main className="max-w-[960px] mx-auto py-8">
                {/* Today's Challenge Header */}
                <div className="text-center mb-8 relative flex flex-col items-center justify-center">
                    <div className="relative w-full flex justify-center items-center">
                        <h1 className="text-4xl font-bold tracking-tighter uppercase mb-2">
                            Challenge #{challenge?.id || '---'}
                        </h1>
                    </div>
                    <p className="text-slate-500 dark:text-slate-400 font-medium">Identify the vehicle from the detail shown</p>
                </div>

                {/* Main Challenge Area */}
                <div className="space-y-6">
                    {/* Hero Card: The Lamp */}
                    <div className="relative group cursor-zoom-in" onClick={() => setShowImageModal(true)}>
                        <div className="absolute -inset-1 bg-gradient-to-r from-primary to-accent-purple rounded-xl blur opacity-20 group-hover:opacity-30 transition duration-1000"></div>
                        <div className="relative bg-white dark:bg-[#192233] rounded-xl overflow-hidden border border-white/5 shadow-2xl">
                            {/* High Quality Car Image */}
                            <div className="w-full aspect-[21/9] relative overflow-hidden">
                                <img
                                    src={challenge?.image_url}
                                    alt="Daily Car Challenge"
                                    className="w-full h-full object-cover object-center"
                                />
                                {/* Overlay Blur for Depth */}
                                <div className="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
                                {/* Enlarge Button Overlay */}
                                <div className="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                                    <button
                                        onClick={(e) => { e.stopPropagation(); setShowImageModal(true); }}
                                        className="bg-primary hover:bg-primary/90 text-white px-6 py-3 rounded-full flex items-center gap-2 font-bold shadow-xl transform scale-95 group-hover:scale-100 transition-transform"
                                    >
                                        <span className="material-symbols-outlined">zoom_in</span>
                                        ENLARGE DETAIL
                                    </button>
                                </div>
                                {/* Hints/UI Elements inside image */}
                                <div className="absolute bottom-4 left-4 flex gap-2">

                                    <button
                                        onClick={(e) => e.stopPropagation()}
                                        className="bg-black/40 backdrop-blur-md border border-white/20 p-2 rounded-lg text-white hover:bg-black/60 transition-colors"
                                    >
                                        <span className="material-symbols-outlined text-lg">share</span>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>

                    {/* Result Message */}
                    {result && (
                        <div className={`p-4 rounded-lg flex items-center gap-3 border ${result.correct ? 'bg-green-500/10 border-green-500/50 text-green-500' : 'bg-red-500/10 border-red-500/50 text-red-500'}`}>
                            <span className="material-symbols-outlined">
                                {result.correct ? 'check_circle' : 'cancel'}
                            </span>
                            <p className="font-bold">{result.message}</p>
                            {result.correct && (
                                <button
                                    onClick={() => setShowSuccessPopup(true)}
                                    className="ml-auto bg-green-500/20 hover:bg-green-500/30 text-green-500 px-3 py-1 rounded-md text-sm font-bold border border-green-500/30 transition-colors"
                                >
                                    VIEW RESULTS
                                </button>
                            )}
                        </div>
                    )}

                    {/* Input Controls: The "Garage" */}
                    <div className="glass-panel border border-white/10 rounded-xl p-8 space-y-8 techno-glow">
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                            {/* Make Selection */}
                            <div className="space-y-2">
                                <label className="text-xs font-bold uppercase tracking-widest text-slate-400 block ml-1">Manufacturer</label>
                                <div className="relative">
                                    <select
                                        value={selectedMake}
                                        onChange={(e) => {
                                            setSelectedMake(e.target.value);
                                            if (result && !result.correct) {
                                                setResult(null);
                                            }
                                        }}
                                        className="w-full h-14 bg-background-dark border border-white/10 rounded-lg px-4 text-white appearance-none focus:border-primary focus:ring-1 focus:ring-primary outline-none cursor-pointer transition-all"
                                    >
                                        <option value="" disabled>Select Make...</option>
                                        {makes.map((make) => (
                                            <option key={make.id} value={make.id}>{make.name}</option>
                                        ))}
                                    </select>
                                    <div className="absolute right-4 top-1/2 -translate-y-1/2 pointer-events-none text-slate-500">
                                        <span className="material-symbols-outlined">expand_more</span>
                                    </div>
                                </div>
                            </div>

                            {/* Model Selection */}
                            <div className="space-y-2">
                                <label className="text-xs font-bold uppercase tracking-widest text-slate-400 block ml-1">Model Name</label>
                                <div className="relative">
                                    <select
                                        value={selectedModel}
                                        onChange={(e) => {
                                            setSelectedModel(e.target.value);
                                            if (result && !result.correct) {
                                                setResult(null);
                                            }
                                        }}
                                        disabled={!selectedMake}
                                        className="w-full h-14 bg-background-dark border border-white/10 rounded-lg px-4 text-white appearance-none focus:border-primary focus:ring-1 focus:ring-primary outline-none cursor-pointer transition-all disabled:opacity-50 disabled:cursor-not-allowed"
                                    >
                                        <option value="" disabled>Select Model...</option>
                                        {models.map((model) => (
                                            <option key={model.name} value={model.id.join(',')}>{model.name}</option>
                                        ))}
                                    </select>
                                    <div className="absolute right-4 top-1/2 -translate-y-1/2 pointer-events-none text-slate-500">
                                        <span className="material-symbols-outlined">expand_more</span>
                                    </div>
                                </div>
                            </div>
                        </div>



                        {/* Submit Button */}
                        <div className="pt-4">
                            <button
                                onClick={handleSubmit}
                                disabled={!isChallengeOver && (!selectedMake || !selectedModel)}
                                className={`w-full py-5 rounded-lg font-black text-xl tracking-widest uppercase transition-all flex items-center justify-center gap-3 group relative overflow-hidden ${isChallengeOver
                                    ? (userStatus?.is_correct
                                        ? 'bg-green-500/20 border border-green-500/30 text-green-400 hover:bg-green-500/30 cursor-pointer'
                                        : 'bg-amber-500/20 border border-amber-500/30 text-amber-400 hover:bg-amber-500/30 cursor-pointer')
                                    : 'bg-primary hover:bg-primary/90 disabled:bg-slate-700 disabled:cursor-not-allowed text-white'
                                    }`}
                            >
                                <div className="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-700"></div>
                                {userStatus?.is_correct ? 'CHALLENGE COMPLETED' : (userStatus?.is_completed ? 'MAX ATTEMPTS REACHED' : 'SUBMIT GUESS')}
                                <span className="material-symbols-outlined font-normal">send</span>
                            </button>
                            <div className="mt-4 flex justify-center items-center gap-4 text-xs font-medium text-slate-500 uppercase tracking-widest">
                                <span>Attempts: {userStatus?.attempts || 0}/{userStatus?.max_attempts || 3}</span>
                                <span className="w-1 h-1 bg-slate-700 rounded-full"></span>
                                <span>Global Accuracy: {stats ? `${Math.round(stats.average_accuracy)}%` : '--%'}</span>
                            </div>
                        </div>
                    </div>
                </div>

                {/* Pre-Footer Banner Ad */}
                <div className="mt-12">
                    <BannerAd adSlot={process.env.NEXT_PUBLIC_ADSENSE_SLOT_HOME_BOTTOM || ''} label="Advertisement" />
                </div>

                {/* Footer Stats */}
                <footer className="mt-8 text-center space-y-6">
                    <div className="flex flex-wrap justify-center gap-8">
                        <div className="flex flex-col items-center">
                            <span className="text-2xl font-bold text-primary">
                                {stats?.players_today.toLocaleString() || '---'}
                            </span>
                            <span className="text-[10px] uppercase tracking-widest text-slate-500">Players Today</span>
                        </div>
                        <div className="flex flex-col items-center">
                            <span className="text-2xl font-bold text-accent-cyan">
                                {stats ? `${Math.round(stats.average_accuracy)}%` : '--%'}
                            </span>
                            <span className="text-[10px] uppercase tracking-widest text-slate-500">Avg. Accuracy</span>
                        </div>
                        <div className="flex flex-col items-center">
                            <span className="text-2xl font-bold text-accent-purple">
                                {stats?.total_bonus_points.toLocaleString() || '---'}
                            </span>
                            <span className="text-[10px] uppercase tracking-widest text-slate-500">Total Bonus Points</span>
                        </div>
                    </div>
                    <div className="pt-16 pb-12 text-slate-600 text-[10px] uppercase tracking-[0.3em] font-black flex flex-col items-center gap-6 border-t border-white/5">
                        <span>Daily Reset in {formatCountdown(secondsUntilReset)}</span>
                        <div className="flex gap-8">
                            <Link href="/privacy" className="hover:text-primary transition-colors">Privacy Policy</Link>
                            <Link href="/terms" className="hover:text-primary transition-colors">Terms and Conditions</Link>
                        </div>
                        <span className="text-slate-700 tracking-normal mt-2 opacity-50">
                            {process.env.NEXT_PUBLIC_APP_VERSION || 'v1.0.0-LOCAL'}
                        </span>
                    </div>
                </footer>
            </main>
        </div>
    );
}

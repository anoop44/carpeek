'use client';

import { useEffect, useState, useRef } from 'react';
import Link from 'next/link';
import BonusInputPopup from './BonusInputPopup';
import StreakMeter from './StreakMeter';

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

interface StreakStats {
    attendance_streak: number;
    submission_streak: number;
    max_attendance_streak: number;
    max_submission_streak: number;
    total_days_participated: number;
    total_days_submitted: number;
}


interface SuccessPopupProps {
    isOpen: boolean;
    onClose: () => void;
    challengeId: number;
    data: {
        make_name: string;
        model_name: string;
        year_range?: string;
        generation?: string;
        codename?: string;
        image_url: string;
        attempts: number;
        points_earned?: number;
        bonus_round?: BonusRoundInfo;
        attempt_history?: {
            is_make_correct: boolean;
            is_model_correct: boolean;
            attempt_number: number;
        }[];
        model_known_for?: string;
        next_challenge_seconds?: number;
        streak_stats?: StreakStats;
    };
    onUpdatePoints?: () => void;
    sessionToken?: string;
    managedFetch?: (url: string, options?: RequestInit) => Promise<Response>;
}

export default function SuccessPopup({ isOpen, onClose, challengeId, data, onUpdatePoints, sessionToken, managedFetch }: SuccessPopupProps) {
    const [secondsRemaining, setSecondsRemaining] = useState<number | null>(data.next_challenge_seconds || null);
    const [activeBonus, setActiveBonus] = useState<{ type: string; label: string; points: number } | null>(null);
    const [bonusState, setBonusState] = useState<BonusRoundInfo | undefined>(data.bonus_round);
    const [submitting, setSubmitting] = useState(false);
    const [feedback, setFeedback] = useState<{ message: string; isCorrect: boolean } | null>(null);
    const [showScrollIndicator, setShowScrollIndicator] = useState(false);
    const scrollContainerRef = useRef<HTMLDivElement>(null);

    const checkScrollable = () => {
        if (scrollContainerRef.current) {
            const { scrollHeight, clientHeight, scrollTop } = scrollContainerRef.current;
            setShowScrollIndicator(scrollHeight > clientHeight && scrollHeight - scrollTop - clientHeight > 15);
        }
    };

    useEffect(() => {
        if (isOpen) {
            const timeout = setTimeout(checkScrollable, 100);
            window.addEventListener('resize', checkScrollable);
            return () => {
                clearTimeout(timeout);
                window.removeEventListener('resize', checkScrollable);
            };
        }
    }, [data, bonusState, isOpen]);

    const handleScroll = (e: React.UIEvent<HTMLDivElement>) => {
        const { scrollTop, scrollHeight, clientHeight } = e.currentTarget;
        if (scrollHeight - scrollTop - clientHeight < 15) {
            setShowScrollIndicator(false);
        } else {
            setShowScrollIndicator(true);
        }
    };

    // Update local state when data changes
    useEffect(() => {
        setBonusState(data.bonus_round);
        if (data.next_challenge_seconds !== undefined) {
            setSecondsRemaining(data.next_challenge_seconds);
        }
    }, [data.bonus_round, data.next_challenge_seconds]);

    useEffect(() => {
        if (secondsRemaining === null) return;

        const timer = setInterval(() => {
            setSecondsRemaining(prev => (prev !== null && prev > 0 ? prev - 1 : 0));
        }, 1000);

        return () => clearInterval(timer);
    }, [secondsRemaining]);

    useEffect(() => {
        if (!isOpen) return;
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === 'Escape') onClose();
        };
        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [isOpen, onClose]);

    const formatTime = (totalSeconds: number | null) => {
        if (totalSeconds === null) return '--:--:--';
        const hours = Math.floor(totalSeconds / 3600);
        const minutes = Math.floor((totalSeconds % 3600) / 60);
        const seconds = totalSeconds % 60;
        return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    };

    const handleBonusSubmit = async (value: string): Promise<{ message: string; isCorrect: boolean }> => {
        if (!activeBonus || !value || submitting) throw new Error('Invalid state');
        setSubmitting(true);
        const { type } = activeBonus;

        try {
            const fetcher = managedFetch || fetch;
            const response = await fetcher('/api/challenge/bonus/submit', {
                method: 'POST',
                body: JSON.stringify({
                    challenge_id: challengeId,
                    bonus_type: type,
                    value: value
                }),
            });

            const result = await response.json();
            if (!response.ok) {
                setFeedback({ message: result.error || 'Bonus submission failed', isCorrect: false });
                setTimeout(() => setFeedback(null), 3000);
                return { message: result.error || 'Bonus submission failed', isCorrect: false };
            }

            // Show feedback briefly or handle state update
            if (bonusState) {
                const newState = { ...bonusState };
                if (type === 'year_range') {
                    newState.year_range_attempted = true;
                    newState.year_range_correct = result.correct;
                }
                if (type === 'generation') {
                    newState.generation_attempted = true;
                    newState.generation_correct = result.correct;
                }
                if (type === 'codename') {
                    newState.codename_attempted = true;
                    newState.codename_correct = result.correct;
                }
                setBonusState(newState);
            }
            if (onUpdatePoints) onUpdatePoints();

            // Build feedback message that always shows the correct value
            let feedbackMsg = result.message;
            if (result.correct_value) {
                if (result.correct) {
                    feedbackMsg = `Correct! Answer: ${result.correct_value}`;
                } else {
                    feedbackMsg = `Incorrect. Answer: ${result.correct_value}`;
                }
            }

            setFeedback({ message: feedbackMsg, isCorrect: result.correct });

            // Auto clear feedback after 3s
            setTimeout(() => setFeedback(null), 3000);
            return { message: feedbackMsg, isCorrect: result.correct };
        } catch (error) {
            console.error('Bonus submission failed:', error);
            throw error;
        } finally {
            setSubmitting(false);
        }
    };

    const handleShare = async () => {
        const history = data.attempt_history || [];

        // 1. Car Attempts Grid
        // 🚗 = guess attempt, 🟩 = correct, ⬛ = wrong
        const grid = history.map(attempt => {
            const makeEmoji = attempt.is_make_correct ? '🟩' : '⬛';
            const modelEmoji = attempt.is_model_correct ? '🟩' : '⬛';
            return `🚗${makeEmoji}${modelEmoji}`;
        }).join('\n');

        // 2. Bonus Indicators
        // 🎯 = Correct, 🚫 = Wrong or Not Attempted
        let bonusGrid = '';
        let bonusWeight = 0;
        let bonusCorrectCount = 0;

        if (bonusState) {
            const types = [
                { enabled: bonusState.year_range_enabled, correct: bonusState.year_range_correct },
                { enabled: bonusState.generation_enabled, correct: bonusState.generation_correct },
                { enabled: bonusState.codename_enabled, correct: bonusState.codename_correct }
            ];

            types.forEach(type => {
                if (type.enabled) {
                    bonusWeight += 1;
                    if (type.correct) {
                        bonusGrid += '🎯';
                        bonusCorrectCount += 1;
                    } else {
                        bonusGrid += '🚫';
                    }
                }
            });
        }

        // 3. Score Calculation (Integer based for share message as per example)
        // Solve points: 1st=5, 2nd=3, 3rd=1
        let solvePoints = 0;
        const solvedAttempt = history.find(a => a.is_make_correct && a.is_model_correct);
        if (solvedAttempt) {
            if (solvedAttempt.attempt_number === 1) solvePoints = 5;
            else if (solvedAttempt.attempt_number === 2) solvePoints = 3;
            else if (solvedAttempt.attempt_number === 3) solvePoints = 1;
        }

        const totalScore = solvePoints + bonusCorrectCount;
        const maxScore = 5 + bonusWeight;

        const shareText = `AutoCorrect #${challengeId}\n\n` +
            `${grid}\n\n` +
            `${bonusGrid ? bonusGrid + '\n\n' : ''}` +
            `Score: ${totalScore}/${maxScore}\n\n` +
            `${window.location.href}`;

        try {
            if (navigator.share) {
                await navigator.share({
                    title: 'CarPeek Results',
                    text: shareText
                });
            } else {
                await navigator.clipboard.writeText(shareText);
                setFeedback({ message: 'Result copied to clipboard!', isCorrect: true });
                setTimeout(() => setFeedback(null), 3000);
            }
        } catch (err) {
            console.error('Share failed:', err);
        }
    };

    if (!isOpen) return null;

    // Helper to render a bonus row
    const renderBonusRow = (type: string, label: string, points: number, enabled: boolean, attempted: boolean, correct?: boolean) => {
        if (!enabled) return null;

        return (
            <div className="spec-row flex flex-col p-3.5 transition-colors">
                <div className="flex items-center justify-between w-full">
                    <span className="text-white/40 text-[10px] uppercase tracking-widest font-bold">{label}</span>

                    {attempted ? (
                        <div className="flex items-center gap-1.5">
                            <span className={`material-symbols-outlined text-sm ${correct ? 'text-green-500' : 'text-red-500'}`}>
                                {correct ? 'check_circle' : 'cancel'}
                            </span>
                            <span className={`text-[10px] font-bold uppercase tracking-wider ${correct ? 'text-green-500/80' : 'text-red-500/80'}`}>
                                {correct ? `Correct (+${points} PTS)` : 'Incorrect'}
                            </span>
                        </div>
                    ) : (
                        <button
                            onClick={() => setActiveBonus({ type, label, points })}
                            className="flex items-center gap-1.5 text-primary hover:text-neon-cyan transition-colors"
                        >
                            <span className="material-symbols-outlined text-sm">edit_square</span>
                            <span className="text-[10px] font-bold uppercase tracking-wider">Guess for +{points} PTS</span>
                        </button>
                    )}
                </div>
            </div>
        );
    };

    return (
        <div className="fixed inset-0 z-30 flex items-center justify-center p-4">
            {/* Dark Overlay */}
            <div className="fixed inset-0 bg-black/75 z-0" onClick={onClose}></div>

            {/* Confetti Background (Static) */}
            <div className="fixed inset-0 z-20 pointer-events-none overflow-hidden" aria-hidden="true">
                <div className="confetti-shape top-1/4 left-1/4 w-4 h-4 bg-primary rotate-45 opacity-60"></div>
                <div className="confetti-shape top-1/3 right-1/4 w-3 h-3 bg-neon-cyan rounded-full opacity-60"></div>
                <div className="confetti-shape bottom-1/4 left-1/3 w-6 h-2 bg-primary -rotate-12 opacity-60"></div>
                <div className="confetti-shape top-2/3 right-1/3 w-2 h-8 bg-neon-cyan rotate-45 opacity-60"></div>
            </div>

            {/* Glass Panel */}
            <div className="glass-panel w-full max-w-[520px] max-h-[80vh] rounded-xl overflow-hidden flex flex-col relative z-30 animate-scale-in">
                {/* Scanline */}
                <div className="h-1 shrink-0 bg-gradient-to-r from-transparent via-cyan-400 to-transparent opacity-50"></div>

                {/* Scrollable Body */}
                <div className="flex-1 relative min-h-0">
                    <div 
                        ref={scrollContainerRef}
                        className="h-full overflow-y-auto custom-scrollbar"
                        onScroll={handleScroll}
                    >
                        <div className="flex flex-col items-center text-center p-6 sm:p-8 sm:pb-4 pb-4">
                        {/* Icon */}
                        <div className="mb-4 relative">
                            <div className="absolute inset-0 bg-primary blur-xl opacity-20"></div>
                            <div className="w-14 h-14 rounded-full bg-primary/20 border border-primary/50 flex items-center justify-center relative z-10">
                                <span className="material-symbols-outlined text-cyan-400 text-3xl neon-glow">check_circle</span>
                            </div>
                        </div>

                        <h1 className="text-white tracking-[0.15em] text-2xl sm:text-3xl font-bold leading-tight uppercase mb-1">
                            Spot Confirmed
                        </h1>
                        <p className="text-cyan-400 text-[10px] font-bold tracking-[0.2em] uppercase mb-6 opacity-80">
                            Vehicle Identified Successfully
                        </p>

                        {/* Feedback Toast */}
                        {feedback && (
                            <div className={`mb-4 px-4 py-2 rounded text-xs font-bold w-full border ${feedback.isCorrect ? 'bg-green-500/10 border-green-500/50 text-green-400' : 'bg-red-500/10 border-red-500/50 text-red-400'}`}>
                                {feedback.message}
                            </div>
                        )}

                        {/* Content Card */}
                        <div className="w-full bg-[#192233]/40 rounded-xl overflow-hidden border border-white/5 mb-6">
                            {/* Image */}
                            <div
                                className="w-full h-40 bg-center bg-no-repeat bg-cover relative"
                                style={{ backgroundImage: `url(${data.image_url})` }}
                            >
                                <div className="absolute inset-0 bg-gradient-to-t from-[#101622] via-transparent to-transparent"></div>
                            </div>

                            <div className="p-5 text-left">
                                <div className="flex flex-col mb-6">
                                    <h3 className="text-white text-2xl font-bold tracking-tight">{data.make_name} {data.model_name}</h3>
                                    <p className="text-primary font-bold text-[10px] uppercase tracking-widest mt-0.5">{data.model_known_for}</p>

                                    {/* Score Display */}
                                    <div className="mt-4 flex flex-col items-start score-display-bg py-2">
                                        <div className="flex items-baseline gap-2">
                                            <span className="text-cyan-400 text-5xl font-black tracking-tighter neon-glow leading-none">+{data.points_earned || 1}</span>
                                            <span className="text-cyan-400 text-xl font-black tracking-tighter neon-glow">POINT</span>
                                        </div>
                                        <div className="flex items-center gap-1.5 mt-1 opacity-70">
                                            <span className="material-symbols-outlined text-white text-xs">military_tech</span>
                                            <p className="text-[10px] font-bold text-white uppercase tracking-[0.15em]">for Main Guess <span className="text-cyan-400">({data.attempts} Attempts)</span></p>
                                        </div>
                                    </div>
                                </div>

                                {/* Bonus Promo */}
                                {bonusState && (bonusState.year_range_enabled || bonusState.generation_enabled || bonusState.codename_enabled) && (
                                    <>
                                        <div className="bg-primary/10 border border-primary/30 rounded-lg p-3 mb-5 flex items-center gap-3">
                                            <span className="material-symbols-outlined text-cyan-400 text-xl shrink-0">add_circle</span>
                                            <p className="text-[10px] text-white/90 leading-relaxed font-bold uppercase tracking-wider">
                                                <span className="text-cyan-400">MAXIMIZE YOUR SCORE:</span> Tackle bonus questions below to earn additional points!
                                            </p>
                                        </div>

                                        {/* Specs Grid (Bonus Questions) */}
                                        <div className="flex flex-col bg-[#0d121c] rounded-lg border border-white/5 overflow-hidden">
                                            {renderBonusRow('year_range', 'Model Year', bonusState.year_range_points || 1, bonusState.year_range_enabled, bonusState.year_range_attempted, bonusState.year_range_correct)}
                                            {renderBonusRow('generation', 'Generation', bonusState.generation_points || 1, bonusState.generation_enabled, bonusState.generation_attempted, bonusState.generation_correct)}
                                            {renderBonusRow('codename', 'Codename', bonusState.codename_points || 1, bonusState.codename_enabled, bonusState.codename_attempted, bonusState.codename_correct)}
                                        </div>
                                    </>
                                )}
                            </div>
                        </div>

                        {/* Streak Meter */}
                        <div className="w-full">
                            <StreakMeter stats={data.streak_stats} />
                        </div>
                    </div>
                    
                    {/* Scroll Indicator Overlay */}
                    {showScrollIndicator && (
                        <div className="absolute bottom-0 left-0 right-0 h-24 bg-gradient-to-t from-[#0d121c]/90 via-[#0d121c]/40 to-transparent pointer-events-none flex items-end justify-center pb-4 transition-opacity duration-300">
                            <div className="animate-bounce flex flex-col items-center text-cyan-400">
                                <span className="text-[10px] font-bold uppercase tracking-[0.2em] opacity-90 drop-shadow-md shadow-black">More Points Below</span>
                                <span className="material-symbols-outlined text-lg opacity-80 mt-1">keyboard_double_arrow_down</span>
                            </div>
                        </div>
                    )}
                </div>
                </div>

                {/* Fixed Footer Actions */}
                <div className="p-6 sm:p-8 pt-4 sm:pt-4 bg-black/20 border-t border-white/5 shrink-0">
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 w-full mb-6">
                        <button onClick={handleShare} className="flex items-center justify-center gap-2.5 h-12 bg-primary hover:bg-primary/90 text-white rounded-lg font-bold transition-all shadow-[0_4px_15px_rgba(13,91,236,0.3)] text-sm">
                            <span className="material-symbols-outlined text-xl">share</span>
                            SHARE RESULTS
                        </button>
                        <Link
                            href="/leaderboard?source=completion"
                            className="flex items-center justify-center gap-2.5 h-12 bg-[#232f48] hover:bg-[#2d3b5a] text-white border border-white/10 rounded-lg font-bold transition-all text-sm"
                        >
                            <span className="material-symbols-outlined text-xl">leaderboard</span>
                            VIEW STATS
                        </Link>
                    </div>

                    <div className="flex items-center justify-center gap-2 text-white/40 text-[10px] uppercase tracking-widest font-bold">
                        <span className="material-symbols-outlined text-sm">schedule</span>
                        NEXT CHALLENGE IN <span className="text-cyan-400 font-mono text-xs">{formatTime(secondsRemaining)}</span>
                    </div>
                </div>
            </div>

            {/* Bonus Input Modal */}
            {activeBonus && (
                <BonusInputPopup
                    isOpen={!!activeBonus}
                    onClose={() => setActiveBonus(null)}
                    onSubmit={handleBonusSubmit}
                    type={activeBonus.type}
                    label={activeBonus.label}
                    points={activeBonus.points}
                />
            )}
        </div>
    );
}

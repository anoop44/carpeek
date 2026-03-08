'use client';

import { useState, useEffect, Suspense, useRef } from 'react';
import Link from 'next/link';
import { GoogleOAuthProvider, GoogleLogin, CredentialResponse } from '@react-oauth/google';
import { jwtDecode } from "jwt-decode";
import AppHeader from '../components/AppHeader';
import BannerAd from '../components/BannerAd';

interface LeaderboardEntry {
    rank: number;
    user_id: string;
    pilot_name: string; // Keeping variable name for compatibility but displaying as Spotter Name
    profile_picture_url?: string;
    level: number;
    level_title: string;
    score: number;
    main_score?: number;
    bonus_score?: number;
    accuracy: number;
    attempts: number;
    time: string;
    is_current_user: boolean;
}

interface UserStatus {
    attempts: number;
    max_attempts: number;
    is_completed: boolean;
    is_correct: boolean;
}

function LeaderboardContent() {
    const [entries, setEntries] = useState<LeaderboardEntry[]>([]);
    const [loading, setLoading] = useState(true);
    const [viewMode, setViewMode] = useState<'daily' | 'alltime'>('daily');
    const [userRank, setUserRank] = useState<LeaderboardEntry | null>(null);
    const [userId, setUserId] = useState<string>('');
    const [isClient, setIsClient] = useState(false);
    const [stats, setStats] = useState<{ players_today: number; total_players: number; average_accuracy: number; global_average_accuracy: number; total_bonus_points: number } | null>(null);
    const [userStatus, setUserStatus] = useState<UserStatus | null>(null);
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [sessionToken, setSessionToken] = useState<string>('');
    const sessionTokenRef = useRef<string>('');

    // Sync ref with state
    useEffect(() => {
        sessionTokenRef.current = sessionToken;
    }, [sessionToken]);

    const BROWSER_SIGNATURE_KEY = 'carpeek_browser_signature';

    const getBrowserSignature = () => {
        let sig = localStorage.getItem(BROWSER_SIGNATURE_KEY);
        if (!sig) {
            sig = window.crypto.randomUUID();
            localStorage.setItem(BROWSER_SIGNATURE_KEY, sig);
        }
        return sig;
    };

    const getHeaders = () => {
        const headers: Record<string, string> = {
            'Content-Type': 'application/json',
            'X-Timezone': Intl.DateTimeFormat().resolvedOptions().timeZone,
        };
        if (sessionTokenRef.current) {
            headers['Authorization'] = `Bearer ${sessionTokenRef.current}`;
        }
        return headers;
    };

    const initSession = async () => {
        try {
            const signature = getBrowserSignature();
            const res = await fetch('/api/auth/session', {
                headers: { 'X-Browser-Signature': signature }
            });
            if (res.ok) {
                const data = await res.json();
                if (data.token) {
                    sessionTokenRef.current = data.token;
                    setSessionToken(data.token);
                    return data.token;
                }
            }
        } catch (e) {
            console.error("Session init failed", e);
        }
        return null;
    };

    const managedFetch = async (url: string, options: RequestInit = {}): Promise<Response> => {
        const currentHeaders = getHeaders();
        let response = await fetch(url, {
            ...options,
            headers: {
                ...currentHeaders,
                ...(options.headers || {}),
            },
        });

        if (response.status === 401) {
            const newToken = await initSession();
            if (newToken) {
                response = await fetch(url, {
                    ...options,
                    headers: {
                        ...getHeaders(),
                        'Authorization': `Bearer ${newToken}`,
                        ...(options.headers || {}),
                    },
                });
            }
        }
        return response;
    };



    useEffect(() => {
        setIsClient(true);
        const storedUserId = localStorage.getItem(BROWSER_SIGNATURE_KEY);
        const storedEmail = localStorage.getItem('autocorrect_user_email');
        if (storedUserId) {
            setUserId(storedUserId);
        }
        if (storedEmail) {
            setIsLoggedIn(true);
        }
        console.log("Google Client ID:", process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID);
    }, []);

    const bootstrap = async () => {
        await initSession();
        fetchData();
        fetchStats();
    };

    const fetchData = async () => {
        setLoading(true);
        try {
            const response = await managedFetch(`/api/leaderboard?type=${viewMode}`);
            const data = await response.json();
            if (!response.ok) {
                setError(data.error || 'Failed to fetch leaderboard');
                return;
            }
            const leaderboardData = data as LeaderboardEntry[];
            setEntries(leaderboardData || []);

            // Check if current user is in the list
            const me = leaderboardData.find((e: LeaderboardEntry) => e.is_current_user);
            if (me) {
                setUserRank(me);
            } else {
                setUserRank(null);
            }
        } catch (error) {
            console.error('Failed to fetch leaderboard:', error);
            setError('Connection failed. Please check your network.');
        } finally {
            setLoading(false);
        }
    };

    const fetchStats = async () => {
        try {
            // Fetch today's challenge to get the ID, then fetch stats
            const challengeRes = await managedFetch('/api/challenge/today');
            const challengeData = await challengeRes.json();

            if (challengeData?.user_status) {
                setUserStatus(challengeData.user_status);
            }

            if (challengeData?.id) {
                const statsRes = await managedFetch(`/api/challenge/stats?challenge_id=${challengeData.id}`);
                const statsData = await statsRes.json();
                setStats(statsData);
            }
        } catch (error) {
            console.error('Failed to fetch stats:', error);
        }
    };

    useEffect(() => {
        if (isClient) {
            bootstrap();
        }
    }, [viewMode, isClient]);

    const handleGoogleSuccess = async (credentialResponse: CredentialResponse) => {
        if (!credentialResponse.credential) return;

        try {
            const signature = getBrowserSignature();
            const response = await fetch('/api/auth/google', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    id_token: credentialResponse.credential,
                    anonymous_id: signature
                }),
            });

            if (response.ok) {
                const user = await response.json();
                if (user.anonymous_id) {
                    localStorage.setItem(BROWSER_SIGNATURE_KEY, user.anonymous_id);
                    setUserId(user.anonymous_id);
                }
                if (user.email) {
                    localStorage.setItem('autocorrect_user_email', user.email);
                    setIsLoggedIn(true);
                }
                // If token returned, save it
                if (user.auth_token) {
                    setSessionToken(user.auth_token);
                }
                fetchData(); // Refresh leaderboard
            } else {
                console.error("Login failed");
            }
        } catch (error) {
            console.error("Login error", error);
            setError('Failed to sign in with Google');
        }
    };

    return (
        <GoogleOAuthProvider clientId={process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || ''}>
            <div className="relative flex h-auto min-h-screen w-full flex-col group/design-root overflow-x-hidden scanline bg-background-light dark:bg-background-dark text-slate-900 dark:text-slate-100">
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
                <div className="layout-container flex h-full grow flex-col">
                    <AppHeader />

                    <main className="flex-1 flex flex-col items-center py-10 px-6 max-w-[1200px] mx-auto w-full">
                        <div className="w-full flex flex-col items-center">
                            <h1 className="text-white tracking-[0.2em] text-4xl md:text-5xl font-bold leading-tight px-4 text-center pb-2 pt-6 font-display uppercase">
                                GLOBAL LEADERBOARD
                            </h1>
                            <p className="text-primary text-xs font-bold leading-normal pb-8 tracking-[0.4em] px-4 text-center uppercase">
                                Live Standings: Updating Spotter Data
                            </p>
                        </div>

                        {/* Google Sign-In Prompt - show whenever user is not signed in */}
                        {isClient && !isLoggedIn && (
                            <div className="w-full mb-8 p-6 md:p-8 rounded-2xl bg-gradient-to-br from-card-dark via-card-dark to-primary/10 border border-primary/30 flex flex-col items-center gap-6 relative overflow-hidden">
                                {/* Background glow effect */}
                                <div className="absolute -top-20 -right-20 w-60 h-60 bg-primary/10 rounded-full blur-3xl pointer-events-none"></div>
                                <div className="absolute -bottom-10 -left-10 w-40 h-40 bg-accent-neon/5 rounded-full blur-2xl pointer-events-none"></div>

                                <div className="flex flex-col items-center text-center relative z-10">
                                    <div className="flex-shrink-0 size-14 rounded-full bg-primary/20 border-2 border-primary/50 flex items-center justify-center text-primary mb-4 animate-pulse">
                                        <span className="material-symbols-outlined text-3xl" style={{ fontVariationSettings: "'FILL' 1" }}>account_circle</span>
                                    </div>
                                    <h3 className="text-white font-bold tracking-widest uppercase text-base md:text-lg mb-2 flex items-center gap-2">
                                        Show the World Who You Are
                                    </h3>
                                    <p className="text-slate-400 text-sm font-medium leading-relaxed max-w-lg">
                                        Sign in with Google to display <span className="text-primary font-semibold">your name</span> and <span className="text-primary font-semibold">profile photo</span> on the leaderboard. Let other spotters know who&apos;s climbing the ranks!
                                    </p>
                                </div>

                                {/* Prominent Google Sign-In Button */}
                                <div className="relative z-10 mt-2">
                                    <div className="rounded-full p-[2px] bg-gradient-to-r from-primary via-accent-neon to-primary animate-gradient-x">
                                        <div className="rounded-full bg-card-dark px-1 py-1">
                                            <GoogleLogin
                                                onSuccess={handleGoogleSuccess}
                                                onError={() => {
                                                    console.log('Login Failed');
                                                }}
                                                theme="filled_black"
                                                text="signin_with"
                                                shape="pill"
                                                size="large"
                                                width="300"
                                                logo_alignment="left"
                                            />
                                        </div>
                                    </div>
                                </div>
                            </div>
                        )}

                        <div className="w-full max-w-md mb-10">
                            <div className="flex p-1 rounded-xl bg-card-dark border border-[#232f48]">
                                <label className={`flex cursor-pointer h-12 grow items-center justify-center overflow-hidden rounded-lg px-2 text-sm font-bold leading-normal transition-all tracking-widest uppercase ${viewMode === 'daily' ? 'bg-primary text-white' : 'text-slate-400'}`}>
                                    <span className="truncate">Daily Circuit</span>
                                    <input checked={viewMode === 'daily'} onChange={() => setViewMode('daily')} className="invisible w-0" name="view-toggle" type="radio" value="daily" />
                                </label>
                                <label className={`flex cursor-pointer h-12 grow items-center justify-center overflow-hidden rounded-lg px-2 text-sm font-bold leading-normal transition-all tracking-widest uppercase ${viewMode === 'alltime' ? 'bg-primary text-white' : 'text-slate-400'}`}>
                                    <span className="truncate">Hall of Fame</span>
                                    <input checked={viewMode === 'alltime'} onChange={() => setViewMode('alltime')} className="invisible w-0" name="view-toggle" type="radio" value="alltime" />
                                </label>
                            </div>
                        </div>

                        {/* Stats Cards */}
                        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 w-full mb-10">
                            <div className="flex flex-col gap-2 rounded-xl p-6 bg-card-dark border border-[#232f48] relative overflow-hidden group">
                                <div className="absolute -top-2 -right-2 opacity-10 group-hover:opacity-20 transition-opacity">
                                    <span className="material-symbols-outlined" style={{ fontSize: '80px' }}>group</span>
                                </div>
                                <p className="text-slate-400 text-xs font-bold tracking-widest uppercase leading-normal">
                                    {viewMode === 'daily' ? 'Total Spotters' : 'Active Today'}
                                </p>
                                <p className="text-white tracking-tight text-3xl font-bold leading-tight font-display">
                                    {stats ? (viewMode === 'daily' ? stats.total_players.toLocaleString() : stats.players_today.toLocaleString()) : '---'}
                                </p>
                            </div>
                            <div className="flex flex-col gap-2 rounded-xl p-6 bg-card-dark border border-primary/50 glow-border relative overflow-hidden">
                                <div className="absolute -top-2 -right-2 opacity-20 text-primary transition-opacity">
                                    <span className="material-symbols-outlined" style={{ fontSize: '80px' }}>speed</span>
                                </div>
                                <p className="text-primary text-xs font-bold tracking-widest uppercase leading-normal">
                                    {viewMode === 'daily' ? 'Today\'s Players' : 'Total Spotters'}
                                </p>
                                <p className="text-white tracking-tight text-3xl font-bold leading-tight font-display">
                                    {viewMode === 'daily' ? (entries.length || '---') : (stats?.total_players.toLocaleString() || '---')}
                                </p>
                            </div>
                            <div className="flex flex-col gap-2 rounded-xl p-6 bg-card-dark border border-[#232f48] relative overflow-hidden">
                                <div className="absolute -top-2 -right-2 opacity-10 group-hover:opacity-20 transition-opacity">
                                    <span className="material-symbols-outlined" style={{ fontSize: '80px' }}>target</span>
                                </div>
                                <p className="text-slate-400 text-xs font-bold tracking-widest uppercase leading-normal">
                                    {viewMode === 'daily' ? 'Crowd Accuracy' : 'Global Accuracy'}
                                </p>
                                <p className="text-white tracking-tight text-3xl font-bold leading-tight font-display">
                                    {stats ? `${Math.round(viewMode === 'daily' ? stats.average_accuracy : stats.global_average_accuracy)}%` : '--%'}
                                </p>
                            </div>
                        </div>


                        {/* Leaderboard Banner Ad */}
                        <div className="w-full mb-6">
                            <BannerAd
                                adSlot={process.env.NEXT_PUBLIC_ADSENSE_SLOT_LEADERBOARD || ''}
                                label="Sponsored"
                            />
                        </div>

                        {/* Leaderboard Table */}
                        <div className="w-full bg-card-dark/40 border border-[#232f48] rounded-xl overflow-hidden backdrop-blur-sm">
                            <div className="overflow-x-auto">
                                <table className="w-full text-left border-collapse">
                                    <thead>
                                        <tr className="bg-card-dark border-b border-[#232f48]">
                                            <th className="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest">Rank</th>
                                            <th className="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest">Spotter Name</th>
                                            <th className="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest text-right">Score</th>
                                            <th className="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest text-right">Accuracy</th>
                                        </tr>
                                    </thead>
                                    <tbody className="divide-y divide-[#232f48]/50">
                                        {loading ? (
                                            <tr>
                                                <td colSpan={4} className="px-6 py-8 text-center text-slate-500 font-medium">Searching for spotters...</td>
                                            </tr>
                                        ) : (
                                            <>
                                                {/* Ghost Row if empty and user played but not logged in */}
                                                {entries.length === 0 && userStatus && !isLoggedIn && (
                                                    <tr className="bg-primary/5 border-b border-primary/20 relative overflow-hidden">
                                                        <td className="px-6 py-5">
                                                            <div className="flex items-center gap-3">
                                                                <span className="text-slate-500 font-bold font-display">--</span>
                                                            </div>
                                                        </td>
                                                        <td className="px-6 py-5">
                                                            <div className="flex items-center gap-3">
                                                                <div className="size-10 rounded-lg flex items-center justify-center text-white bg-slate-800 border border-white/10 border-dashed">
                                                                    <span className="material-symbols-outlined text-slate-500">person</span>
                                                                </div>
                                                                <div>
                                                                    <p className="text-slate-300 font-bold tracking-widest uppercase text-sm">
                                                                        YOU <span className="text-primary text-[10px] ml-1">(UNCLAIMED)</span>
                                                                    </p>
                                                                    <p className="text-slate-500 text-[10px] font-bold">
                                                                        THIS COULD BE YOU IF LOGGED IN
                                                                    </p>
                                                                </div>
                                                            </div>
                                                        </td>
                                                        <td className="px-6 py-5 font-display text-slate-500 text-right">
                                                            --
                                                        </td>
                                                        <td className="px-6 py-5 font-display text-slate-300 text-right">
                                                            {userStatus.is_correct ? '100%' : '0%'}
                                                        </td>
                                                    </tr>
                                                )}

                                                {/* Standard Empty State (only if no ghost row is shown, or maybe if ghost row is shown but we still want to say "no other pilots"?) 
                                                Actually, if ghost row is shown, we probably don't want "No pilots" or maybe we do? 
                                                Let's only show "No pilots" if entries is empty AND no ghost row.
                                            */}
                                                {entries.length === 0 && (!userStatus || isLoggedIn) && (
                                                    <tr>
                                                        <td colSpan={4} className="px-6 py-8 text-center text-slate-500 font-medium tracking-widest uppercase">No spotters found for this circuit yet. Be the first!</td>
                                                    </tr>
                                                )}

                                                {entries.map((entry) => {
                                                    // Determine specific styles based on rank
                                                    let rowClass = "hover:bg-white/5 transition-colors";
                                                    let rankTextClass = "text-slate-500 font-bold font-display";
                                                    let avatarClass = "bg-slate-800";
                                                    let accuracyClass = "text-slate-300";
                                                    let pilotNameClass = "text-slate-300";

                                                    if (entry.rank === 1) {
                                                        rowClass = "bg-primary/10 group hover:bg-primary/20 transition-colors";
                                                        rankTextClass = "text-primary font-black text-2xl font-display italic neon-text-glow";
                                                        avatarClass = "bg-gradient-to-br from-primary to-accent-neon";
                                                        accuracyClass = "text-primary";
                                                        pilotNameClass = "text-white";
                                                    } else if (entry.rank === 2) {
                                                        rowClass = "bg-white/5 group hover:bg-white/10 transition-colors";
                                                        rankTextClass = "text-slate-300 font-black text-2xl font-display italic";
                                                        avatarClass = "bg-slate-700";
                                                        accuracyClass = "text-white";
                                                        pilotNameClass = "text-white";
                                                    } else if (entry.rank === 3) {
                                                        rowClass = "bg-white/5 group hover:bg-white/10 transition-colors";
                                                        rankTextClass = "text-amber-600/80 font-black text-2xl font-display italic";
                                                        avatarClass = "bg-slate-800";
                                                        accuracyClass = "text-white";
                                                        pilotNameClass = "text-white";
                                                    }

                                                    // Override if current user
                                                    if (entry.is_current_user) {
                                                        rowClass = "bg-gradient-to-r from-orange-500/20 to-transparent border-l-4 border-orange-500 relative z-10 shadow-[0_4px_20px_rgba(249,115,22,0.15)]";
                                                        pilotNameClass = "text-white drop-shadow-md";
                                                        accuracyClass = "text-white font-bold drop-shadow-md";
                                                    }

                                                    return (
                                                        <tr key={entry.rank} className={rowClass}>
                                                            <td className={entry.rank <= 3 ? "px-6 py-6" : "px-6 py-5"}>
                                                                <div className={entry.rank <= 3 ? "flex items-center gap-3" : ""}>
                                                                    <span className={rankTextClass}>
                                                                        {String(entry.rank).padStart(2, '0')}
                                                                    </span>
                                                                    {entry.rank === 1 && (
                                                                        <span className="material-symbols-outlined text-primary" style={{ fontVariationSettings: "'FILL' 1" }}>military_tech</span>
                                                                    )}
                                                                    {entry.rank === 2 && (
                                                                        <span className="material-symbols-outlined text-slate-400">military_tech</span>
                                                                    )}
                                                                    {entry.rank === 3 && (
                                                                        <span className="material-symbols-outlined text-amber-700/80">military_tech</span>
                                                                    )}
                                                                </div>
                                                            </td>
                                                            <td className={entry.rank <= 3 ? "px-6 py-6" : "px-6 py-5"}>
                                                                {entry.rank <= 3 ? (
                                                                    <div className="flex items-center gap-3">
                                                                        <div className={`size-10 rounded-lg flex items-center justify-center text-white ${avatarClass} overflow-hidden`}>
                                                                            {entry.profile_picture_url ? (
                                                                                <img src={entry.profile_picture_url} alt={entry.pilot_name} className="w-full h-full object-cover" referrerPolicy="no-referrer" />
                                                                            ) : (
                                                                                <span className="material-symbols-outlined">
                                                                                    {entry.rank === 1 ? 'hexagon' : (entry.rank === 2 ? 'token' : 'grid_view')}
                                                                                </span>
                                                                            )}
                                                                        </div>
                                                                        <div>
                                                                            <p className={`${pilotNameClass} font-bold tracking-widest uppercase text-sm flex items-center gap-2`}>
                                                                                {entry.pilot_name} {entry.is_current_user ? <span className="px-1.5 py-0.5 rounded text-[9px] bg-orange-700 text-white font-black animate-pulse">YOU</span> : ''}
                                                                            </p>
                                                                            <div className="flex items-center gap-1.5 mt-0.5">
                                                                                <span className={`text-[9px] md:text-[11px] font-black px-1.5 py-0.5 rounded text-black ${entry.rank === 1 ? "bg-primary" : "bg-slate-400"}`}>
                                                                                    LVL {entry.level}
                                                                                </span>
                                                                                <span className={entry.rank === 1 ? "text-primary text-[10px] md:text-xs font-bold" : "text-slate-500 text-[10px] md:text-xs font-bold"}>
                                                                                    {entry.level_title}
                                                                                </span>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                ) : (
                                                                    <div className="flex items-center gap-3">
                                                                        {entry.profile_picture_url && (
                                                                            <div className="size-8 rounded-lg overflow-hidden flex-shrink-0">
                                                                                <img src={entry.profile_picture_url} alt={entry.pilot_name} className="w-full h-full object-cover" referrerPolicy="no-referrer" />
                                                                            </div>
                                                                        )}
                                                                        <div className="flex flex-col">
                                                                            <p className={`${pilotNameClass} font-medium tracking-wider uppercase text-sm flex items-center gap-2`}>
                                                                                {entry.pilot_name} {entry.is_current_user ? <span className="px-1.5 py-0.5 rounded text-[9px] bg-orange-700 text-white font-black animate-pulse">YOU</span> : ''}
                                                                            </p>
                                                                            <div className="flex items-center gap-1.5 mt-0.5 opacity-60">
                                                                                <span className="text-[9px] md:text-[11px] font-black px-1.5 py-0.5 rounded bg-slate-700 text-slate-300">
                                                                                    LVL {entry.level}
                                                                                </span>
                                                                                <span className="text-slate-500 text-[9px] md:text-xs font-bold uppercase">
                                                                                    {entry.level_title}
                                                                                </span>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                )}
                                                            </td>
                                                            <td className={`${entry.rank <= 3 ? "px-6 py-6 text-lg" : "px-6 py-5"} font-display ${accuracyClass} ${entry.rank <= 3 ? "font-bold" : ""} text-right`}>
                                                                {viewMode === 'daily' && entry.main_score !== undefined && entry.bonus_score !== undefined ? (
                                                                    <span>
                                                                        <span className="font-bold">{entry.score}</span>
                                                                        <span className="text-xs text-slate-500 ml-1 opacity-70">({entry.main_score}+{entry.bonus_score})</span>
                                                                    </span>
                                                                ) : (
                                                                    <span className="font-bold">{entry.score} points</span>
                                                                )}
                                                            </td>
                                                            <td className={`${entry.rank <= 3 ? "px-6 py-6 text-lg" : "px-6 py-5"} font-display ${accuracyClass} ${entry.rank <= 3 ? "font-bold" : ""} text-right`}>
                                                                {`${entry.accuracy.toFixed(1)}%`}
                                                            </td>
                                                        </tr>
                                                    );
                                                })}
                                            </>
                                        )}
                                    </tbody>
                                </table>
                            </div>
                        </div>

                        <div className="mt-20 flex flex-col md:flex-row items-center justify-between w-full text-slate-500 text-[10px] tracking-[0.3em] font-black uppercase gap-4 border-t border-white/5 pt-8 pb-12">
                            <div className="flex items-center gap-2">
                                <span className="inline-block size-1.5 rounded-full bg-green-500 animate-pulse ring-4 ring-green-500/5"></span>
                                LIVE NETWORK DATA
                            </div>
                            <div className="flex items-center gap-8">
                                <Link className="hover:text-primary transition-colors" href="/privacy">Privacy Policy</Link>
                                <Link className="hover:text-primary transition-colors" href="/terms">Terms and Conditions</Link>
                                <span className="text-slate-800 tracking-normal">STABLE-v2.4.0</span>
                            </div>
                        </div>
                    </main>
                </div>
            </div>
        </GoogleOAuthProvider>
    );
}

export default function LeaderboardPage() {
    return (
        <Suspense fallback={<div className="min-h-screen bg-background-dark text-white flex items-center justify-center">Loading...</div>}>
            <LeaderboardContent />
        </Suspense>
    );
}

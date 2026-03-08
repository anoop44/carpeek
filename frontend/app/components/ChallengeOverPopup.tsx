'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';

interface ChallengeOverPopupProps {
    isOpen: boolean;
    onClose: () => void;
    isCorrect: boolean;
    nextChallengeSeconds: number | null;
    solution?: {
        make_name: string;
        model_name: string;
    };
}

export default function ChallengeOverPopup({ isOpen, onClose, isCorrect, nextChallengeSeconds, solution }: ChallengeOverPopupProps) {
    const [countdown, setCountdown] = useState<number | null>(nextChallengeSeconds);

    useEffect(() => {
        if (!isOpen) return;
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === 'Escape') onClose();
        };
        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [isOpen, onClose]);

    useEffect(() => {
        setCountdown(nextChallengeSeconds);
    }, [nextChallengeSeconds]);

    useEffect(() => {
        if (countdown === null || countdown <= 0) return;
        const timer = setInterval(() => {
            setCountdown(prev => (prev && prev > 0 ? prev - 1 : 0));
        }, 1000);
        return () => clearInterval(timer);
    }, [countdown]);

    const formatTime = (totalSeconds: number | null) => {
        if (totalSeconds === null || totalSeconds <= 0) return { hours: '00', minutes: '00', seconds: '00' };
        const hours = Math.floor(totalSeconds / 3600).toString().padStart(2, '0');
        const minutes = Math.floor((totalSeconds % 3600) / 60).toString().padStart(2, '0');
        const secs = (totalSeconds % 60).toString().padStart(2, '0');
        return { hours, minutes, seconds: secs };
    };

    if (!isOpen) return null;

    const time = formatTime(countdown);

    return (
        <div className="fixed inset-0 z-[100] flex items-center justify-center p-4" onClick={onClose}>
            {/* Backdrop */}
            <div className="absolute inset-0 bg-black/70 backdrop-blur-sm animate-fade-in"></div>

            {/* Popup Content */}
            <div
                className="relative w-full max-w-md animate-scale-in"
                onClick={(e) => e.stopPropagation()}
            >
                {/* Glow Effect */}
                <div className={`absolute -inset-1 rounded-2xl blur-lg opacity-40 ${isCorrect ? 'bg-gradient-to-r from-green-500 to-accent-neon' : 'bg-gradient-to-r from-amber-500 to-primary'}`}></div>

                <div className="relative bg-card-dark border border-white/10 rounded-2xl overflow-hidden">
                    {/* Top Accent Bar */}
                    <div className={`h-1 ${isCorrect ? 'bg-gradient-to-r from-green-500 via-accent-neon to-green-500' : 'bg-gradient-to-r from-amber-500 via-primary to-amber-500'}`}></div>

                    {/* Close Button */}
                    <button
                        onClick={onClose}
                        className="absolute top-4 right-4 z-10 text-slate-500 hover:text-white transition-colors"
                    >
                        <span className="material-symbols-outlined">close</span>
                    </button>

                    <div className="p-8 flex flex-col items-center text-center">
                        {/* Status Icon */}
                        <div className={`size-20 rounded-full flex items-center justify-center mb-6 ${isCorrect ? 'bg-green-500/15 border-2 border-green-500/40' : 'bg-amber-500/15 border-2 border-amber-500/40'}`}>
                            <span
                                className={`material-symbols-outlined text-5xl ${isCorrect ? 'text-green-400' : 'text-amber-400'}`}
                                style={{ fontVariationSettings: "'FILL' 1" }}
                            >
                                {isCorrect ? 'verified' : 'block'}
                            </span>
                        </div>

                        {/* Heading */}
                        <h2 className="text-white font-bold text-xl tracking-widest uppercase mb-2">
                            {isCorrect ? 'Already Spotted!' : 'No Attempts Left'}
                        </h2>
                        <p className="text-slate-400 text-sm leading-relaxed mb-2 max-w-xs">
                            {isCorrect
                                ? "You've already identified today's car successfully."
                                : "You've exhausted all your attempts for today's challenge."
                            }
                        </p>

                        {/* Solution Display (if available) */}
                        {solution && (
                            <div className={`mt-2 mb-6 px-4 py-2 rounded-lg border ${isCorrect ? 'bg-green-500/5 border-green-500/20' : 'bg-amber-500/5 border-amber-500/20'}`}>
                                <p className="text-xs text-slate-500 uppercase tracking-widest mb-1">Answer</p>
                                <p className={`font-bold tracking-wider uppercase ${isCorrect ? 'text-green-400' : 'text-amber-400'}`}>
                                    {solution.make_name} {solution.model_name}
                                </p>
                            </div>
                        )}

                        {/* Countdown Timer */}
                        <div className="w-full mt-4 mb-6">
                            <p className="text-slate-500 text-[10px] font-bold uppercase tracking-[0.3em] mb-4">
                                Next Challenge In
                            </p>
                            <div className="flex items-center justify-center gap-3">
                                {/* Hours */}
                                <div className="flex flex-col items-center">
                                    <div className="bg-background-dark border border-white/10 rounded-xl px-5 py-4 min-w-[72px]">
                                        <span className="text-3xl font-black font-display text-white tracking-tight">
                                            {time.hours}
                                        </span>
                                    </div>
                                    <span className="text-[9px] text-slate-600 uppercase tracking-widest mt-2 font-bold">Hours</span>
                                </div>

                                <span className="text-2xl font-bold text-slate-600 -mt-4 animate-pulse">:</span>

                                {/* Minutes */}
                                <div className="flex flex-col items-center">
                                    <div className="bg-background-dark border border-white/10 rounded-xl px-5 py-4 min-w-[72px]">
                                        <span className="text-3xl font-black font-display text-white tracking-tight">
                                            {time.minutes}
                                        </span>
                                    </div>
                                    <span className="text-[9px] text-slate-600 uppercase tracking-widest mt-2 font-bold">Minutes</span>
                                </div>

                                <span className="text-2xl font-bold text-slate-600 -mt-4 animate-pulse">:</span>

                                {/* Seconds */}
                                <div className="flex flex-col items-center">
                                    <div className="bg-background-dark border border-white/10 rounded-xl px-5 py-4 min-w-[72px]">
                                        <span className="text-3xl font-black font-display text-primary tracking-tight">
                                            {time.seconds}
                                        </span>
                                    </div>
                                    <span className="text-[9px] text-slate-600 uppercase tracking-widest mt-2 font-bold">Seconds</span>
                                </div>
                            </div>
                        </div>

                        {/* Action Buttons */}
                        <div className="w-full flex flex-col gap-3">
                            <Link
                                href="/leaderboard"
                                className="w-full py-4 bg-primary hover:bg-primary/90 text-white rounded-xl font-bold text-sm tracking-widest uppercase transition-all flex items-center justify-center gap-2"
                            >
                                <span className="material-symbols-outlined text-lg">leaderboard</span>
                                View Leaderboard
                            </Link>
                            <button
                                onClick={onClose}
                                className="w-full py-3 bg-white/5 hover:bg-white/10 text-slate-400 rounded-xl font-bold text-xs tracking-widest uppercase transition-all border border-white/5"
                            >
                                Dismiss
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

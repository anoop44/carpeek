'use client';

import { useEffect } from 'react';

interface StreakStats {
    attendance_streak: number; // Consecutive logins
    submission_streak: number; // Consecutive played games
    max_attendance_streak: number;
    max_submission_streak: number;
    total_days_participated: number;
    total_days_submitted: number;
}

interface StreakPopupProps {
    isOpen: boolean;
    onClose: () => void;
    stats: StreakStats | null | undefined;
}

export default function StreakPopup({ isOpen, onClose, stats }: StreakPopupProps) {
    useEffect(() => {
        if (!isOpen) return;
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === 'Escape') onClose();
        };
        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [isOpen, onClose]);

    if (!isOpen || !stats) return null;

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 animate-in fade-in duration-200">
            {/* Backdrop */}
            <div className="fixed inset-0 bg-black/80 backdrop-blur-sm z-0" onClick={onClose}></div>

            {/* Modal */}
            <div className="glass-panel w-full max-w-[400px] rounded-xl overflow-hidden flex flex-col relative z-20 border border-white/10 shadow-2xl animate-in zoom-in-95 duration-200 scanline">
                {/* Header Scanline */}
                <div className="h-1 bg-gradient-to-r from-transparent via-orange-500 to-transparent opacity-80 animate-pulse"></div>

                <div className="p-6 relative">
                    {/* Close Button */}
                    <button
                        onClick={onClose}
                        className="absolute top-4 right-4 text-white/40 hover:text-white transition-colors z-30"
                    >
                        <span className="material-symbols-outlined">close</span>
                    </button>

                    <div className="flex flex-col items-center mb-8 relative">
                        {/* Background glow behind icon */}
                        <div className="absolute top-0 left-1/2 -translate-x-1/2 w-32 h-32 bg-orange-500/20 blur-[50px] rounded-full pointer-events-none"></div>

                        <div className="w-20 h-20 rounded-full bg-gradient-to-b from-[#192233] to-black border border-orange-500/30 flex items-center justify-center mb-4 shadow-[0_0_20px_rgba(249,115,22,0.3)] relative z-10">
                            <span className="material-symbols-outlined text-orange-500 text-4xl drop-shadow-[0_0_8px_rgba(249,115,22,0.8)]" style={{ fontVariationSettings: "'FILL' 1" }}>local_fire_department</span>
                        </div>
                        <h2 className="text-2xl font-bold text-white tracking-widest uppercase mb-1 drop-shadow-md">Spotter Status</h2>
                        <p className="text-[10px] font-bold text-orange-400 uppercase tracking-[0.2em] opacity-90">Spotting Records</p>
                    </div>

                    {/* Main Stats Grid */}
                    <div className="grid grid-cols-2 gap-4 mb-6">
                        {/* Current Streak (Main Focus) */}
                        <div className="col-span-2 bg-[#0F141E]/80 border border-orange-500/30 rounded-lg p-5 flex items-center justify-between relative overflow-hidden group">
                            <div className="absolute inset-0 bg-gradient-to-r from-orange-500/5 to-transparent opacity-0 group-hover:opacity-100 transition-opacity"></div>
                            <div className="relative z-10">
                                <p className="text-[10px] font-bold text-orange-200/70 uppercase tracking-wider mb-1">Current Streak</p>
                                <div className="flex items-baseline gap-2">
                                    <p className="text-4xl font-black text-white tracking-tighter drop-shadow-[0_0_10px_rgba(255,255,255,0.3)]">{stats.submission_streak}</p>
                                    <span className="text-xs font-bold text-orange-500 uppercase">Days</span>
                                </div>
                            </div>
                            <span className="material-symbols-outlined text-5xl text-orange-500 opacity-10 absolute right-4 bottom-2 -rotate-12">trending_up</span>
                        </div>

                        {/* Best Streak */}
                        <div className="bg-[#0F141E]/60 border border-white/5 rounded-lg p-4 hover:bg-[#0F141E]/80 transition-colors">
                            <div className="flex items-center gap-1.5 mb-2 opacity-50">
                                <span className="material-symbols-outlined text-sm">stars</span>
                                <p className="text-[9px] font-bold uppercase tracking-wider">Best Streak</p>
                            </div>
                            <div className="flex items-baseline gap-1">
                                <p className="text-2xl font-bold text-white">{stats.max_submission_streak}</p>
                                <span className="text-[9px] text-white/30 uppercase font-bold">Days</span>
                            </div>
                        </div>

                        {/* Total Solved */}
                        <div className="bg-[#0F141E]/60 border border-white/5 rounded-lg p-4 hover:bg-[#0F141E]/80 transition-colors">
                            <div className="flex items-center gap-1.5 mb-2 opacity-50">
                                <span className="material-symbols-outlined text-sm">assignment_turned_in</span>
                                <p className="text-[9px] font-bold uppercase tracking-wider">Total Found</p>
                            </div>
                            <div className="flex items-baseline gap-1">
                                <p className="text-2xl font-bold text-white">{stats.total_days_submitted}</p>
                                <span className="text-[9px] text-white/30 uppercase font-bold">Cars</span>
                            </div>
                        </div>

                        {/* Recent Activity (Attendance) -> "Login Streak" maybe? */}
                        <div className="col-span-2 bg-[#0F141E]/40 border border-white/5 rounded-lg p-3 flex items-center justify-between opacity-70 hover:opacity-100 transition-opacity">
                            <div className="flex items-center gap-2">
                                <span className="material-symbols-outlined text-cyan-400 text-lg">garage</span>
                                <span className="text-[10px] font-bold text-gray-300 uppercase tracking-wider">Garage Visits</span>
                            </div>
                            <span className="text-sm font-bold text-white">{stats.attendance_streak} Days</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

'use client';

import { useState, useEffect } from 'react';

interface HowToPlayPopupProps {
    isOpen: boolean;
    onClose: () => void;
}

export default function HowToPlayPopup({ isOpen, onClose }: HowToPlayPopupProps) {
    useEffect(() => {
        if (!isOpen) return;
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === 'Escape') onClose();
        };
        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [isOpen, onClose]);

    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 z-[100] flex items-center justify-center p-4" onClick={onClose}>
            {/* Backdrop */}
            <div className="absolute inset-0 bg-black/70 backdrop-blur-sm animate-fade-in"></div>

            {/* Popup Content */}
            <div
                className="relative w-full max-w-md md:max-w-2xl animate-scale-in"
                onClick={(e) => e.stopPropagation()}
            >
                {/* Glow Effect */}
                <div className="absolute -inset-1 rounded-2xl blur-lg opacity-30 bg-gradient-to-r from-primary to-accent-neon"></div>

                <div className="relative bg-card-dark border border-white/10 rounded-2xl overflow-hidden flex flex-col max-h-[85vh]">
                    {/* Top Accent Bar */}
                    <div className="h-1 bg-gradient-to-r from-primary via-accent-neon to-primary"></div>

                    {/* Header */}
                    <div className="relative p-6 pb-2 border-b border-white/5 flex items-center justify-between shrink-0">
                        <div className="flex items-center gap-3">
                            <div className="size-10 rounded-full bg-primary/20 border border-primary/30 flex items-center justify-center">
                                <span className="material-symbols-outlined text-primary text-xl" style={{ fontVariationSettings: "'FILL' 1" }}>
                                    help
                                </span>
                            </div>
                            <h2 className="text-white font-bold text-lg md:text-2xl tracking-widest uppercase">
                                How to Play
                            </h2>
                        </div>
                        <button
                            onClick={onClose}
                            className="w-8 h-8 flex items-center justify-center rounded-full bg-white/5 hover:bg-white/10 text-slate-400 hover:text-white transition-colors"
                        >
                            <span className="material-symbols-outlined text-sm">close</span>
                        </button>
                    </div>

                    {/* Scrollable Content */}
                    <div className="p-6 overflow-y-auto overflow-x-hidden space-y-6 custom-scrollbar">

                        {/* The Objective */}
                        <div>
                            <h3 className="text-primary font-bold text-sm md:text-base tracking-widest uppercase mb-2 flex items-center gap-2">
                                <span className="material-symbols-outlined text-base">directions_car</span>
                                The Objective
                            </h3>
                            <p className="text-slate-300 text-sm md:text-base leading-relaxed">
                                Identify the car from the daily image. You have <strong>3 attempts</strong> to correctly guess the <strong>Make</strong> and <strong>Model</strong>.
                            </p>
                        </div>

                        {/* Clues & Feedback */}
                        <div className='hidden'>
                            <h3 className="text-primary font-bold text-sm tracking-widest uppercase mb-3 flex items-center gap-2">
                                <span className="material-symbols-outlined text-base">track_changes</span>
                                Feedback System
                            </h3>
                            <p className="text-slate-300 text-sm leading-relaxed mb-3">
                                After each guess, the tiles will change colors to show how close you are to the actual car:
                            </p>

                            <div className="space-y-3">
                                <div className="flex items-center gap-3">
                                    <div className="w-8 h-8 rounded border border-green-500/40 bg-green-500/20 text-green-400 flex items-center justify-center font-bold text-xs shrink-0">
                                        <span className="material-symbols-outlined text-sm">check</span>
                                    </div>
                                    <p className="text-slate-300 text-sm"><span className="text-green-400 font-bold">Green:</span> Exact match! You got this part perfectly correct.</p>
                                </div>
                                <div className="flex items-center gap-3">
                                    <div className="w-8 h-8 rounded border border-amber-500/40 bg-amber-500/20 text-amber-400 flex items-center justify-center font-bold text-xs shrink-0">
                                        <span className="material-symbols-outlined text-sm">cycle</span>
                                    </div>
                                    <p className="text-slate-300 text-sm"><span className="text-amber-400 font-bold">Yellow/Orange:</span> Partial match! This usually means the make is right, but the model or generation is wrong.</p>
                                </div>
                                <div className="flex items-center gap-3">
                                    <div className="w-8 h-8 rounded border border-red-500/40 bg-red-500/20 text-red-400 flex items-center justify-center font-bold text-xs shrink-0">
                                        <span className="material-symbols-outlined text-sm">arrow_upward</span>
                                    </div>
                                    <p className="text-slate-300 text-sm"><span className="text-red-400 font-bold">Arrows:</span> For numbers (like years), arrows will indicate if the correct answer is higher (↑) or lower (↓).</p>
                                </div>
                                <div className="flex items-center gap-3">
                                    <div className="w-8 h-8 rounded border border-white/10 bg-slate-800 text-slate-400 flex items-center justify-center font-bold text-xs shrink-0">
                                        <span className="material-symbols-outlined text-sm">close</span>
                                    </div>
                                    <p className="text-slate-300 text-sm"><span className="text-slate-400 font-bold">Gray:</span> Incorrect completely. Not the right make, model, or generation.</p>
                                </div>
                            </div>
                        </div>

                        {/* Bonus Questions */}
                        <div>
                            <h3 className="text-primary font-bold text-sm md:text-base tracking-widest uppercase mb-2 flex items-center gap-2">
                                <span className="material-symbols-outlined text-base">stars</span>
                                Bonus Questions
                            </h3>
                            <p className="text-slate-300 text-sm md:text-base leading-relaxed mb-2">
                                After a successful guess, you might be asked <strong>Generation</strong> and <strong>Year</strong> bonus questions. These are not available for every challenge.
                            </p>
                            <ul className="text-slate-300 text-sm md:text-base leading-relaxed space-y-2 list-disc pl-5 marker:text-primary">
                                <li><strong>Generation:</strong> The specific platform or iteration (e.g., Mk7, E46). For numbered generations, use the full word (e.g., <em>First</em>, <em>Second</em>) without adding "gen" or "generation".</li>
                                <li><strong>Production Year:</strong> Guessing <em>any</em> year within the model's production range (e.g., 2015–2020) is considered correct!</li>
                            </ul>
                        </div>

                        {/* Points & Leaderboard */}
                        <div>
                            <h3 className="text-primary font-bold text-sm md:text-base tracking-widest uppercase mb-2 flex items-center gap-2">
                                <span className="material-symbols-outlined text-base">military_tech</span>
                                Points System
                            </h3>
                            <div className="space-y-2">
                                <div className="flex justify-between items-center text-slate-300 text-sm md:text-base">
                                    <span>Correct Make & Model (1st Try)</span>
                                    <span className="text-accent-neon font-bold">5 Points</span>
                                </div>
                                <div className="flex justify-between items-center text-slate-300 text-sm md:text-base">
                                    <span>Correct Make & Model (2nd Try)</span>
                                    <span className="text-accent-neon font-bold">3 Points</span>
                                </div>
                                <div className="flex justify-between items-center text-slate-300 text-sm md:text-base">
                                    <span>Correct Make & Model (3rd Try)</span>
                                    <span className="text-accent-neon font-bold">1 Point</span>
                                </div>
                                <div className="flex justify-between items-center text-slate-400 text-xs md:text-sm border-t border-white/5 pt-1">
                                    <span>Correct Make Only (3rd Try)</span>
                                    <span className="text-accent-neon/80 font-bold">0.5 Points</span>
                                </div>
                                <div className="flex justify-between items-center text-slate-400 text-xs md:text-sm">
                                    <span>Each Bonus Question Correct</span>
                                    <span className="text-accent-neon/80 font-bold">+1 Point</span>
                                </div>
                            </div>
                            <p className="text-slate-300 text-[10px] md:text-xs leading-relaxed mt-4 italic text-center opacity-50">
                                Check back every day at Midnight for a new challenge!
                            </p>
                        </div>
                    </div>

                    {/* Action Button */}
                    <div className="p-6 pt-2 shrink-0 border-t border-white/5">
                        <button
                            onClick={onClose}
                            className="w-full py-3.5 bg-primary hover:bg-primary/90 text-white rounded-xl font-bold text-sm tracking-widest uppercase transition-all"
                        >
                            Got It, Let's Play!
                        </button>
                    </div>
                </div>
            </div>
            {/* Custom scrollbar styles if not globally added */}
            <style jsx global>{`
                .custom-scrollbar::-webkit-scrollbar {
                    width: 6px;
                }
                .custom-scrollbar::-webkit-scrollbar-track {
                    background: rgba(255, 255, 255, 0.02);
                    border-radius: 8px;
                }
                .custom-scrollbar::-webkit-scrollbar-thumb {
                    background: rgba(255, 255, 255, 0.1);
                    border-radius: 8px;
                }
                .custom-scrollbar::-webkit-scrollbar-thumb:hover {
                    background: rgba(255, 255, 255, 0.2);
                }
            `}</style>
        </div>
    );
}

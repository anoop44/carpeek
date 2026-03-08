'use client';

import { useState, useEffect, useRef } from 'react';

interface BonusInputPopupProps {
    isOpen: boolean;
    onClose: () => void;
    onSubmit: (value: string) => Promise<{ message: string; isCorrect: boolean }>;
    type: string; // 'year_range', 'generation', 'codename'
    label: string;
    points: number;
}

export default function BonusInputPopup({
    isOpen,
    onClose,
    onSubmit,
    type,
    label,
    points,
}: BonusInputPopupProps) {
    const [inputValue, setInputValue] = useState('');
    const [submitting, setSubmitting] = useState(false);
    const [feedback, setFeedback] = useState<{ message: string; isCorrect: boolean } | null>(null);
    const timerRef = useRef<NodeJS.Timeout | null>(null);

    useEffect(() => {
        return () => {
            if (timerRef.current) clearTimeout(timerRef.current);
        };
    }, []);

    // Filter feedback when open/closed
    const [secondsRemainingBeforeClose, setSecondsRemainingBeforeClose] = useState<number | null>(null);

    const handleClose = () => {
        if (timerRef.current) clearTimeout(timerRef.current);
        onClose();
    }

    useEffect(() => {
        if (!isOpen) return;
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === 'Escape') handleClose();
        };
        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [isOpen, onClose]);

    if (!isOpen) return null;

    const getPlaceholder = () => {
        switch (type) {
            case 'year_range': return 'ENTER YEAR RANGE (E.G. 2020-2024)';
            case 'generation': return 'ENTER GENERATION (E.G. MK8)';
            case 'codename': return 'ENTER CODENAME (E.G. G80)';
            default: return `ENTER ${label.toUpperCase()}`;
        }
    };

    const handleConfirm = async () => {
        if (!inputValue.trim() || submitting) return;
        setSubmitting(true);

        try {
            const result = await onSubmit(inputValue);
            setFeedback({ message: result.message, isCorrect: result.isCorrect });

            // Wait 5 seconds before closing so user can read the answer
            if (timerRef.current) clearTimeout(timerRef.current);
            timerRef.current = setTimeout(() => {
                handleClose();
                setFeedback(null);
                setInputValue('');
            }, 5000);
        } catch (error) {
            setFeedback({ message: 'Submission failed. Please try again.', isCorrect: false });
            if (timerRef.current) clearTimeout(timerRef.current);
            timerRef.current = setTimeout(() => setFeedback(null), 3000);
        } finally {
            setSubmitting(false);
        }
    };

    return (
        <div className="fixed inset-0 z-[60] flex items-center justify-center p-4">
            {/* Dark Overlay */}
            <div className="fixed inset-0 bg-black/80 z-40 backdrop-blur-sm" onClick={handleClose}></div>

            {/* Modal Body */}
            <div className="glass-panel w-full max-w-[440px] rounded-2xl overflow-hidden relative border-t-2 border-t-neon-cyan/50 z-50">
                <button
                    onClick={handleClose}
                    className="absolute top-4 right-4 text-white/40 hover:text-white transition-colors"
                >
                    <span className="material-symbols-outlined">close</span>
                </button>

                <div className="p-8 flex flex-col items-center">
                    <div className="mb-6 relative">
                        <div className="absolute inset-0 bg-neon-cyan blur-xl opacity-20"></div>
                        <div className="size-16 rounded-xl bg-neon-cyan/10 border border-neon-cyan/30 flex items-center justify-center relative z-10 rotate-45">
                            <span className="material-symbols-outlined text-neon-cyan text-3xl neon-glow -rotate-45">search_insights</span>
                        </div>
                    </div>

                    <h1 className="text-white tracking-[0.15em] text-xl font-bold leading-tight uppercase mb-2 text-center">
                        Identify {label}
                    </h1>

                    <p className="text-neon-cyan text-[10px] font-bold tracking-[0.2em] uppercase mb-8 opacity-80 text-center">
                        Precision bonus available: +{points} PTS
                    </p>

                    {/* Feedback Message */}
                    {feedback && (
                        <div className={`w-full mb-6 p-3 rounded-lg border text-center text-xs font-bold animate-in fade-in zoom-in duration-300 ${feedback.isCorrect
                            ? 'bg-green-500/10 border-green-500/50 text-green-400'
                            : 'bg-red-500/10 border-red-500/50 text-red-400'
                            }`}>
                            <div className="flex items-center justify-center gap-2">
                                <span className="material-symbols-outlined text-sm">
                                    {feedback.isCorrect ? 'check_circle' : 'error'}
                                </span>
                                {feedback.message}
                            </div>
                        </div>
                    )}

                    <div className="w-full space-y-6">
                        <div className="relative group">
                            <label className="absolute -top-2.5 left-4 px-2 bg-[#101622] text-[9px] font-bold text-neon-cyan tracking-widest uppercase z-10">
                                System Input
                            </label>
                            <input
                                autoFocus
                                className="w-full bg-[#192233]/60 border border-white/10 focus:border-neon-cyan focus:ring-1 focus:ring-neon-cyan/30 rounded-lg py-5 px-6 text-white text-lg font-bold placeholder:text-white/10 placeholder:text-[10px] placeholder:tracking-[0.1em] transition-all outline-none text-center tracking-widest uppercase neon-border-glow"
                                placeholder={getPlaceholder()}
                                type="text"
                                value={inputValue}
                                onChange={(e) => setInputValue(e.target.value)}
                                onKeyDown={(e) => e.key === 'Enter' && handleConfirm()}
                                disabled={submitting || !!feedback}
                            />
                        </div>

                        <div className="grid grid-cols-1 gap-4 pt-2">
                            <button
                                onClick={handleConfirm}
                                disabled={submitting || !inputValue.trim() || !!feedback}
                                className="group relative flex items-center justify-center gap-3 h-14 bg-primary hover:bg-primary/90 text-white rounded-lg font-extrabold transition-all shadow-[0_4px_20px_rgba(19,91,236,0.4)] text-sm tracking-[0.1em] overflow-hidden disabled:opacity-50"
                            >
                                <div className="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-700"></div>
                                <span className="uppercase">{submitting ? 'Verifying...' : (feedback?.isCorrect ? 'Mission Success' : 'Confirm Guess')}</span>
                                <span className="material-symbols-outlined text-xl">sensors</span>
                            </button>

                            <button
                                onClick={handleClose}
                                className="h-10 text-white/40 hover:text-white/70 text-[10px] font-bold uppercase tracking-[0.15em] transition-colors"
                            >
                                Return to Mission Summary
                            </button>
                        </div>
                    </div>

                    <div className="mt-8 w-full flex items-center justify-between opacity-20">
                        <div className="h-px flex-1 bg-gradient-to-r from-transparent to-white"></div>
                        <div className="px-3 flex gap-1">
                            <div className="size-1 bg-white rounded-full"></div>
                            <div className="size-1 bg-white rounded-full"></div>
                            <div className="size-1 bg-white rounded-full"></div>
                        </div>
                        <div className="h-px flex-1 bg-gradient-to-l from-transparent to-white"></div>
                    </div>
                </div>
            </div>
        </div>
    );
}

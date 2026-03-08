'use client';

interface StreakStats {
    attendance_streak: number;
    submission_streak: number;
    max_attendance_streak: number;
    max_submission_streak: number;
    total_days_participated: number;
    total_days_submitted: number;
}

interface StreakMeterProps {
    stats: StreakStats | null | undefined;
}

export default function StreakMeter({ stats }: StreakMeterProps) {
    if (!stats) return null;

    // Generate 7-dot progress indicator (mobile only)
    const renderDots = (current: number, color: string) => {
        const dots = [];
        for (let i = 0; i < 7; i++) {
            dots.push(
                <span
                    key={i}
                    className={`size-1.5 rounded-full transition-all ${i < current
                        ? `${color} shadow-[0_0_4px_currentColor]`
                        : 'bg-white/10'
                        }`}
                />
            );
        }
        return dots;
    };

    return (
        <>
            {/* ── Mobile: compact inline banner ── */}
            <div className="md:hidden flex items-center gap-3 px-3 py-2 bg-[#192233]/50 rounded-lg border border-white/5 backdrop-blur-sm">
                <span
                    className={`material-symbols-outlined text-lg leading-none ${stats.submission_streak > 0 ? 'text-orange-400' : 'text-white/15'
                        }`}
                    style={{ fontVariationSettings: "'FILL' 1" }}
                >
                    local_fire_department
                </span>

                <div className="flex items-center gap-2 min-w-0">
                    <span className="text-white text-sm font-black tabular-nums leading-none">
                        {stats.submission_streak}
                    </span>
                    <div className="flex items-center gap-0.5">
                        {renderDots(Math.min(stats.submission_streak, 7), 'bg-primary')}
                    </div>
                </div>

                <span className="w-px h-4 bg-white/10 shrink-0" />

                <div className="flex items-center gap-2 min-w-0">
                    <span className="material-symbols-outlined text-accent-cyan text-sm leading-none opacity-60">
                        visibility
                    </span>
                    <span className="text-white/70 text-sm font-bold tabular-nums leading-none">
                        {stats.attendance_streak}
                    </span>
                </div>

                <div className="flex-1" />

                <div className="flex items-center gap-1 px-1.5 py-0.5 bg-primary/10 rounded border border-primary/20 shrink-0">
                    <span className="material-symbols-outlined text-primary text-[10px] leading-none">star</span>
                    <span className="text-primary text-[9px] font-black uppercase leading-none tracking-tight">
                        {stats.max_submission_streak}
                    </span>
                </div>
            </div>

            {/* ── Desktop: full card layout ── */}
            <div className="hidden md:flex flex-col gap-4 p-4 bg-[#192233]/40 rounded-xl border border-white/5 backdrop-blur-sm">
                <div className="flex items-center justify-between">
                    <div className="flex items-center gap-2">
                        <span className="material-symbols-outlined text-primary text-xl">speed</span>
                        <h3 className="text-white text-xs font-bold uppercase tracking-widest">Active Streak</h3>
                    </div>
                    <div className="flex items-center gap-1.5 px-2 py-0.5 bg-primary/20 rounded-full border border-primary/30">
                        <span className="material-symbols-outlined text-primary text-xs leading-none">stars</span>
                        <span className="text-primary text-[10px] font-black uppercase tracking-tighter">PERSONAL BEST: {stats.max_submission_streak}</span>
                    </div>
                </div>

                <div className="grid grid-cols-2 gap-3">
                    {/* Submission Streak */}
                    <div className="flex flex-col p-3 bg-black/20 rounded-lg border border-white/5 relative overflow-hidden group">
                        <div className="absolute top-0 right-0 p-2 opacity-10 group-hover:opacity-20 transition-opacity">
                            <span className="material-symbols-outlined text-4xl -rotate-12">ads_click</span>
                        </div>
                        <span className="text-white/40 text-[9px] font-bold uppercase tracking-wider mb-1">Guess Streak</span>
                        <div className="flex items-baseline gap-2">
                            <span className="text-white text-2xl font-black tracking-tighter">{stats.submission_streak}</span>
                            <span className="text-primary text-[10px] font-bold uppercase">Days</span>
                        </div>
                        <div className="mt-2 h-1 w-full bg-white/5 rounded-full overflow-hidden">
                            <div
                                className="h-full bg-primary shadow-[0_0_8px_rgba(19,91,236,0.6)]"
                                style={{ width: `${Math.min((stats.submission_streak / 7) * 100, 100)}%` }}
                            ></div>
                        </div>
                    </div>

                    {/* Participation Streak */}
                    <div className="flex flex-col p-3 bg-black/20 rounded-lg border border-white/5 relative overflow-hidden group">
                        <div className="absolute top-0 right-0 p-2 opacity-10 group-hover:opacity-20 transition-opacity">
                            <span className="material-symbols-outlined text-4xl -rotate-12">garage</span>
                        </div>
                        <span className="text-white/40 text-[9px] font-bold uppercase tracking-wider mb-1">Daily Visits</span>
                        <div className="flex items-baseline gap-2">
                            <span className="text-white text-2xl font-black tracking-tighter">{stats.attendance_streak}</span>
                            <span className="text-accent-cyan text-[10px] font-bold uppercase">Days</span>
                        </div>
                        <div className="mt-2 h-1 w-full bg-white/5 rounded-full overflow-hidden">
                            <div
                                className="h-full bg-accent-cyan shadow-[0_0_8px_rgba(0,242,255,0.6)]"
                                style={{ width: `${Math.min((stats.attendance_streak / 7) * 100, 100)}%` }}
                            ></div>
                        </div>
                    </div>
                </div>

                <div className="flex items-center justify-between pt-1 opacity-60">
                    <div className="flex items-center gap-1">
                        <span className="text-[10px] font-bold text-white/50 uppercase tracking-widest">Total Spots:</span>
                        <span className="text-[10px] font-black text-primary uppercase">{stats.total_days_submitted}</span>
                    </div>
                    <div className="flex items-center gap-1">
                        <span className="text-[10px] font-bold text-white/50 uppercase tracking-widest">Garage Visits:</span>
                        <span className="text-[10px] font-black text-accent-cyan uppercase">{stats.total_days_participated}</span>
                    </div>
                </div>
            </div>
        </>
    );
}

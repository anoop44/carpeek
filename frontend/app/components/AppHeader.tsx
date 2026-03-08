'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { useState, useEffect, useRef } from 'react';

interface AppHeaderProps {
    streakCount?: number;
    onStreakClick?: () => void;
    onHelpClick?: () => void;
}

export default function AppHeader({ streakCount, onStreakClick, onHelpClick }: AppHeaderProps) {
    const pathname = usePathname();
    const [menuOpen, setMenuOpen] = useState(false);
    const drawerRef = useRef<HTMLDivElement>(null);

    const isActive = (path: string) => pathname === path;

    // Close on Escape key
    useEffect(() => {
        const handleKey = (e: KeyboardEvent) => {
            if (e.key === 'Escape') setMenuOpen(false);
        };
        document.addEventListener('keydown', handleKey);
        return () => document.removeEventListener('keydown', handleKey);
    }, []);

    // Prevent body scroll when drawer is open
    useEffect(() => {
        document.body.style.overflow = menuOpen ? 'hidden' : '';
        return () => { document.body.style.overflow = ''; };
    }, [menuOpen]);

    const closeMenu = () => setMenuOpen(false);

    return (
        <>
            <header className="border-b border-white/10 bg-background-light dark:bg-background-dark/80 sticky top-0 z-50 backdrop-blur-md">
                <div className="max-w-[1200px] mx-auto h-16 flex items-center justify-between px-4 md:px-8">
                    {/* Logo */}
                    <Link href="/" className="flex items-center gap-3" onClick={closeMenu}>
                        <div className="p-2 bg-primary rounded-lg text-white">
                            <span className="material-symbols-outlined block">electric_car</span>
                        </div>
                        <h2 className="text-lg md:text-xl font-bold tracking-tight uppercase">AutoCorrect</h2>
                    </Link>

                    {/* ── Desktop right side (md and above) ── */}
                    <div className="hidden md:flex items-center gap-6">
                        <nav className="flex items-center gap-6">
                            <Link
                                className={`text-sm font-medium transition-colors ${isActive('/') ? 'text-white border-b-2 border-primary pb-0.5' : 'hover:text-primary'}`}
                                href="/"
                            >
                                Race
                            </Link>
                            <Link
                                className={`text-sm font-medium transition-colors ${isActive('/leaderboard') ? 'text-white border-b-2 border-primary pb-0.5' : 'hover:text-primary'}`}
                                href="/leaderboard"
                            >
                                Leaderboard
                            </Link>
                        </nav>
                        {streakCount !== undefined && (
                            <button
                                onClick={onStreakClick}
                                className="flex items-center gap-2 px-3 py-1.5 bg-primary/10 border border-primary/20 rounded-full hover:bg-primary/20 hover:border-primary/40 transition-colors cursor-pointer group"
                            >
                                <span
                                    className={`material-symbols-outlined text-sm ${streakCount > 0 ? 'text-orange-500 animate-pulse' : 'text-white/20'}`}
                                    style={{ fontVariationSettings: "'FILL' 1" }}
                                >
                                    local_fire_department
                                </span>
                                <span className="text-sm font-bold uppercase tracking-tighter group-hover:text-white transition-colors">
                                    {streakCount} Day Streak
                                </span>
                            </button>
                        )}
                        {onHelpClick && (
                            <button
                                onClick={onHelpClick}
                                className="w-10 h-10 flex items-center justify-center rounded-full bg-slate-800 hover:bg-slate-700 text-slate-300 hover:text-white transition-colors border border-white/10"
                                title="How to Play"
                            >
                                <span className="material-symbols-outlined text-xl">help</span>
                            </button>
                        )}
                    </div>

                    {/* ── Mobile right side (below md) ── */}
                    <div className="flex md:hidden items-center gap-2">
                        {/* Compact streak pill — always visible on mobile when there's a streak */}
                        {streakCount !== undefined && streakCount > 0 && (
                            <button
                                onClick={onStreakClick}
                                aria-label={`${streakCount} day streak`}
                                className="flex items-center gap-1.5 px-2.5 py-1 bg-primary/10 border border-primary/20 rounded-full transition-colors active:scale-95"
                            >
                                <span
                                    className="material-symbols-outlined text-orange-500 animate-pulse"
                                    style={{ fontVariationSettings: "'FILL' 1", fontSize: '16px' }}
                                >
                                    local_fire_department
                                </span>
                                <span className="text-xs font-bold uppercase tracking-tight">
                                    {streakCount}
                                </span>
                            </button>
                        )}

                        {/* Hamburger toggle */}
                        <button
                            id="mobile-menu-button"
                            onClick={() => setMenuOpen(prev => !prev)}
                            aria-label={menuOpen ? 'Close menu' : 'Open menu'}
                            aria-expanded={menuOpen}
                            aria-controls="mobile-menu-drawer"
                            className="w-10 h-10 flex items-center justify-center rounded-xl bg-slate-800/80 hover:bg-slate-700 text-slate-300 hover:text-white transition-all border border-white/10 active:scale-95"
                        >
                            <span
                                className="material-symbols-outlined transition-all duration-200"
                                style={{ fontSize: '22px' }}
                            >
                                {menuOpen ? 'close' : 'menu'}
                            </span>
                        </button>
                    </div>
                </div>
            </header>

            {/* ── Mobile Drawer Backdrop ── */}
            <div
                className={`fixed inset-0 z-40 bg-black/60 backdrop-blur-sm transition-opacity duration-300 md:hidden ${menuOpen ? 'opacity-100 pointer-events-auto' : 'opacity-0 pointer-events-none'
                    }`}
                onClick={closeMenu}
                aria-hidden="true"
            />

            {/* ── Mobile Drawer Panel ── */}
            <div
                ref={drawerRef}
                id="mobile-menu-drawer"
                role="dialog"
                aria-modal="true"
                aria-label="Navigation menu"
                className={`fixed top-16 left-0 right-0 z-40 md:hidden transition-all duration-300 ease-out origin-top ${menuOpen
                        ? 'opacity-100 scale-y-100 translate-y-0'
                        : 'opacity-0 scale-y-95 -translate-y-2 pointer-events-none'
                    }`}
            >
                <div className="mx-3 mt-1 rounded-2xl border border-white/10 bg-background-dark/95 backdrop-blur-xl shadow-2xl overflow-hidden">

                    {/* Nav links */}
                    <nav className="flex flex-col p-3 gap-1" aria-label="Mobile navigation">
                        <Link
                            href="/"
                            onClick={closeMenu}
                            className={`flex items-center gap-3 px-4 py-3.5 rounded-xl text-sm font-bold uppercase tracking-widest transition-colors ${isActive('/')
                                    ? 'bg-primary/20 text-primary border border-primary/30'
                                    : 'text-slate-300 hover:bg-white/5 hover:text-white'
                                }`}
                        >
                            <span
                                className="material-symbols-outlined"
                                style={{ fontVariationSettings: isActive('/') ? "'FILL' 1" : "'FILL' 0", fontSize: '20px' }}
                            >
                                sports_motorsports
                            </span>
                            Race
                            {isActive('/') && (
                                <span className="ml-auto w-1.5 h-1.5 rounded-full bg-primary animate-pulse" />
                            )}
                        </Link>

                        <Link
                            href="/leaderboard"
                            onClick={closeMenu}
                            className={`flex items-center gap-3 px-4 py-3.5 rounded-xl text-sm font-bold uppercase tracking-widest transition-colors ${isActive('/leaderboard')
                                    ? 'bg-primary/20 text-primary border border-primary/30'
                                    : 'text-slate-300 hover:bg-white/5 hover:text-white'
                                }`}
                        >
                            <span
                                className="material-symbols-outlined"
                                style={{ fontVariationSettings: isActive('/leaderboard') ? "'FILL' 1" : "'FILL' 0", fontSize: '20px' }}
                            >
                                leaderboard
                            </span>
                            Leaderboard
                            {isActive('/leaderboard') && (
                                <span className="ml-auto w-1.5 h-1.5 rounded-full bg-primary animate-pulse" />
                            )}
                        </Link>

                        {onHelpClick && (
                            <button
                                onClick={() => { onHelpClick(); closeMenu(); }}
                                className="flex items-center gap-3 px-4 py-3.5 rounded-xl text-sm font-bold uppercase tracking-widest text-slate-300 hover:bg-white/5 hover:text-white transition-colors w-full text-left"
                            >
                                <span className="material-symbols-outlined" style={{ fontSize: '20px' }}>
                                    help
                                </span>
                                How to Play
                            </button>
                        )}
                    </nav>

                    {/* Streak card — only shown when streak prop is provided */}
                    {streakCount !== undefined && (
                        <>
                            <div className="mx-4 border-t border-white/[0.06]" />
                            <div className="p-3">
                                <button
                                    onClick={() => { onStreakClick?.(); closeMenu(); }}
                                    className="w-full flex items-center gap-3 px-4 py-3 rounded-xl bg-orange-500/5 border border-orange-500/20 hover:bg-orange-500/10 transition-colors"
                                >
                                    <span
                                        className={`material-symbols-outlined ${streakCount > 0 ? 'text-orange-500 animate-pulse' : 'text-slate-600'}`}
                                        style={{ fontVariationSettings: "'FILL' 1", fontSize: '22px' }}
                                    >
                                        local_fire_department
                                    </span>
                                    <div className="text-left">
                                        <p className="text-xs font-black uppercase tracking-widest text-white/80">
                                            {streakCount} Day Streak
                                        </p>
                                        <p className="text-[10px] text-slate-500 font-medium uppercase tracking-wider">
                                            Tap to view details
                                        </p>
                                    </div>
                                    <span className="ml-auto material-symbols-outlined text-slate-600" style={{ fontSize: '16px' }}>
                                        chevron_right
                                    </span>
                                </button>
                            </div>
                        </>
                    )}
                </div>
            </div>
        </>
    );
}

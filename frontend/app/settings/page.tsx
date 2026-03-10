'use client';

import { useAuth } from '../components/AuthProvider';
import AppHeader from '../components/AppHeader';
import Link from 'next/link';

export default function SettingsPage() {
    const { isLoggedIn, isSubscriber, userEmail, subscriptionProvider, customerInfo, isLoading, logout } = useAuth();

    return (
        <div className="relative flex h-auto min-h-screen w-full flex-col group/design-root overflow-x-hidden scanline bg-background-light dark:bg-background-dark text-slate-900 dark:text-slate-100">
            <AppHeader />
            
            <main className="flex-1 flex flex-col items-center py-12 px-6 max-w-2xl mx-auto w-full">
                <h1 className="text-white text-3xl font-bold font-display uppercase tracking-widest mb-10 text-center">
                    Settings
                </h1>

                {isLoading ? (
                    <div className="py-20 flex flex-col items-center justify-center">
                        <div className="size-8 border-4 border-primary/30 border-t-primary rounded-full animate-spin mb-4"></div>
                        <p className="text-slate-400 uppercase tracking-widest text-xs font-bold">Loading account...</p>
                    </div>
                ) : (
                    <div className="w-full flex justify-between bg-card-dark border border-primary/20 p-6 rounded-2xl flex-col gap-8 shadow-lg shadow-black/20 relative overflow-hidden">
                        
                        {/* Account Section */}
                        <section className="flex flex-col gap-4 relative z-10 w-full">
                            <h2 className="text-lg font-bold text-white uppercase tracking-widest flex items-center gap-2 border-b border-white/10 pb-2">
                                <span className="material-symbols-outlined text-primary">account_circle</span>
                                Account
                            </h2>
                            
                            {isLoggedIn ? (
                                <div className="flex flex-col md:flex-row md:items-center justify-between gap-4 bg-white/[0.02] p-4 rounded-xl">
                                    <div className="flex flex-col">
                                        <p className="text-slate-400 text-xs font-bold uppercase tracking-wider mb-1">Signed in as</p>
                                        <p className="text-white font-medium">{userEmail || 'Unknown Email'}</p>
                                    </div>
                                    <button 
                                        onClick={logout}
                                        className="px-4 py-2 rounded-lg bg-red-500/10 text-red-400 border border-red-500/30 hover:bg-red-500/20 hover:text-red-300 font-bold uppercase tracking-wider text-xs transition-colors self-start md:self-auto"
                                    >
                                        Sign Out
                                    </button>
                                </div>
                            ) : (
                                <div className="flex flex-col items-start gap-3 bg-white/[0.02] p-4 rounded-xl">
                                    <p className="text-slate-400 text-sm">You are playing as an <span className="text-white font-bold">Anonymous Spotter</span>.</p>
                                    <p className="text-slate-500 text-xs">Sign in from the Leaderboard page to sync your progress.</p>
                                    <Link href="/leaderboard" className="mt-2 text-primary font-bold uppercase tracking-widest text-xs hover:underline flex items-center gap-1">
                                        Go to Leaderboard <span className="material-symbols-outlined text-[14px]">arrow_forward</span>
                                    </Link>
                                </div>
                            )}
                        </section>

                        {/* Subscription Section */}
                        <section className="flex flex-col gap-4 relative z-10 w-full animate-fade-in delay-100 fill-mode-both">
                            <h2 className="text-lg font-bold text-white uppercase tracking-widest flex items-center gap-2 border-b border-white/10 pb-2">
                                <span className="material-symbols-outlined text-accent-neon">workspace_premium</span>
                                Subscription
                            </h2>
                            
                            {isSubscriber ? (
                                <div className="flex flex-col gap-4 bg-gradient-to-br from-green-500/10 to-transparent p-5 rounded-xl border border-green-500/20">
                                    <div className="flex items-center gap-3">
                                        <div className="size-10 rounded-full bg-green-500/20 text-green-400 flex items-center justify-center">
                                            <span className="material-symbols-outlined">verified</span>
                                        </div>
                                        <div className="flex-1">
                                            <h3 className="text-white font-bold text-lg">AutoCorrect Pro</h3>
                                            <p className="text-green-400/80 text-xs font-bold uppercase tracking-wider">Active via {subscriptionProvider === 'razorpay' ? 'Razorpay' : 'RevenueCat'}</p>
                                        </div>
                                    </div>
                                    
                                    <div className="mt-2 pt-4 border-t border-green-500/10 flex flex-col md:flex-row md:items-center justify-between gap-4">
                                        {subscriptionProvider === 'razorpay' ? (
                                            <p className="text-slate-400 text-xs">
                                                To manage or cancel your subscription, please check your email from Razorpay or use the Razorpay mobile app/website.
                                            </p>
                                        ) : customerInfo?.managementURL ? (
                                            <a 
                                                href={customerInfo.managementURL}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="px-5 py-2.5 rounded-lg bg-green-500 hover:bg-green-400 text-white font-bold uppercase tracking-wider text-xs transition-colors text-center inline-flex items-center justify-center gap-2"
                                            >
                                                Manage Billing <span className="material-symbols-outlined text-[16px]">open_in_new</span>
                                            </a>
                                        ) : (
                                            <p className="text-slate-400 text-xs">
                                                Manage your subscription via the Paddle/RevenueCat portal sent to your email.
                                            </p>
                                        )}
                                        <p className="text-slate-500 text-xs text-center md:text-left md:max-w-[200px]">
                                            Thank you for supporting AutoCorrect!
                                        </p>
                                    </div>
                                </div>
                            ) : (
                                <div className="flex flex-col md:flex-row items-center justify-between gap-6 bg-white/[0.02] p-5 rounded-xl border border-white/5">
                                    <div className="flex flex-col items-center md:items-start text-center md:text-left flex-1">
                                        <h3 className="text-white font-bold mb-1">Remove Ads</h3>
                                        <p className="text-slate-400 text-sm">Enjoy an uninterrupted experience and support the app.</p>
                                    </div>
                                    <Link 
                                        href="/subscribe"
                                        className="px-6 py-3 rounded-xl bg-gradient-to-r from-primary to-accent-neon hover:to-primary hover:from-accent-neon text-white font-bold uppercase tracking-widest text-xs transition-all w-full md:w-auto text-center glow-[primary]"
                                    >
                                        Subscribe Now
                                    </Link>
                                </div>
                            )}
                        </section>

                    </div>
                )}
            </main>
        </div>
    );
}

import AppHeader from '../components/AppHeader'
import Link from 'next/link'
import { Metadata } from 'next'

export const metadata: Metadata = {
    title: 'Circuit Rules | AutoCorrect Terms',
    description: 'The terms of play and service rules for AutoCorrect by Pixel Fairy Studio.',
}

export default function TermsPage() {
    return (
        <div className="min-h-screen bg-background-dark text-slate-100 flex flex-col">
            <AppHeader />
            <main className="flex-1 max-w-4xl mx-auto px-6 py-16 text-slate-400 font-medium leading-relaxed">
                <div className="mb-12">
                    <h1 className="text-4xl md:text-5xl font-black tracking-tight uppercase mb-4 font-display text-white">
                        Terms & Conditions
                    </h1>
                    <p className="text-primary text-xs font-bold tracking-[0.4em] uppercase">
                        Circuit Rules — Version 1.0 — March 2026
                    </p>
                </div>

                <div className="space-y-12">
                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">1. Acceptance of Terms</h2>
                        <p>
                            By accessing AutoCorrect (&quot;App&quot;), a service provided by Pixel Fairy Studio, you agree to comply with and be bound by these Terms and Conditions (&quot;Terms&quot;). If you do not agree, please do not use the service.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">2. User Account and Data</h2>
                        <p>
                            You represent that any information provided via Google OAuth is accurate. You are responsible for maintaining the privacy of your session and are accountable for all activities conducted under your identity on the leaderboard.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">3. Fair Play Protocol</h2>
                        <div className="space-y-4">
                            <p>To maintain a fair competitive environment, users agree to:</p>
                            <ul className="list-disc pl-5 space-y-2">
                                <li>Not use any automated bots or scripts to guess challenges.</li>
                                <li>Not attempt to reverse engineer the game logic or bypass server-side verification.</li>
                                <li>Not exploit bugs or technical glitches to manipulate the leaderboard rankings.</li>
                            </ul>
                            <p>We reserve the right to remove entries that violate these rules.</p>
                        </div>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">4. Intellectual Property</h2>
                        <p>
                            AutoCorrect, its logo, design, codebase, and unique car challenges are the intellectual property of Pixel Fairy Studio. You may not reproduce, copy, or distribute any part of the service without explicit permission.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">5. Limitation of Liability</h2>
                        <p>
                            AutoCorrect is provided on an &quot;as is&quot; and &quot;as available&quot; basis. Pixel Fairy Studio disclaims all warranties, either express or implied. We satisfy no guarantee that the service will be error-free or uninterrupted.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">6. Modification of Terms</h2>
                        <p>
                            We reserve the right to modify these Terms at any time. Changes will be effective upon posting update versions to the App. Your continued use of AutoCorrect signifies your acceptance of updated terms.
                        </p>
                    </section>
                </div>

                <div className="mt-16 pt-8 border-t border-white/5 flex justify-center">
                    <Link href="/" className="text-xs font-bold uppercase tracking-widest hover:text-primary transition-colors">
                        Return to Race
                    </Link>
                </div>
            </main>
        </div>
    )
}

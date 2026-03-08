import AppHeader from '../components/AppHeader'
import Link from 'next/link'
import { Metadata } from 'next'

export const metadata: Metadata = {
    title: 'Privacy Protocol | AutoCorrect',
    description: 'How we protect your data and privacy at Pixel Fairy Studio.',
}

export default function PrivacyPage() {
    return (
        <div className="min-h-screen bg-background-dark text-slate-100 flex flex-col">
            <AppHeader />
            <main className="flex-1 max-w-4xl mx-auto px-6 py-16">
                <div className="mb-12">
                    <h1 className="text-4xl md:text-5xl font-black tracking-tight uppercase mb-4 font-display">
                        Privacy Policy
                    </h1>
                    <p className="text-primary text-xs font-bold tracking-[0.4em] uppercase">
                        Protocol Version 1.0 — Last Updated March 2026
                    </p>
                </div>

                <div className="space-y-12 text-slate-400 font-medium leading-relaxed">
                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">1. Overview</h2>
                        <p>
                            AutoCorrect (&quot;we&quot;, &quot;us&quot;, or &quot;our&quot;), a service provided by Pixel Fairy Studio, is committed to protecting your privacy. This Privacy Policy explains how we collect, use, and safeguard your information when you use our car identification challenge app.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">2. Information Collection</h2>
                        <div className="space-y-4">
                            <p>We collect information in the following ways:</p>
                            <ul className="list-disc pl-5 space-y-2">
                                <li><span className="text-white">Google OAuth:</span> When you sign in with Google, we receive your name, email address, and profile picture URL. This is used solely to display your identity on the global leaderboard.</li>
                                <li><span className="text-white">Anonymous Identifiers:</span> For users who do not sign in, we generate a unique browser signature to track your daily progress and scores locally.</li>
                                <li><span className="text-white">Game Data:</span> We store your challenge attempts, scores, and streak data to provide the game service.</li>
                            </ul>
                        </div>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">3. Cookies & Advertisements</h2>
                        <p>
                            We use &quot;cookies&quot; to manage user sessions and for Google AdSense to serve relevant advertisements. These cookies help us understand how users interact with the site and help fund the development of AutoCorrect. By using the site, you consent to the use of cookies.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">4. Data Security</h2>
                        <p>
                            We implement industry-standard security protocols to protect your data. Your Google credentials are never stored directly on our servers; we only store the public tokens and profile information provided by the Google OAuth service.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">5. Third-Party Services</h2>
                        <p>
                            Our app integrates with Google services. Please review Google&apos;s Privacy Policy regarding how they handle data across their platforms.
                        </p>
                    </section>

                    <section>
                        <h2 className="text-xl font-bold text-white uppercase tracking-widest mb-4">6. Contact</h2>
                        <p>
                            For any questions regarding this protocol, please reach out to us at <span className="text-primary">support@pixelfairystudio.com</span>.
                        </p>
                    </section>
                </div>

                <div className="mt-16 pt-8 border-t border-white/5 flex justify-center">
                    <Link href="/" className="text-xs font-bold uppercase tracking-widest hover:text-primary transition-colors">
                        Back to Circuit
                    </Link>
                </div>
            </main>
        </div>
    )
}

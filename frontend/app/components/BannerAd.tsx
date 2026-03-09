'use client';

import { useEffect, useRef } from 'react';
import { useAuth } from './AuthProvider';
import Link from 'next/link';

interface BannerAdProps {
    /**
     * The AdSense ad slot ID. Get this from your AdSense account.
     * Example: "1234567890"
     */
    adSlot: string;
    /**
     * Optional label shown above the ad. Defaults to "Advertisement".
     * Pass an empty string to hide it.
     */
    label?: string;
    /**
     * Extra CSS class names for the outer wrapper (for placement-specific tweaks).
     */
    className?: string;
}

/**
 * BannerAd — Google AdSense display ad.
 *
 * The ad blends into the dark UI by using a muted container that matches
 * the card/surface colours of the rest of the app.
 *
 * REQUIRES:
 *   - NEXT_PUBLIC_ADSENSE_CLIENT_ID set in your .env (e.g. "ca-pub-XXXXXXXXXXXXXXXXX")
 *   - The AdSense <script> tag added to app/layout.tsx (see ADSENSE_SETUP.md)
 */
export default function BannerAd({
    adSlot,
    label = 'Advertisement',
    className = '',
}: BannerAdProps) {
    const adRef = useRef<HTMLModElement>(null);
    const clientId = process.env.NEXT_PUBLIC_ADSENSE_CLIENT_ID;
    const { isSubscriber, isLoading } = useAuth();

    useEffect(() => {
        // Only run on client when the global adsbygoogle is available and user is NOT a subscriber
        if (typeof window === 'undefined' || !clientId || isLoading || isSubscriber) {
            return;
        }

        try {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            ((window as any).adsbygoogle = (window as any).adsbygoogle || []).push({});
        } catch (e) {
            console.error('[BannerAd] adsbygoogle push error:', e);
        }
    }, [clientId, isSubscriber, isLoading]);

    // If no client ID configured, or if waiting for sub status, or if user is subscriber, render nothing
    if (!clientId || isLoading || isSubscriber) return null;

    return (
        <div
            className={`w-full flex flex-col items-center gap-1.5 ${className}`}
            aria-label="Sponsored content"
        >
            <div className="flex w-full max-w-[970px] justify-between items-end px-2">
                {label && (
                    <p className="text-[9px] font-bold uppercase tracking-[0.25em] text-slate-500 select-none">
                        {label}
                    </p>
                )}
                <Link 
                    href="/subscribe" 
                    className="text-[10px] text-primary/80 hover:text-primary font-bold uppercase tracking-wider transition-colors drop-shadow-sm flex items-center gap-1"
                >
                    Hide Ads
                    <span className="material-symbols-outlined text-[10px]" style={{ fontVariationSettings: "'FILL' 1" }}>workspace_premium</span>
                </Link>
            </div>
            
            {/* Muted surface that blends with the page background */}
            <div className="w-full max-w-[970px] rounded-xl overflow-hidden bg-card-dark/30 border border-white/[0.04] backdrop-blur-sm">
                <ins
                    ref={adRef}
                    className="adsbygoogle block"
                    style={{ display: 'block' }}
                    data-ad-client={clientId}
                    data-ad-slot={adSlot}
                    data-ad-format="auto"
                    data-full-width-responsive="true"
                />
            </div>
        </div>
    );
}

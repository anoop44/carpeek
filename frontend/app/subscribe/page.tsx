'use client';

import { useEffect, useState } from 'react';
import AppHeader from '../components/AppHeader';
import { useAuth } from '../components/AuthProvider';
import GoogleLoginButton from '../components/GoogleLoginButton';
import { Purchases, Package } from '@revenuecat/purchases-js';
import Link from 'next/link';
import { isInIndia } from '../utils/region';
import { loadRazorpay } from '../utils/razorpay';

export default function SubscribePage() {
    const { isLoggedIn, isSubscriber, isLoading } = useAuth();
    const [packages, setPackages] = useState<Package[]>([]);
    const [fetchingOfferings, setFetchingOfferings] = useState(true);
    const [purchaseError, setPurchaseError] = useState<string | null>(null);
    const [isPurchasing, setIsPurchasing] = useState(false);
    const [isIndianUser, setIsIndianUser] = useState<boolean | null>(null);

    useEffect(() => {
        // Step 1: Detect Region
        const inIndia = isInIndia();
        setIsIndianUser(inIndia);

        // Step 2: Fetch Offerings or Load Razorpay
        const initialize = async () => {
            if (typeof window === 'undefined' || isLoading) return;
            
            if (!inIndia) {
                // International: Load RevenueCat
                try {
                    setFetchingOfferings(true);
                    const offerings = await Purchases.getSharedInstance().getOfferings();
                    
                    if (offerings.current !== null && offerings.current.availablePackages.length > 0) {
                        setPackages(offerings.current.availablePackages);
                    } else {
                        console.log("No current offerings available in RevenueCat");
                    }
                } catch (e) {
                    console.error("Error fetching RC offerings", e);
                    setPurchaseError("Failed to load subscription options.");
                } finally {
                    setFetchingOfferings(false);
                }
            } else {
                // India: Pre-load Razorpay script
                await loadRazorpay();
                setFetchingOfferings(false);
            }
        };

        initialize();
    }, [isLoading]);

    const handleRCPurchase = async (pkg: Package) => {
        if (!isLoggedIn) {
            setPurchaseError("Please sign in first to subscribe.");
            return;
        }

        try {
            setIsPurchasing(true);
            setPurchaseError(null);
            
            const purchaseResult = await Purchases.getSharedInstance().purchasePackage(pkg);
            
            const entitlementId = process.env.NEXT_PUBLIC_REVENUECAT_ENTITLEMENT || 'ad_free';
            if (entitlementId in purchaseResult.customerInfo.entitlements.active) {
                window.location.href = '/settings';
            } else {
                setPurchaseError("Purchase completed but subscription not activated. Please contact support.");
            }

        } catch (e: unknown) {
            const error = e as { userCancelled?: boolean; message?: string };
            if (!error.userCancelled) {
                console.error("Purchase error", e);
                setPurchaseError(error.message || "An error occurred during purchase.");
            }
        } finally {
            setIsPurchasing(false);
        }
    };

    const handleRazorpayPurchase = async () => {
        if (!isLoggedIn) {
            setPurchaseError("Please sign in first to subscribe.");
            return;
        }

        try {
            setIsPurchasing(true);
            setPurchaseError(null);
            
            // 1. Create Subscription in Backend
            const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'}/api/v1/subscription/razorpay/create`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error("Failed to initialize payment service.");
            }

            const { subscription_id, key_id } = await response.json();

            // 2. Open Razorpay Checkout
            const options = {
                key: key_id,
                subscription_id: subscription_id,
                name: "AutoCorrect Pro",
                description: "Monthly Ad-Free Subscription",
                image: "/images/icon.png",
                handler: function(response: any) {
                    // Success! Razorpay takes care of it, we just redirect or show success
                    window.location.href = '/settings?payment=success';
                },
                prefill: {
                  email: JSON.parse(localStorage.getItem('user_info') || '{}').email || '',
                },
                theme: {
                    color: "#0d5bec"
                }
            };

            const rzp = new (window as any).Razorpay(options);
            rzp.on('payment.failed', function (response: any) {
                setPurchaseError(response.error.description || "Payment failed.");
            });
            rzp.open();

        } catch (e: any) {
            console.error("Razorpay error", e);
            setPurchaseError(e.message || "Failed to start payment process.");
        } finally {
            setIsPurchasing(false);
        }
    };

    return (
        <div className="relative flex h-auto min-h-screen w-full flex-col group/design-root overflow-x-hidden scanline bg-background-light dark:bg-background-dark text-slate-900 dark:text-slate-100">
            <AppHeader />
            
            <main className="flex-1 flex flex-col items-center py-12 px-6 max-w-4xl mx-auto w-full">
                <div className="text-center mb-10">
                    <h1 className="text-white text-4xl md:text-5xl font-bold font-display uppercase tracking-widest mb-4">
                        Upgrade Your Experience
                    </h1>
                    <p className="text-slate-400 text-lg max-w-2xl mx-auto">
                        Support AutoCorrect and enjoy an uninterrupted, ad-free spotting experience.
                    </p>
                </div>

                {/* Benefits Section */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full mb-12">
                    <div className="bg-card-dark border border-primary/20 rounded-2xl p-6 flex flex-col items-start gap-4 shadow-lg shadow-primary/5">
                        <div className="size-12 rounded-xl bg-green-500/20 text-green-400 flex items-center justify-center">
                            <span className="material-symbols-outlined text-2xl" style={{ fontVariationSettings: "'FILL' 1" }}>block</span>
                        </div>
                        <div>
                            <h3 className="text-white font-bold text-xl mb-1">Zero Ads</h3>
                            <p className="text-slate-400 text-sm">Focus entirely on the challenge without any distractions. No banners, no popups.</p>
                        </div>
                    </div>
                    
                    <div className="bg-card-dark border border-primary/20 rounded-2xl p-6 flex flex-col items-start gap-4 shadow-lg shadow-primary/5">
                        <div className="size-12 rounded-xl bg-orange-500/20 text-orange-400 flex items-center justify-center">
                            <span className="material-symbols-outlined text-2xl" style={{ fontVariationSettings: "'FILL' 1" }}>favorite</span>
                        </div>
                        <div>
                            <h3 className="text-white font-bold text-xl mb-1">Support Development</h3>
                            <p className="text-slate-400 text-sm">Your subscription directly helps keep servers running and funds the creation of new features.</p>
                        </div>
                    </div>
                </div>

                {/* Actions Section */}
                <div className="w-full max-w-md bg-card-dark border border-primary/30 rounded-3xl p-8 flex flex-col items-center relative overflow-hidden">
                    <div className="absolute top-0 right-0 w-64 h-64 bg-primary/10 rounded-full blur-3xl -mr-32 -mt-32 pointer-events-none"></div>

                    {isSubscriber ? (
                        <div className="flex flex-col items-center text-center">
                            <div className="size-16 rounded-full bg-green-500/20 text-green-400 flex items-center justify-center mb-4">
                                <span className="material-symbols-outlined text-3xl" style={{ fontVariationSettings: "'FILL' 1" }}>check_circle</span>
                            </div>
                            <h2 className="text-2xl font-bold text-white mb-2">You are Ad-Free!</h2>
                            <p className="text-slate-400 mb-6">Thank you for supporting AutoCorrect.</p>
                            <Link href="/settings" className="px-6 py-3 rounded-full bg-white/10 text-white font-bold hover:bg-white/20 transition-colors uppercase tracking-widest text-sm">
                                Manage Subscription
                            </Link>
                        </div>
                    ) : isLoading || isIndianUser === null || fetchingOfferings ? (
                         <div className="py-12 flex flex-col items-center justify-center">
                             <div className="size-8 border-4 border-primary/30 border-t-primary rounded-full animate-spin mb-4"></div>
                             <p className="text-slate-400 uppercase tracking-widest text-xs font-bold">Loading plans...</p>
                         </div>
                    ) : !isLoggedIn ? (
                        <div className="flex flex-col items-center text-center">
                            <div className="size-16 rounded-full bg-primary/20 text-primary flex items-center justify-center mb-6">
                                <span className="material-symbols-outlined text-3xl" style={{ fontVariationSettings: "'FILL' 1" }}>person</span>
                            </div>
                            <h2 className="text-2xl font-bold text-white mb-3">Sign in to Subscribe</h2>
                            <p className="text-slate-400 text-sm mb-6">
                                Please sign in with Google first. This ensures you can access your subscription across all your devices.
                            </p>
                            <GoogleLoginButton />
                        </div>
                    ) : (
                        <div className="w-full flex flex-col items-center">
                            <h2 className="text-2xl font-bold text-white mb-6">Choose a Plan</h2>
                            
                            {purchaseError && (
                                <div className="w-full bg-red-500/10 border border-red-500/30 text-red-400 text-sm p-3 rounded-lg mb-6 text-center">
                                    {purchaseError}
                                </div>
                            )}

                            {isIndianUser ? (
                                <div className="w-full flex flex-col gap-4">
                                    <button
                                        onClick={handleRazorpayPurchase}
                                        disabled={isPurchasing}
                                        className="w-full relative group overflow-hidden rounded-2xl border border-primary hover:border-accent-neon transition-colors"
                                    >
                                        <div className="absolute inset-0 bg-gradient-to-br from-primary/10 to-transparent group-hover:from-accent-neon/20 transition-colors"></div>
                                        <div className="relative p-6 flex flex-col">
                                            <div className="flex justify-between items-end mb-2">
                                                 <h3 className="text-white font-bold text-lg">AutoCorrect Pro</h3>
                                                 <p className="text-2xl font-bold text-white font-display">
                                                     ₹99
                                                     <span className="text-sm text-slate-400 font-sans ml-1 font-normal uppercase tracking-wider">
                                                         / mo
                                                     </span>
                                                 </p>
                                            </div>
                                            <p className="text-slate-400 text-sm text-left">
                                                Go ad-free with UPI Autopay. Standard monthly renewal.
                                            </p>
                                        </div>
                                    </button>
                                    
                                    <div className="flex flex-wrap justify-center gap-4 mt-2 mb-2 opacity-60 grayscale hover:grayscale-0 transition-all text-white/50 text-[10px] uppercase tracking-tighter">
                                        <span>UPI</span> • <span>Cards</span> • <span>Netbanking</span>
                                    </div>

                                    <p className="text-slate-500 text-xs text-center mt-2">
                                        Cancel anytime. Secure payments processed by Razorpay.
                                    </p>
                                </div>
                            ) : (
                                <div className="w-full flex flex-col gap-4">
                                    {packages.length === 0 ? (
                                        <div className="text-slate-400 text-center py-6">
                                            No subscription plans available in your region.
                                        </div>
                                    ) : packages.map((pkg) => (
                                        <button
                                            key={pkg.identifier}
                                            onClick={() => handleRCPurchase(pkg)}
                                            disabled={isPurchasing}
                                            className="w-full relative group overflow-hidden rounded-2xl border border-primary hover:border-accent-neon transition-colors"
                                        >
                                            <div className="absolute inset-0 bg-gradient-to-br from-primary/10 to-transparent group-hover:from-accent-neon/20 transition-colors"></div>
                                            <div className="relative p-6 flex flex-col">
                                                <div className="flex justify-between items-end mb-2">
                                                     <h3 className="text-white font-bold text-lg">{pkg.rcBillingProduct.displayName || pkg.identifier}</h3>
                                                     <p className="text-2xl font-bold text-white font-display">
                                                         {pkg.rcBillingProduct.currentPrice.formattedPrice}
                                                         <span className="text-sm text-slate-400 font-sans ml-1 font-normal uppercase tracking-wider">
                                                             / yr
                                                         </span>
                                                     </p>
                                                </div>
                                                <p className="text-slate-400 text-sm text-left line-clamp-2">
                                                    {pkg.rcBillingProduct.description || 'Remove all ads and support the app.'}
                                                </p>
                                            </div>
                                        </button>
                                    ))}
                                    
                                    <p className="text-slate-500 text-xs text-center mt-4">
                                        Cancel anytime. Secure payments processed through Paddle.
                                    </p>
                                </div>
                            )}
                        </div>
                    )}
                </div>
            </main>
        </div>
    );
}

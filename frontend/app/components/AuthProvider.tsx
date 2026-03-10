'use client';

import React, { createContext, useContext, useState, useEffect, ReactNode, useRef } from 'react';
import { Purchases, CustomerInfo } from '@revenuecat/purchases-js';

interface UserData {
    id?: number;
    anonymous_id?: string;
    email?: string;
    display_name?: string;
    profile_picture_url?: string;
    is_subscriber?: boolean;
    subscription_status?: string;
    subscription_provider?: string;
}

interface AuthContextType {
    isLoggedIn: boolean;
    isSubscriber: boolean;
    sessionToken: string | null;
    userId: string | null;
    userEmail: string | null;
    isLoading: boolean;
    customerInfo: CustomerInfo | null;
    subscriptionProvider: string | null;
    login: (token: string, anonymousId?: string, email?: string) => void;
    logout: () => void;
    managedFetch: (url: string, options?: RequestInit) => Promise<Response>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const BROWSER_SIGNATURE_KEY = 'carpeek_browser_signature';

export const getBrowserSignature = () => {
    let sig = localStorage.getItem(BROWSER_SIGNATURE_KEY);
    if (!sig) {
        sig = window.crypto.randomUUID();
        localStorage.setItem(BROWSER_SIGNATURE_KEY, sig);
    }
    return sig;
};

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isSubscriber, setIsSubscriber] = useState(false);
    const [sessionToken, setSessionToken] = useState<string | null>(null);
    const [userId, setUserId] = useState<string | null>(null);
    const [userEmail, setUserEmail] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [customerInfo, setCustomerInfo] = useState<CustomerInfo | null>(null);
    const [subscriptionProvider, setSubscriptionProvider] = useState<string | null>(null);
    const [rcInitialized, setRcInitialized] = useState(false);

    const sessionTokenRef = useRef<string | null>(null);

    useEffect(() => {
        sessionTokenRef.current = sessionToken;
    }, [sessionToken]);

    const initRevenueCat = (appUserId: string) => {
        if (typeof window === 'undefined' || rcInitialized) return;

        const apiKey = process.env.NEXT_PUBLIC_REVENUECAT_API_KEY;
        if (!apiKey) {
            console.warn('RevenueCat API Key not found');
            return;
        }

        try {
            Purchases.configure(apiKey, appUserId);
            setRcInitialized(true);
            
            // Fetch initial customer info to double check entitlements
            Purchases.getSharedInstance().getCustomerInfo().then(info => {
                setCustomerInfo(info);
                const entitlementId = process.env.NEXT_PUBLIC_REVENUECAT_ENTITLEMENT || 'ad_free';
                if (entitlementId in info.entitlements.active) {
                    setIsSubscriber(true);
                }
            }).catch(e => console.error("Failed to get RC customer info", e));

        } catch (e) {
            console.error('Failed to initialize RevenueCat', e);
        }
    };

    const initSession = async () => {
        try {
            const signature = getBrowserSignature();
            setUserId(signature);
            initRevenueCat(signature); // Initialize RC as soon as we have a signature

            const email = localStorage.getItem('autocorrect_user_email');
            if (email) {
                setUserEmail(email);
                setIsLoggedIn(true);
            }

            const res = await fetch('/api/auth/session', {
                headers: { 'X-Browser-Signature': signature }
            });

            if (res.ok) {
                const data = await res.json();
                if (data.token) {
                    setSessionToken(data.token);
                    sessionTokenRef.current = data.token;
                }
                if (data.is_subscriber) {
                    setIsSubscriber(true);
                }
                if (data.subscription_provider) {
                    setSubscriptionProvider(data.subscription_provider);
                }
                if (data.is_logged_in) {
                    setIsLoggedIn(true);
                }
            }
        } catch (e) {
            console.error("Session init failed", e);
        } finally {
            setIsLoading(false);
        }
    };

    useEffect(() => {
        initSession();
    }, []);

    const login = (token: string, anonymousId?: string, email?: string) => {
        setSessionToken(token);
        sessionTokenRef.current = token;
        setIsLoggedIn(true);
        if (anonymousId) {
            setUserId(anonymousId);
            localStorage.setItem(BROWSER_SIGNATURE_KEY, anonymousId);
            initRevenueCat(anonymousId); // Re-init RC with potentially new ID (unlikely to change, but good practice)
        }
        if (email) {
            setUserEmail(email);
            localStorage.setItem('autocorrect_user_email', email);
        }
    };

    const logout = () => {
        setSessionToken(null);
        sessionTokenRef.current = null;
        setIsLoggedIn(false);
        setUserEmail(null);
        localStorage.removeItem('autocorrect_user_email');
        
        // Don't clear browser signature on logout to maintain anonymous data
        // But do log out of RevenueCat explicitly if needed using Purchases.getSharedInstance().logOut()
    };

    const getHeaders = () => {
        const headers: Record<string, string> = {
            'Content-Type': 'application/json',
            'X-Timezone': Intl.DateTimeFormat().resolvedOptions().timeZone,
        };
        if (sessionTokenRef.current) {
            headers['Authorization'] = `Bearer ${sessionTokenRef.current}`;
        }
        return headers;
    };

    const managedFetch = async (url: string, options: RequestInit = {}): Promise<Response> => {
        const currentHeaders = getHeaders();
        let response = await fetch(url, {
            ...options,
            headers: {
                ...currentHeaders,
                ...(options.headers || {}),
            },
        });

        if (response.status === 401) {
            await initSession(); // Re-init session gets new token
            if (sessionTokenRef.current) {
                response = await fetch(url, {
                    ...options,
                    headers: {
                        ...getHeaders(),
                        ...(options.headers || {}),
                    },
                });
            }
        }
        return response;
    };

    return (
        <AuthContext.Provider value={{
            isLoggedIn,
            isSubscriber,
            sessionToken,
            userId,
            userEmail,
            isLoading,
            customerInfo,
            subscriptionProvider,
            login,
            logout,
            managedFetch
        }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};

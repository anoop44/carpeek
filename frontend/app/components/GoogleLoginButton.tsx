'use client';

import { GoogleOAuthProvider, GoogleLogin, CredentialResponse } from '@react-oauth/google';
import { useAuth, getBrowserSignature } from './AuthProvider';
import { useState } from 'react';

export default function GoogleLoginButton() {
    const { login } = useAuth();
    const [error, setError] = useState<string | null>(null);

    const handleGoogleSuccess = async (credentialResponse: CredentialResponse) => {
        if (!credentialResponse.credential) return;

        try {
            const signature = getBrowserSignature();
            const response = await fetch('/api/auth/google', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    id_token: credentialResponse.credential,
                    anonymous_id: signature
                }),
            });

            if (response.ok) {
                const user = await response.json();
                login(user.auth_token, user.anonymous_id, user.email);
            } else {
                console.error("Login failed");
                setError('Login failed from server.');
            }
        } catch (error) {
            console.error("Login error", error);
            setError('Failed to sign in with Google');
        }
    };

    return (
        <GoogleOAuthProvider clientId={process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || ''}>
            <div className="flex flex-col items-center gap-2">
                <GoogleLogin
                    onSuccess={handleGoogleSuccess}
                    onError={() => {
                        console.error('Google Login Failed');
                        setError('Google Login window closed or failed');
                    }}
                    theme="filled_black"
                    shape="pill"
                    text="continue_with"
                    useOneTap={false}
                />
                {error && (
                    <p className="text-red-400 text-xs mt-1">{error}</p>
                )}
            </div>
        </GoogleOAuthProvider>
    );
}

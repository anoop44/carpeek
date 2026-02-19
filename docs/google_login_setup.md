# Google Login Setup Guide

This guide details the steps to configure Google Login for the AutoCorrect application.

## Prerequisites

- A Google Cloud Platform (GCP) Project
- Access to the AutoCorrect codebase

## Step 1: Create Google OAuth Credentials

1.  Go to the [Google Cloud Console](https://console.cloud.google.com/).
2.  Select your project or create a new one.
3.  Navigate to **APIs & Services** > **Credentials**.
4.  Click **Create Credentials** and select **OAuth client ID**.
5.  If prompted, configure the **OAuth consent screen** (External is usually fine for testing).
6.  Select **Web application** as the Application type.
7.  Set the name (e.g., "AutoCorrect Web").
8.  **Authorized JavaScript origins**:
    - Add `http://localhost:3000` (for local development).
    - Add your production domain (e.g., `https://your-app.com`).
9.  **Authorized redirect URIs**:
    - Add `http://localhost:3000` (and production URL).
10. Click **Create**.
11. Copy the **Client ID** (you will need this for the frontend).

## Step 2: Configure Frontend (Next.js)

1.  Open `frontend/.env.local` (create if it doesn't exist).
2.  Add the Client ID:
    ```env
    NEXT_PUBLIC_GOOGLE_CLIENT_ID=your-google-client-id-here
    ```

## Step 3: Configure Backend (Go)

The backend verifies the ID token sent by the frontend.

1.  Open `backend/.env` (or set environment variables in your deployment).
2.  Ensure you have established a database connection string.
3.  (Optional) If you implement server-side validation libraries that require the secret (not currently used in the basic flow, but good practice), add `GOOGLE_CLIENT_SECRET`.

## Step 4: Verify Implementation

1.  Start the backend (`go run cmd/server/main.go`).
2.  Start the frontend (`npm run dev`).
3.  Navigate to the Leaderboard page.
4.  Click the **"Sign in with Google"** button (or the banner after completing a challenge).
5.  Complete the Google sign-in flow.
6.  Verify that your user ID persists and your stats are linked.

## Troubleshooting

-   **"Origin mismatch" error**: Ensure `http://localhost:3000` is exactly listed in Authorized JavaScript origins.
-   **"Popup closed by user"**: The user closed the window before finishing.
-   **Token verification failed**: Check backend logs. Ensure the token sent matches the one expected by Google's tokeninfo endpoint.

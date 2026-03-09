# Subscription Feature: Manual Tasks & Local Testing Guide

The RevenueCat integration and overarching subscription feature have been fully implemented across the frontend and backend. 
To finalize the setup and test it locally, there are manual configurations required in the RevenueCat dashboard and your local development environment.

## Phase 1: RevenueCat Dashboard Tasks (Manual)

1. **Project & App Creation**
   - Go to your [RevenueCat Dashboard](https://app.revenuecat.com/).
   - Create a new project for "AutoCorrect".
   - Under Project Settings, add a new **Web App** (and mobile apps if necessary). Note down the **Web API Key**.

2. **Define Entitlements and Products**
   - **Entitlements**: Create an entitlement named `ad_free` (this is the value configured in code. If you use a different name, update `NEXT_PUBLIC_REVENUECAT_ENTITLEMENT` in the `.env`).
   - **Products**: Connect your payment gateway (e.g., Stripe for web). Create products in Stripe and import them into RevenueCat.
   - **Offerings**: Create a new offering (e.g., `default`). Attach your products to this offering. The frontend uses `Purchases.getSharedInstance().getOfferings()` to fetch these options on the Subscribe page.

3. **Configure Webhook**
   - Once your server is publicly accessible (or via Ngrok for testing), go to **Project Settings > Webhooks** in RevenueCat.
   - Add your webhook URL. Production example: `https://api.yourdomain.com/api/v1/subscription/webhook`.
   - Add an Authorization header: `Bearer <YOUR_SECRET_TOKEN>`.

---

## Phase 2: Environment Variables

Ensure your environment files are populated with the correct keys.

### Frontend (`frontend/.env.local` or `.env`)
```
NEXT_PUBLIC_REVENUECAT_API_KEY=rc_your_web_api_key_here
NEXT_PUBLIC_REVENUECAT_ENTITLEMENT=ad_free
```

### Backend (`backend/.env`)
```
# The token you enter in the RevenueCat dashboard exactly as-is under Authorization header (e.g. Bearer my_secret_key)
REVENUECAT_WEBHOOK_SECRET=my_secret_key
```

---

## Phase 3: How to Test Subscriptions Locally

Testing the end-to-end flow requires your local backend to receive HTTP POST requests from RevenueCat's servers when a subscription event occurs.

### Step 1: Tunnel Localhost with Ngrok
Since your local backend isn't accessible from the internet, use `ngrok` to expose it:
```bash
ngrok http 8080
```
*Note the HTTPS forwarding URL given by ngrok (e.g., `https://random-123.ngrok-free.app`).*

### Step 2: Point RevenueCat Intermittently
Navigate to your RevenueCat Webhook settings and temporarily set the URL to your ngrok tunnel:
`https://random-123.ngrok-free.app/api/v1/subscription/webhook`
*(Make sure to change this back to your production server URL when done testing!)*

### Step 3: Run the Local Apps
- Backend: `make run` or `go run ./cmd/server`
- Frontend: `npm run dev`

### Step 4: Perform a Test Purchase
1. **Google Login:** Open the locally running frontend app, go to the Leaderboard, and sign in with Google. Subscriptions require Google login as per your requirements.
2. **Access Subscription Flow:** Go to `Settings` > `Subscribe Now` or click "Remove Ads" on a Banner Ad.
3. **Execute Purchase:** On the test build, select a RevenueCat offering. Since you're using a Stripe connection in Test Mode, complete the purchase using a [Stripe test card](https://stripe.com/docs/testing) (e.g., `4242 4242 4242 4242`).
4. **Validation:**
   - **Frontend:** You should be redirected back to the `/settings` page, which should now reflect **"Ad-Free Active"**.
   - **Backend:** Check your backend console logs. You should see a `REVENUECAT` log detailing that a `INITIAL_PURCHASE` webhook was received from RevenueCat, confirming the database synced and updated the `is_subscriber` flag.

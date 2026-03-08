import type { Metadata } from 'next'
import { Space_Grotesk } from 'next/font/google'
import './globals.css'

const spaceGrotesk = Space_Grotesk({
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-space-grotesk',
})

export const metadata: Metadata = {
  title: 'AutoCorrect | Daily Car Challenge',
  description: 'Identify the vehicle by its lighting signature. New challenge every day.',
}

const adsenseClientId = process.env.NEXT_PUBLIC_ADSENSE_CLIENT_ID;

// Server-side logging for container verification
if (adsenseClientId) {
  console.log(`[Configuration] AdSense Client ID detected: ${adsenseClientId}`);
} else {
  console.log('[Configuration] AdSense Client ID not found. Ads will be disabled.');
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className={`${spaceGrotesk.variable} dark`}>
      <head>
        <link
          href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:wght,FILL@100..700,0..1&display=swap"
          rel="stylesheet"
        />
        {/* Google AdSense — only injected when client ID is configured */}
        {adsenseClientId && (
          <script
            async
            src={`https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=${adsenseClientId}`}
            crossOrigin="anonymous"
          />
        )}
      </head>
      <body className="bg-background-light dark:bg-background-dark text-slate-900 dark:text-white min-h-screen">
        {children}
      </body>
    </html>
  )
}

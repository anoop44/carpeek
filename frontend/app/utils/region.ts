/**
 * Detects if the user is likely in India based on timezone or locale.
 * This is a client-side heuristic. For stricter checks, a GeoIP API would be needed.
 */
export const isInIndia = (): boolean => {
  try {
    // 1. Check Timezone (Asia/Kolkata)
    const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    if (timezone === 'Asia/Kolkata') return true;

    // 2. Check Locale (en-IN, hi-IN, etc.)
    const locale = navigator.language || (navigator as any).userLanguage;
    if (locale && locale.endsWith('-IN')) return true;

    // 3. Fallback to storage if user previously selected India (optional)
    if (typeof window !== 'undefined') {
      const stored = localStorage.getItem('user_region');
      if (stored === 'IN') return true;
    }
  } catch (e) {
    console.warn('Region detection failed:', e);
  }

  return false;
};

/**
 * Returns the detected country code
 */
export const getDetectedCountry = (): string => {
  return isInIndia() ? 'IN' : 'GLOBAL';
};

import { NextResponse } from 'next/server';

export async function proxyToBackend(
    request: Request,
    path: string,
    options: {
        method?: string;
        body?: unknown;
        headers?: Record<string, string>;
    } = {}
) {
    const BACKEND_URL = process.env.NEXT_PUBLIC_API_URL;
    const targetUrl = `${BACKEND_URL}${path}`;

    console.log(`Proxying request to backend: ${options.method || 'GET'} ${targetUrl}`);

    const headers = new Headers();
    headers.set('Content-Type', 'application/json');

    // Forward relevant headers from the client
    const forwardedHeaders = [
        'x-user-id',
        'x-timezone',
        'authorization',
        'x-forwarded-for',
        'x-browser-signature'
    ];

    forwardedHeaders.forEach(h => {
        const val = request.headers.get(h);
        if (val) headers.set(h, val);
    });

    // Add custom headers if provided
    if (options.headers) {
        Object.entries(options.headers).forEach(([k, v]) => headers.set(k, v));
    }

    try {
        const response = await fetch(targetUrl, {
            method: options.method || 'GET',
            headers: headers,
            body: options.body ? JSON.stringify(options.body) : undefined,
        });

        const contentType = response.headers.get('content-type') || '';
        let data;

        if (contentType.includes('application/json')) {
            data = await response.json();
        } else {
            // This should rarely happen now that backend is fixed, but good for safety
            const text = await response.text();
            data = { error: text.trim() || 'Unknown error from backend', status: response.status };
        }

        const res = NextResponse.json(data, { status: response.status });

        // Forward User Identity if updated (X-User-ID header)
        const updatedUserId = response.headers.get('X-User-ID');
        if (updatedUserId) {
            res.headers.set('X-User-ID', updatedUserId);
        }

        return res;
    } catch (error) {
        console.error(`BFF Proxy Error (${path}):`, error);
        return NextResponse.json(
            { error: 'Backend connection failed' },
            { status: 503 }
        );
    }
}
export async function proxyBinaryToBackend(
    request: Request,
    path: string
) {
    const BACKEND_URL = process.env.NEXT_PUBLIC_API_URL;
    const targetUrl = `${BACKEND_URL}${path}`;

    const headers = new Headers();
    // Forward relevant headers from the client
    const forwardedHeaders = [
        'x-user-id',
        'x-timezone',
        'authorization',
        'x-forwarded-for',
        'x-browser-signature'
    ];

    forwardedHeaders.forEach(h => {
        const val = request.headers.get(h);
        if (val) headers.set(h, val);
    });

    try {
        const response = await fetch(targetUrl, {
            method: 'GET',
            headers: headers,
        });

        if (!response.ok) {
            return new NextResponse(null, { status: response.status });
        }

        const blob = await response.blob();
        const res = new NextResponse(blob);

        // Forward content-type
        const contentType = response.headers.get('content-type');
        if (contentType) {
            res.headers.set('content-type', contentType);
        }

        return res;
    } catch (error) {
        console.error(`BFF Binary Proxy Error (${path}):`, error);
        return new NextResponse(null, { status: 503 });
    }
}

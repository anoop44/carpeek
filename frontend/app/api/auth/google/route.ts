import { proxyToBackend } from '../../utils';

export async function POST(request: Request) {
    const body = await request.json();
    return proxyToBackend(request, '/api/v1/auth/google', {
        method: 'POST',
        body
    });
}

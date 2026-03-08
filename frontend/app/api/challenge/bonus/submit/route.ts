import { proxyToBackend } from '../../../utils';

export async function POST(request: Request) {
    const body = await request.json();
    return proxyToBackend(request, '/api/v1/challenge/bonus/submit', {
        method: 'POST',
        body
    });
}

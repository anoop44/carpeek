import { proxyToBackend } from '../utils';

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url);
    const type = searchParams.get('type') || 'daily';
    return proxyToBackend(request, `/api/v1/leaderboard?type=${type}`);
}

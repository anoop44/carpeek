import { proxyToBackend } from '../../utils';

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url);
    const challenge_id = searchParams.get('challenge_id');
    return proxyToBackend(request, `/api/v1/challenge/stats?challenge_id=${challenge_id}`);
}

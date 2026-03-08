import { proxyToBackend } from '../../../utils';

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url);
    const challengeId = searchParams.get('challenge_id');
    return proxyToBackend(request, `/api/v1/challenge/bonus/status?challenge_id=${challengeId}`);
}

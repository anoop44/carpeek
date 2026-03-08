import { proxyToBackend } from '../../utils';

export async function GET(request: Request) {
    return proxyToBackend(request, '/api/v1/challenge/today');
}

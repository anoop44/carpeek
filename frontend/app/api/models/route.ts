import { proxyToBackend } from '../utils';

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url);
    const make_id = searchParams.get('make_id');
    return proxyToBackend(request, `/api/v1/models?make_id=${make_id}`);
}

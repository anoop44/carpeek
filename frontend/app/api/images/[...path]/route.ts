import { proxyBinaryToBackend } from '../../utils';

export async function GET(
    request: Request,
    { params }: { params: Promise<{ path: string[] }> }
) {
    const { path } = await params;
    const fullPath = path.join('/');
    return proxyBinaryToBackend(request, `/images/${fullPath}`);
}

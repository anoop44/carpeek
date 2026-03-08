'use client';

import { useEffect } from 'react';

interface ImageModalProps {
    isOpen: boolean;
    onClose: () => void;
    imageUrl: string;
}

export default function ImageModal({ isOpen, onClose, imageUrl }: ImageModalProps) {
    useEffect(() => {
        if (!isOpen) return;
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === 'Escape') onClose();
        };
        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [isOpen, onClose]);

    if (!isOpen) return null;

    return (
        <div
            className="fixed inset-0 z-[110] flex items-center justify-center p-4 cursor-zoom-out"
            onClick={onClose}
        >
            {/* Backdrop */}
            <div className="absolute inset-0 bg-black/95 backdrop-blur-md animate-fade-in"></div>

            {/* Close Button */}
            <button
                onClick={onClose}
                className="absolute top-6 right-6 z-[120] w-12 h-12 flex items-center justify-center rounded-full bg-white/10 hover:bg-white/20 text-white transition-all shadow-2xl border border-white/20"
                title="Close"
            >
                <span className="material-symbols-outlined text-2xl">close</span>
            </button>

            {/* Image Container */}
            <div
                className="relative max-w-full max-h-full flex items-center justify-center animate-scale-in"
                onClick={(e) => e.stopPropagation()}
            >
                <img
                    src={imageUrl}
                    alt="Enlarged Car Detail"
                    className="max-w-full max-h-[90vh] md:max-h-[95vh] object-contain rounded-lg shadow-2xl border border-white/10 pointer-events-none"
                />
            </div>
        </div>
    );
}

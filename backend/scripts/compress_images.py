#!/usr/bin/env python3
"""
Script to compress images in a directory structure.

This script:
- Takes a source directory and output directory as arguments
- Recreates the source directory structure in the output directory
- Compresses JPG, JPEG, and PNG images to reduce file size
- Preserves original filenames and folder structure
- Uses Pillow (PIL) for image processing

Usage:
  ./compress_images.py <source_dir> <output_dir> [--quality QUALITY] [--max-size MAX_SIZE] [--dry-run]

Arguments:
  source_dir    Source directory containing images
  output_dir    Output directory for compressed images
  --quality     JPEG quality (1-100, default: 85)
  --max-size    Maximum dimension for resizing (default: None, no resizing)
  --dry-run1    Show what would be done without actually compressing
"""

import os
import sys
import argparse
import logging
from pathlib import Path
try:
    from PIL import Image
    PIL_AVAILABLE = True
except ImportError:
    PIL_AVAILABLE = False
    print("Warning: Pillow (PIL) library not found. Please install with: pip install Pillow")

def compress_image(input_path, output_path, quality=85, max_size=None):
    """
    Compress a single image file.
    
    Args:
        input_path: Path to input image
        output_path: Path to output image
        quality: JPEG quality (1-100)
        max_size: Maximum dimension for resizing (width or height)
    
    Returns:
        tuple: (success, message)
    """
    try:
        # Get original file size
        original_size = os.path.getsize(input_path)
        
        # Open the image
        with Image.open(input_path) as img:
            # Convert to RGB if necessary (for JPEG saving)
            if img.mode in ("RGBA", "P"):
                img = img.convert("RGB")
            
            # Resize if max_size is specified
            if max_size:
                # Get current dimensions
                width, height = img.size
                
                # Calculate new dimensions while maintaining aspect ratio
                if width > height:
                    if width > max_size:
                        new_width = max_size
                        new_height = int(height * (max_size / width))
                    else:
                        new_width, new_height = width, height
                else:
                    if height > max_size:
                        new_height = max_size
                        new_width = int(width * (max_size / height))
                    else:
                        new_width, new_height = width, height
                
                # Resize if needed
                if (new_width, new_height) != (width, height):
                    img = img.resize((new_width, new_height), Image.Resampling.LANCZOS)
            
            # Save the image with compression
            if input_path.suffix.lower() in ['.jpg', '.jpeg']:
                img.save(output_path, 'JPEG', quality=quality, optimize=True, progressive=True)
            elif input_path.suffix.lower() == '.png':
                # For PNG, use optimize and reduce colors if possible
                img.save(output_path, 'PNG', optimize=True, compress_level=6)
            else:
                # Should not happen due to filtering, but just in case
                img.save(output_path)
        
        # Get compressed file size
        compressed_size = os.path.getsize(output_path)
        size_reduction = original_size - compressed_size
        reduction_percent = (size_reduction / original_size) * 100 if original_size > 0 else 0
        
        return True, f"Compressed: {input_path} -> {output_path} ({original_size:,}B → {compressed_size:,}B, {reduction_percent:.1f}% reduction)"
    
    except Exception as e:
        return False, f"Error compressing {input_path}: {str(e)}"

def process_directory(source_dir, output_dir, quality=85, max_size=None, dry_run=False):
    """
    Process all images in a directory recursively.
    
    Args:
        source_dir: Source directory path
        output_dir: Output directory path
        quality: JPEG quality
        max_size: Maximum dimension for resizing
        dry_run: If True, only show what would be done without actual compression
    
    Returns:
        tuple: (success_count, error_count, total_count)
    """
    success_count = 0
    error_count = 0
    total_count = 0
    
    # Walk through the source directory
    for root, dirs, files in os.walk(source_dir):
        # Create corresponding directory in output
        relative_path = os.path.relpath(root, source_dir)
        output_root = os.path.join(output_dir, relative_path)
        
        # Create directory if it doesn't exist
        if not dry_run:
            os.makedirs(output_root, exist_ok=True)
        
        # Process each file
        for filename in files:
            file_path = os.path.join(root, filename)
            output_path = os.path.join(output_root, filename)
            
            # Check if file is an image we want to compress
            if filename.lower().endswith(('.jpg', '.jpeg', '.png')):
                total_count += 1
                
                if dry_run:
                    print(f"Would compress: {file_path} -> {output_path}")
                    success_count += 1
                else:
                    success, message = compress_image(Path(file_path), Path(output_path), quality, max_size)
                    if success:
                        success_count += 1
                        print(f"✓ {message}")
                    else:
                        error_count += 1
                        print(f"✗ {message}")
    
    return success_count, error_count, total_count

def main():
    parser = argparse.ArgumentParser(description='Compress images in a directory structure')
    parser.add_argument('source_dir', help='Source directory containing images')
    parser.add_argument('output_dir', help='Output directory for compressed images')
    parser.add_argument('--quality', type=int, default=85, 
                        help='JPEG quality (1-100, default: 85)')
    parser.add_argument('--max-size', type=int, 
                        help='Maximum dimension for resizing (default: None, no resizing)')
    parser.add_argument('--dry-run', action='store_true',
                        help='Show what would be done without actually compressing')
    
    args = parser.parse_args()
    
    # Validate input directories
    source_path = Path(args.source_dir)
    output_path = Path(args.output_dir)
    
    if not source_path.exists():
        print(f"Error: Source directory '{args.source_dir}' does not exist.")
        sys.exit(1)
    
    if not source_path.is_dir():
        print(f"Error: '{args.source_dir}' is not a directory.")
        sys.exit(1)
    
    # Create output directory if it doesn't exist
    output_path.mkdir(parents=True, exist_ok=True)
    
    print(f"Starting image compression...")
    print(f"Source: {source_path}")
    print(f"Output: {output_path}")
    print(f"JPEG Quality: {args.quality}")
    if args.max_size:
        print(f"Max Size: {args.max_size}px")
    print("-" * 50)
    
    # Process the directory
    success_count, error_count, total_count = process_directory(
        source_path, output_path, args.quality, args.max_size, args.dry_run
    )
    
    # Print summary
    print("\n" + "="*50)
    print("COMPRESSION SUMMARY")
    print("="*50)
    print(f"Total images processed: {total_count}")
    print(f"Successfully compressed: {success_count}")
    print(f"Errors: {error_count}")
    
    if error_count == 0:
        print("✓ All images processed successfully!")
        sys.exit(0)
    else:
        print(f"⚠ {error_count} errors occurred.")
        sys.exit(1)

if __name__ == "__main__":
    main()
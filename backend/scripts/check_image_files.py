#!/usr/bin/env python3
"""
Script to verify image files exist in the images/actual directory.

This script reads image URLs from backend/migrations/002_seed_makes_models_up.sql
and checks whether the corresponding files exist in images/actual.
It creates two lists:
1. Files that are missing from the images/actual directory
2. Image URLs that contain dots (.) in the filename part (before the extension)
"""

import os
import re
import sys

def main():
    # Define paths
    sql_file_path = os.path.join(os.path.dirname(__file__), '..', 'migrations', '002_seed_makes_models_up.sql')
    images_dir = os.path.join(os.path.dirname(__file__), '..', '..', 'images', 'actual')
    
    # Check if SQL file exists
    if not os.path.exists(sql_file_path):
        print(f"Error: SQL file not found at {sql_file_path}")
        sys.exit(1)
    
    # Check if images directory exists
    if not os.path.exists(images_dir):
        print(f"Error: Images directory not found at {images_dir}")
        sys.exit(1)
    
    # Read SQL file
    with open(sql_file_path, 'r') as f:
        sql_content = f.read()
    
    # Extract image URLs using regex
    # Pattern matches: 'image_url', 'path/to/image.jpg'
    # We're looking for values in the INSERT INTO models statement
    image_urls = []
    
    # First, try to find the specific pattern used in the SQL file
    # Based on the sample, the format is: 'image_url', 'make/filename.ext'
    pattern = r"image_url,\s*'([^']*)'"
    matches = re.findall(pattern, sql_content)
    
    # If no matches, try a more flexible approach
    if not matches:
        # Look for all values in the VALUES clause that are likely image URLs
        # Image URLs typically contain file extensions and are quoted
        pattern2 = r"'([^']+?\.(jpg|jpeg|png|gif|webp))'"
        all_matches = re.findall(pattern2, sql_content)
        matches = [match[0] for match in all_matches]
    
    # Filter out empty strings, duplicates, and non-image values
    image_urls = []
    for url in matches:
        url = url.strip()
        if url and url.endswith(('.jpg', '.jpeg', '.png', '.gif', '.webp')):
            image_urls.append(url)
    
    # Remove duplicates while preserving order
    seen = set()
    unique_urls = []
    for url in image_urls:
        if url not in seen:
            seen.add(url)
            unique_urls.append(url)
    
    image_urls = unique_urls
    
    print(f"Found {len(image_urls)} unique image URLs in the SQL file.")
    
    # Lists to store results
    missing_files = []
    urls_with_dots = []
    
    # Check each image URL
    for url in image_urls:
        # Skip empty URLs
        if not url:
            continue
            
        # Check if URL contains dots (.) before the extension
        # Split on last dot to separate filename from extension
        if '.' in url:
            # Get the part before the extension (filename without extension)
            if url.count('.') > 1:
                # More than one dot means there's at least one dot before the extension
                urls_with_dots.append(url)
            elif url.count('.') == 1:
                # Exactly one dot - check if it's before the extension
                # If the dot is at the beginning or very early, it might be part of the path
                # But for our purpose, we'll consider any single dot as potentially problematic
                # unless it's clearly part of the extension pattern
                if not url.endswith('.jpg') and not url.endswith('.jpeg') and not url.endswith('.png') and not url.endswith('.gif') and not url.endswith('.webp'):
                    urls_with_dots.append(url)
        
        # Construct full path to check if file exists
        # URL format is like: opel/Opel_Astra_1.2_Turbo_Ultimate_(L)_–_f_13122024.jpg
        # So we need to join images_dir with the URL
        full_path = os.path.join(images_dir, url)
        
        # Check if file exists
        if not os.path.exists(full_path):
            missing_files.append(url)
    
    # Print results
    print(f"\nMissing files ({len(missing_files)}):")
    for url in missing_files:
        print(f"  - {url}")
    
    print(f"\nURLs with dots in filename (before extension) ({len(urls_with_dots)}):")
    for url in urls_with_dots:
        print(f"  - {url}")
    
    # Summary
    print(f"\nSummary:")
    print(f"  Total image URLs: {len(image_urls)}")
    print(f"  Missing files: {len(missing_files)}")
    print(f"  URLs with dots in filename: {len(urls_with_dots)}")
    
    # Exit with error code if there are missing files
    if missing_files:
        sys.exit(1)
    else:
        print("All image files found.")

if __name__ == "__main__":
    main()
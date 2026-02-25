#!/usr/bin/env python3

import os
import requests
import urllib.parse
import sys
import time
from pathlib import Path

def main():
    # Read the car_images.txt file
    with open('/home/darsana/Projects/apps/CarPeek/backend/scripts/car_images.txt', 'r') as f:
        lines = f.readlines()
    
    # Create directory for images
    image_dir = Path('/home/darsana/Projects/apps/CarPeek/backend/scripts/car_images')
    image_dir.mkdir(exist_ok=True)
    
    # Track used filenames, duplicate URLs, and already downloaded files
    filename_counter = {}
    duplicate_urls = []
    processed_urls = set()
    
    # Set headers to mimic a real browser (to avoid 403 errors from Wikipedia)
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36',
        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8',
        'Accept-Language': 'en-US,en;q=0.5',
        'Accept-Encoding': 'gzip, deflate',
        'Connection': 'keep-alive',
        'Upgrade-Insecure-Requests': '1',
        'Sec-Fetch-Dest': 'document',
        'Sec-Fetch-Mode': 'navigate',
        'Sec-Fetch-Site': 'none',
        'Sec-Fetch-User': '?1',
    }
    
    # Get list of already downloaded files
    existing_files = set(f.stem for f in image_dir.iterdir() if f.is_file() and f.name != 'duplicate_urls.txt')
    
    # Process each line
    downloaded_count = 0
    total_lines = len(lines)
    
    for i, line in enumerate(lines):
        line = line.strip()
        if not line:
            continue
            
        try:
            # Split by " - " to get parts
            parts = line.split(' - ')
            if len(parts) < 4:
                print(f"Warning: Line {i+1} doesn't have enough parts: {line}")
                continue
                
            manufacturer = parts[0].strip()
            model = parts[1].strip()
            generation = parts[2].strip()
            url = parts[3].strip()
            
            # Skip if URL is "NO IMAGE FOUND"
            if url == "NO IMAGE FOUND":
                print(f"Skipping line {i+1}: No image found")
                continue
                
            # Create base filename
            base_filename = f"{manufacturer}_{model}_{generation}"
            # Clean up the filename
            base_filename = ''.join(c if c.isalnum() or c in '_-' else '_' for c in base_filename)
            
            # Check if this file is already downloaded (with or without offset)
            filename_to_check = base_filename
            if filename_to_check in existing_files:
                # File already exists, check if we need offset
                counter = 1
                while f"{filename_to_check}_{counter}" in existing_files:
                    counter += 1
                filename_to_check = f"{filename_to_check}_{counter}"
            elif any(f.startswith(filename_to_check + '_') and f.split('_')[-1].isdigit() for f in existing_files):
                # File with offset exists, find next available
                counter = 1
                while f"{filename_to_check}_{counter}" in existing_files:
                    counter += 1
                filename_to_check = f"{filename_to_check}_{counter}"
            
            # Check if URL is duplicate
            if url in processed_urls:
                duplicate_urls.append(url)
            else:
                processed_urls.add(url)
            
            # Skip if file already exists
            if filename_to_check in existing_files:
                print(f"Skipping {i+1}/{total_lines}: {manufacturer} {model} {generation} (already downloaded)")
                continue
            
            # Download the image with proper headers and rate limiting
            print(f"Downloading {i+1}/{total_lines}: {manufacturer} {model} {generation}")
            try:
                # Add delay to avoid rate limiting (1 second between requests)
                time.sleep(1)
                
                response = requests.get(url, headers=headers, timeout=10)
                response.raise_for_status()
                
                # Get the file extension from the URL
                url_path = urllib.parse.urlparse(url).path
                if url_path.endswith('.jpg') or url_path.endswith('.jpeg'):
                    ext = '.jpg'
                elif url_path.endswith('.png'):
                    ext = '.png'
                elif url_path.endswith('.gif'):
                    ext = '.gif'
                else:
                    ext = '.jpg'  # default to jpg
                
                # Save the image
                with open(image_dir / f"{filename_to_check}{ext}", 'wb') as f:
                    f.write(response.content)
                    
                downloaded_count += 1
                print(f"✓ Downloaded: {filename_to_check}{ext}")
                
            except requests.exceptions.RequestException as e:
                print(f"✗ Error downloading {url}: {e}")
                # Try with thumbnail URL as fallback
                try:
                    # Extract filename from URL path
                    url_parts = urllib.parse.urlparse(url)
                    filename = os.path.basename(url_parts.path)
                    # Create thumbnail URL (Wikipedia thumbnail format)
                    thumb_url = f"https://upload.wikimedia.org/wikipedia/commons/thumb/{url_parts.path[1:]}/800px{url_parts.path}"
                    print(f"Trying thumbnail URL: {thumb_url}")
                    response = requests.get(thumb_url, headers=headers, timeout=10)
                    response.raise_for_status()
                    
                    # Save with original filename
                    with open(image_dir / f"{filename_to_check}{ext}", 'wb') as f:
                        f.write(response.content)
                        
                    downloaded_count += 1
                    print(f"✓ Downloaded (thumbnail): {filename_to_check}{ext}")
                    
                except Exception as e2:
                    print(f"✗ Thumbnail also failed: {e2}")
                    continue
                
        except Exception as e:
            print(f"Error processing line {i+1}: {e}")
            continue
    
    # Write duplicate URLs
    with open(image_dir / 'duplicate_urls.txt', 'w') as f:
        for url in duplicate_urls:
            f.write(url + '\n')
    
    print(f"\nDownload complete!")
    print(f"Total images downloaded: {downloaded_count}")
    print(f"Duplicate URLs found: {len(duplicate_urls)}")
    print(f"Total files in directory: {len([f for f in image_dir.iterdir() if f.is_file()])}")

if __name__ == "__main__":
    main()
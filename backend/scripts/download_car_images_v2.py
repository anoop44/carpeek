#!/usr/bin/env python3

import os
import requests
import urllib.parse
import sys
from pathlib import Path

def main():
    # Read the car_images.txt file
    with open('/home/darsana/Projects/apps/CarPeek/backend/scripts/car_images.txt', 'r') as f:
        lines = f.readlines()
    
    # Create directory for images
    image_dir = Path('/home/darsana/Projects/apps/CarPeek/backend/scripts/car_images')
    image_dir.mkdir(exist_ok=True)
    
    # Track used filenames and duplicate URLs
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
    
    # Process each line
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
                print(f"Skipping line {i+1}: No image found for {manufacturer} {model} {generation}")
                continue
                
            # Create base filename
            # Replace spaces with underscores and remove special characters
            base_filename = f"{manufacturer}_{model}_{generation}"
            # Clean up the filename (remove special characters that might cause issues)
            base_filename = ''.join(c if c.isalnum() or c in '_-' else '_' for c in base_filename)
            
            # Check if URL is duplicate
            if url in processed_urls:
                duplicate_urls.append(url)
                # Add offset number to filename
                if base_filename in filename_counter:
                    filename_counter[base_filename] += 1
                    filename = f"{base_filename}_{filename_counter[base_filename]}"
                else:
                    filename_counter[base_filename] = 1
                    filename = f"{base_filename}_1"
            else:
                processed_urls.add(url)
                filename = base_filename
            
            # Ensure filename is unique by checking if file already exists
            final_filename = filename
            counter = 1
            while (image_dir / f"{final_filename}.jpg").exists():
                final_filename = f"{filename}_{counter}"
                counter += 1
            
            # Download the image with proper headers
            print(f"Downloading {i+1}/{len(lines)}: {manufacturer} {model} {generation}")
            try:
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
                with open(image_dir / f"{final_filename}{ext}", 'wb') as f:
                    f.write(response.content)
                    
            except requests.exceptions.RequestException as e:
                print(f"Error downloading {url}: {e}")
                continue
                
        except Exception as e:
            print(f"Error processing line {i+1}: {e}")
            continue
    
    # Write duplicate URLs to a file
    with open(image_dir / 'duplicate_urls.txt', 'w') as f:
        for url in duplicate_urls:
            f.write(url + '\n')
    
    print(f"\nDownload complete!")
    print(f"Total images downloaded: {len([f for f in image_dir.iterdir() if f.is_file() and f.suffix in ['.jpg', '.jpeg', '.png', '.gif']])}")
    print(f"Duplicate URLs found: {len(duplicate_urls)}")
    print(f"Duplicate URLs saved to: {image_dir}/duplicate_urls.txt")

if __name__ == "__main__":
    main()
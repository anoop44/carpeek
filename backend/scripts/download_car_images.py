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
            
            # Download the image
            print(f"Downloading {i+1}/{len(lines)}: {manufacturer} {model} {generation}")
            try:
                response = requests.get(url, timeout=10)
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
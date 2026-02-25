#!/usr/bin/env python3

import os
from pathlib import Path

def main():
    # Read the car_images.txt file
    with open('/home/darsana/Projects/apps/CarPeek/backend/scripts/car_images.txt', 'r') as f:
        lines = f.readlines()
    
    # Get list of downloaded files (excluding duplicate_urls.txt)
    image_dir = Path('/home/darsana/Projects/apps/CarPeek/backend/scripts/car_images')
    downloaded_files = [f for f in image_dir.iterdir() if f.is_file() and f.name != 'duplicate_urls.txt']
    
    # Extract filenames without extension for comparison
    downloaded_basenames = set()
    for f in downloaded_files:
        # Remove extension and get base name
        basename = f.stem
        downloaded_basenames.add(basename)
    
    # Parse car_images.txt to get expected filenames
    expected_filenames = []
    missing_filenames = []
    
    for i, line in enumerate(lines):
        line = line.strip()
        if not line:
            continue
            
        try:
            # Split by " - " to get parts
            parts = line.split(' - ')
            if len(parts) < 4:
                continue
                
            manufacturer = parts[0].strip()
            model = parts[1].strip()
            generation = parts[2].strip()
            
            # Create base filename (same logic as download script)
            base_filename = f"{manufacturer}_{model}_{generation}"
            # Clean up the filename (remove special characters that might cause issues)
            base_filename = ''.join(c if c.isalnum() or c in '_-' else '_' for c in base_filename)
            
            expected_filenames.append(base_filename)
            
            # Check if this filename is in downloaded files
            if base_filename not in downloaded_basenames:
                # Also check if there are offset versions (e.g., base_filename_1, base_filename_2)
                found = False
                for j in range(1, 10):  # Check up to offset 9
                    if f"{base_filename}_{j}" in downloaded_basenames:
                        found = True
                        break
                if not found:
                    missing_filenames.append(f"{i+1}: {manufacturer} - {model} - {generation}")
                    
        except Exception as e:
            print(f"Error processing line {i+1}: {e}")
            continue
    
    # Print results
    print(f"Total cars in car_images.txt: {len(expected_filenames)}")
    print(f"Total downloaded images: {len(downloaded_basenames)}")
    print(f"Missing images: {len(missing_filenames)}")
    
    if missing_filenames:
        print("\nMissing images:")
        for missing in missing_filenames:
            print(f"  {missing}")
        
        # Write missing files to a file
        with open(image_dir / 'missing_images.txt', 'w') as f:
            f.write(f"Total missing: {len(missing_filenames)}\n\n")
            for missing in missing_filenames:
                f.write(f"{missing}\n")
    
    print(f"\nMissing images saved to: {image_dir}/missing_images.txt")

if __name__ == "__main__":
    main()
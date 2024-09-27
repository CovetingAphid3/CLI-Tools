#!/usr/bin/env python3

import os
import argparse
from tabulate import tabulate

def count_lines_in_file(file_path):
    """Count lines in a single file."""
    with open(file_path, 'r', encoding='utf-8') as file:
        return sum(1 for line in file if line.strip())  # Count non-empty lines

def count_lines_in_directory(directory):
    """Count lines in all files in a directory."""
    total_lines = 0
    file_counts = []  # List to hold tuples of (file_path, line_count)
    
    for root, _, files in os.walk(directory):
        for file in files:
            if file.endswith(('.py', '.js', '.html', '.css', '.ts', '.d.ts')):  # Include .ts files
                file_path = os.path.join(root, file)
                lines = count_lines_in_file(file_path)
                file_counts.append((file_path, lines))
                total_lines += lines
    
    return total_lines, file_counts

def main():
    parser = argparse.ArgumentParser(description='Count lines of code in a file or directory.')
    parser.add_argument('-f', '--file', type=str, help='Path to the file to count lines')
    parser.add_argument('-d', '--directory', type=str, help='Path to the directory to count lines')

    args = parser.parse_args()

    if args.file and args.directory:
        print("Please specify either a file or a directory, not both.")
        return

    if args.file:
        if os.path.isfile(args.file):
            lines = count_lines_in_file(args.file)
            print(f"Lines in file '{args.file}': {lines}")
        else:
            print("Invalid file path. Please enter a valid file.")

    elif args.directory:
        if os.path.isdir(args.directory):
            total_lines, file_counts = count_lines_in_directory(args.directory)
            print(tabulate(file_counts, headers=["File Path", "Lines Counted"], tablefmt="fancy_grid"))
            print(f"\nTotal lines in directory '{args.directory}': {total_lines}")
        else:
            print("Invalid directory path. Please enter a valid directory.")

    else:
        print("Please specify either a file with -f or a directory with -d.")

if __name__ == "__main__":
    main()


#!/usr/bin/env python3

import os
import requests
import argparse
from dotenv import load_dotenv
from prettytable import PrettyTable

load_dotenv()

GITHUB_TOKEN = os.getenv("GITHUB_TOKEN")
BASE_URL = "https://api.github.com"

def get_headers():
    return {
        "Authorization": f"token {GITHUB_TOKEN}",
        "Accept": "application/vnd.github.v3+json"
    }

def get_repo_info(owner, repo):
    response = requests.get(f"{BASE_URL}/repos/{owner}/{repo}", headers=get_headers())
    return response.json()

def list_repositories(owner):
    """List repositories for a specific user."""
    response = requests.get(f"{BASE_URL}/users/{owner}/repos", headers=get_headers())
    return response.json()



def display_repo_info(repo_info):
    """Display repository information in a table format."""
    table = PrettyTable()
    table.field_names = ["Key", "Value"]

    # Set table attributes for better display
    table.align = "l"  # Left-align all columns
    table.valign = "t"  # Top-align all cells
    table.max_width = 30  # Set a maximum width for the 'Value' column

    # Add rows to the table
    for key, value in repo_info.items():
        # Convert the value to string for consistent length checking
        value_str = str(value)

        if isinstance(value, dict):
            value_str = ", ".join(f"{k}: {v}" for k, v in value.items())
        
        # Limit the value display to avoid excessive width
        if len(value_str) > 30:
            value_str = value_str[:27] + '...'
        
        table.add_row([key, value_str])

    print(table)


def display_repositories(repos):
    """Display a list of repositories in a table format."""
    table = PrettyTable()
    table.field_names = ["Repository Name", "Private"]

    for repo in repos:
        table.add_row([repo['name'], repo['private']])

    print(table)

def main():
    parser = argparse.ArgumentParser(description='GitHub API CLI Tool')
    subparsers = parser.add_subparsers(dest='command')

    # Sub-command for repository info
    repo_parser = subparsers.add_parser('repo', help='Repository commands')
    repo_subparsers = repo_parser.add_subparsers(dest='repo_command')

    # Command to get repo info
    repo_info_parser = repo_subparsers.add_parser('info', help='Get repository information')
    repo_info_parser.add_argument('owner', type=str, help='Repository owner')
    repo_info_parser.add_argument('repo', type=str, help='Repository name')

    # Command to list repositories
    repo_list_parser = repo_subparsers.add_parser('list', help='List repositories for a user')
    repo_list_parser.add_argument('owner', type=str, help='User or organization owner')

    args = parser.parse_args()

    if args.command == 'repo' and args.repo_command == 'info':
        info = get_repo_info(args.owner, args.repo)
        display_repo_info(info)

    elif args.command == 'repo' and args.repo_command == 'list':
        repos = list_repositories(args.owner)
        display_repositories(repos)

if __name__ == "__main__":
    main()




# GitHub API CLI Tool

## Overview

The GitHub API CLI Tool is a command-line interface that allows users to interact with the GitHub API seamlessly. This tool provides functionalities to fetch repository information, user profiles, search repositories, manage issues, and more—all from the command line. Designed to simplify interactions with GitHub, this tool is perfect for developers looking to enhance their productivity.

## Features

- **Repository Information**: Retrieve detailed information about a specific repository, including stars, forks, and open issues.
- **User Profile**: Access information about GitHub users, including their public repositories, followers, and following count.
- **Search Repositories**: Search for repositories based on keywords or topics.
- **List Issues**: List open issues in any public repository.
- **Create an Issue**: Create new issues in specified repositories directly from the CLI.
- **Fork a Repository**: Fork a repository to your account with a simple command.
- **Rate Limiting Status**: Check the current rate limit status for API requests.

## Technologies Used

- **Python 3**: Programming language used to implement the CLI tool.
- **Requests Library**: Used to make HTTP requests to the GitHub API.
- **Argparse**: For parsing command-line arguments.
- **PrettyTable**: For displaying data in a structured table format.

## Installation

### Prerequisites

- [Python 3](https://www.python.org/downloads/) (v3.6 or higher)
- [pip](https://pip.pypa.io/en/stable/installation/) (Python package manager)

### Steps

1. **Clone the repository**:
   ```bash
   git clone git@github.com:CovetingAphid3/github-api-cli-tool.git
   ```

2. **Navigate to the project folder**:
   ```bash
   cd github-api-cli-tool
   ```

3. **Install dependencies**:
   ```bash
   pip install -r requirements.txt
   ```

4. **Obtain a GitHub Personal Access Token**:
   - Go to [GitHub Settings](https://github.com/settings/tokens).
   - Click on "Generate new token."
   - Select the necessary scopes based on the functionalities you wish to access (e.g., `repo` for repository access).
   - Copy the token for later use.

5. **Set up your environment**:
   Create a `.env` file in the project root and add your GitHub token:
   ```plaintext
   GITHUB_TOKEN=your_personal_access_token
   ```

## Usage

The tool provides several commands. Here are some examples:

### 1. Get Repository Information
```bash
python main.py repo info <owner> <repo>
```
*Example*:
```bash
python main.py repo info octocat Hello-World
```

### 2. Get User Profile Information
```bash
python main.py user info <username>
```
*Example*:
```bash
python main.py user info octocat
```

### 3. Search for Repositories
```bash
python main.py search repos <keyword>
```
*Example*:
```bash
python main.py search repos machine learning
```

### 4. List Open Issues in a Repository
```bash
python main.py repo issues <owner> <repo>
```
*Example*:
```bash
python main.py repo issues octocat Hello-World
```

### 5. Create an Issue in a Repository
```bash
python main.py repo create_issue <owner> <repo> "<issue_title>" "<issue_body>"
```
*Example*:
```bash
python main.py repo create_issue octocat Hello-World "New Issue" "This is the body of the new issue."
```


### 6. Check Rate Limiting Status
```bash
python main.py rate_limit
```

## Error Handling

The tool includes error handling for various scenarios:
- Invalid input (e.g., incorrect repository or user names).
- API errors (e.g., rate limits exceeded).
- Network errors during API requests.



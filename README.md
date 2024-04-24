
# Git-Auto
Automate Git repository creation and initialization with GitHub integration using this Go CLI tool. To view how this command works visit my Twitter https://twitter.com/SENachi27

### Description
Git Automator is a command-line tool built in Go that simplifies the process of setting up a Git repository with GitHub integration. It automates the steps of creating a GitHub repository, initializing a local Git repository, adding the GitHub repository as a remote origin, committing changes, and pushing them to the remote repository, all in one command.

### Features
- Create a GitHub repository with the name of the current directory.
- Initialize a local Git repository.
- Add the GitHub repository as a remote origin.
- Add all files to the staging area.
- Prompt for a commit message and commit the changes.
- Rename the branch to 'main'.
- Push changes to the remote repository.

### Usage 
- First up clone this repo in your root directory
``` 
    git clone https://github.com/yourusername/git-automator.git

```
- Build the Go executable:
```
cd git-automator
go build -o git-automate main.go
```
- Then add the this Cloned directory path to your environment variables. How to add path to environment variables can be found on this link : http://surl.li/sxrdz

- To use Git Automator, simply run the following command in your terminal:
``` 
    git-auto

```
- Follow the prompts to enter a commit message when prompted.

### Prerequisites
- Go programming language installed
- GitHub CLI (gh) installed and authenticated

### Future Scope

To use an analyzer which analyzes your codebase and then generates a commit message and uses it as the commit message for the commit command.

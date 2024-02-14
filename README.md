# git-branch-rename

Opens the current branch name with your configured `$EDITOR`, to easily change the name without re-typing it.

## Install

```bash
curl -sSfL https://raw.githubusercontent.com/benjlevesque/git-branch-rename/main/godownloader.sh | sh -s -- -b /usr/local/bin
git config --global alias.rename-branch "!git-branch-rename"
```

You can change the installation path by editing `/usr/local/bin` in the above command. Make sure the directory is in your `$PATH`

## Usage
```bash
git rename-branch
```


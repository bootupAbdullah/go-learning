# Understanding PATH Environment Variable

## What PATH Does

**PATH is a shell variable containing directories where executable programs live.**

```bash
echo $PATH
# Output example:
# /usr/bin:/usr/local/bin:/home/user/go/bin
#    ^          ^              ^
#   dir1       dir2           dir3
```

## How Command Execution Works

### When You Type a Command

```bash
$ go version
```

**Shell searches PATH directories in order:**

1. Checks `/usr/bin/go` → not found
2. Checks `/usr/local/bin/go` → found! Execute this
3. (stops searching)

### Without PATH

```bash
# Would need full path every time:
/usr/local/bin/go version
/home/user/go/bin/myprogram

# Instead of just:
go version
myprogram
```

## Your Go PATH Commands Explained

### Command 1: Set GOPATH

```bash
export GOPATH=$HOME/go
```

**What it does:**
- Creates variable `GOPATH` pointing to `~/go`
- Go tooling uses this for workspace location
- Tells Go where your projects/packages live

### Command 2: Update PATH

```bash
export PATH=$PATH:$GOPATH/bin
```

**Breakdown:**

```
$PATH              :  $GOPATH/bin
  ^                      ^
existing PATH      append new directory
```

**Effect:**
- Appends `~/go/bin` to existing PATH
- Programs in `~/go/bin` now executable from anywhere
- Can run Go-compiled binaries without full path

## PATH Syntax

### Colon Separator

```bash
PATH=/dir1:/dir2:/dir3
#         ^     ^
#      separators
```

**Each directory separated by `:`**

### Order Matters

```bash
PATH=/usr/local/bin:/usr/bin
#         ^              ^
#      checked first   checked second
```

**First match wins** - if same program exists in multiple directories, first one found executes.

## Common PATH Operations

### View Current PATH

```bash
# All on one line:
echo $PATH

# Each directory on separate line:
echo $PATH | tr ':' '\n'
```

### Find Where Command Lives

```bash
which go        # Shows full path to go executable
which python3   # Shows full path to python3
type go         # Shows if alias/function/executable
```

### Temporarily Add Directory

```bash
# Only for current shell session:
export PATH=$PATH:/new/directory
```

### Permanently Add Directory

**Edit shell config file:**

```bash
# For bash:
nano ~/.bashrc

# For zsh:
nano ~/.zshrc

# Add line:
export PATH=$PATH:/new/directory

# Reload:
source ~/.bashrc
```

## Mental Model

**PATH = "Phone book of program locations"**

- Shell checks this "phone book" when you type a command
- Finds program's address without you specifying full path
- Searches in order until found

## Quick Reference Commands

```bash
# See PATH contents
echo $PATH

# See PATH formatted
echo $PATH | tr ':' '\n'

# Find executable location
which <command>

# Show all environment variables
env

# Show specific variable
echo $GOPATH
echo $HOME
echo $USER
```

## Common Variables Used With PATH

```bash
$HOME     # User's home directory (/home/username)
$USER     # Current username
$PWD      # Current working directory
$GOPATH   # Go workspace location
```

## Example Scenario

### Before Adding to PATH

```bash
# Have to use full path:
~/go/bin/myapp

# Or navigate to directory first:
cd ~/go/bin
./myapp
```

### After Adding to PATH

```bash
export PATH=$PATH:~/go/bin

# Now works from anywhere:
myapp
```

## Learning Resources

### Immediate Practice

```bash
man bash          # Search for "PATH"
man environ       # Environment variables
env               # See all variables
which <command>   # Find programs
type <command>    # Command info
```

### Experimentation

```bash
# Backup current PATH
OLDPATH=$PATH

# Modify PATH
export PATH=$PATH:/test/directory

# Restore if needed
export PATH=$OLDPATH
```

### Recommended Reading

- "The Linux Command Line" by William Shotts (free PDF)
- `man bash` - Bash manual
- `man environ` - Environment documentation

## Key Takeaways

1. **PATH contains directories** - not individual programs
2. **Colon-separated list** - order matters
3. **Shell searches sequentially** - stops at first match
4. **Executables in PATH directories** - run from anywhere
5. **Add with `export PATH=$PATH:/new/dir`** - appends to existing PATH
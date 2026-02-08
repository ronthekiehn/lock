# lock

Block any website from your terminal. One command to lock, one to unlock.

## How it works

`lock` adds entries to `/etc/hosts` pointing the domain to `0.0.0.0`, then flushes your DNS cache. Simple, effective, immediate.

## Install

```bash
# Clone the repo
git clone https://github.com/yourusername/lock.git

# Add to your PATH (add this to your .zshrc or .bashrc)
alias lock="/path/to/lock/lock"
```

## Usage

```bash
# Block a website
lock x.com

# Unblock a website
lock x.com --unlock
```

Works with any domain:

```bash
lock reddit.com
lock instagram.com
lock youtube.com
```

## Requirements

- macOS (uses `dscacheutil` and `mDNSResponder` for DNS flushing)
- `sudo` access (required to modify `/etc/hosts`)

## License

MIT

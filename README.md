# lock

Block any website from your terminal. One command to lock, one to unlock.

## How it works

`lock` adds entries to `/etc/hosts` pointing the domain to `0.0.0.0`, then flushes your DNS cache. Simple, effective, immediate.

## Install

One-line install:

```bash
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/install.sh | bash
```

Or manually:

```bash
# Download and install
sudo curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/lock -o /usr/local/bin/lock
sudo chmod +x /usr/local/bin/lock
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

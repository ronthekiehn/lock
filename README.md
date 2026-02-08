# lock

Block any website from your terminal. One command to lock, one to unlock.

## How it works

`lock` adds entries to `/etc/hosts` pointing the domain to `0.0.0.0`, then flushes your DNS cache. Simple, effective, immediate.

## Install

**Recommended** (inspect before running):

```bash
# Download the installer
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/install.sh -o install.sh

# Review it (optional but recommended)
cat install.sh

# Run it
bash install.sh
```

**One-line** (for the brave):

```bash
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/install.sh | bash
```

**Install specific version**:

```bash
# Set version (e.g., v1.0.0 or commit hash)
LOCK_VERSION=v1.0.0 bash install.sh
```

**Manual install**:

```bash
# Download as unprivileged user, then install
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/lock -o lock
chmod +x lock
sudo install -m 0755 lock /usr/local/bin/lock
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

- **macOS only** (uses `dscacheutil` and `mDNSResponder` for DNS flushing)
- `sudo` access (required to modify `/etc/hosts`)

Linux/Windows support: The script will exit with an error on non-Darwin systems. PRs welcome to add cross-platform DNS flushing support!

## License

MIT

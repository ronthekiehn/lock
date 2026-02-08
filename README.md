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

**One-line** (if you're feeling brave):

```bash
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/install.sh | bash
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

- Unix-like system with `/etc/hosts` (macOS, Linux, BSD, WSL)
- `sudo` access (required to modify `/etc/hosts`)

**Note:** DNS cache flushing is handled automatically on macOS and most Linux systems. On other systems, you may need to restart your browser for changes to take effect.

## License

MIT

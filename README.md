# lock

Block any website from your terminal. One command to lock, one to unlock.

## How it works

`lock` adds entries to `/etc/hosts` pointing the domain to `0.0.0.0`, then flushes your DNS cache. Simple, effective, immediate.

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/install.sh | sh
```

Or manually:

```bash
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/lock -o lock
chmod +x lock
sudo install -m 0755 lock /usr/local/bin/lock
```

The install script is [open source](https://github.com/ronthekiehn/lock/blob/main/install.sh) - feel free to review it first.

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

# lock

Block distracting websites from your terminal.

## Why?

Browser extensions are too easy to disable. `lock` makes blocking instant but unlocking deliberately hard - you have to manually edit `/etc/hosts` with sudo. Just enough friction to break the habit.

## How it works

Adds entries to `/etc/hosts` pointing domains to `0.0.0.0`, then flushes DNS cache. Works immediately across all browsers.

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

Block any distracting website:

```bash
lock x.com
lock reddit.com
lock youtube.com
lock instagram.com
```

To unblock, you'll need to manually edit `/etc/hosts`:

```bash
sudo nano /etc/hosts
# Find and delete the lines with "# lock: domain.com"
```

The friction is the point - if it was easy to unlock, it wouldn't work.

## Requirements

- Unix-like system with `/etc/hosts` (macOS, Linux, BSD, WSL)
- `sudo` access (required to modify `/etc/hosts`)

**Note:** DNS cache flushing is handled automatically on macOS and most Linux systems. On other systems, you may need to restart your browser for changes to take effect.

## Credits

Built with [Claude](https://claude.ai) (Anthropic).

## License

MIT

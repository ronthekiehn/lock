# lock

A quick script I use to lock Twitter (and other sites) across all browsers. Written by Claude.

The idea: make locking frictionless, but unlocking hard. One command locks a site instantly via `/etc/hosts`. To unlock, you have to manually edit the file with sudo. Just enough friction to break the habit.

There's intentionally no `unlock` command.

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

Lock any distracting website:

```bash
lock x.com
lock reddit.com
lock youtube.com
lock instagram.com
```

To unlock, you'll need to manually edit `/etc/hosts`:

```bash
sudo nano /etc/hosts
# Find and delete the lines with "# lock: domain.com"
```

The friction is the point - if it was easy to unlock, it wouldn't work.

## Requirements

- Unix-like system with `/etc/hosts` (macOS, Linux, BSD, WSL)
- `sudo` access (required to modify `/etc/hosts`)

**Note:** DNS cache flushing is handled automatically on macOS and most Linux systems. On other systems, you may need to restart your browser for changes to take effect.

## Flags

- `-t`, `--kill-terminal`: Close the current terminal session after locking the domain.
- `-j`, `--disable-js`: Disable JavaScript for the domain in Chrome preferences and restart Chrome (best effort).

## Credits

Built with [Claude](https://claude.ai) (Anthropic).

## License

MIT

# lock

Block distracting websites from your terminal. One command, instant effect.

## Why?

We all have those websites that eat our time. Twitter, Reddit, YouTube - you know the ones. The problem isn't accessing them, it's *how easy* it is to access them. One moment of weakness and you're doomscrolling for an hour.

Browser extensions can be disabled with a click. Screen time restrictions feel patronizing. What you need is *friction* - just enough resistance to break the habit loop.

`lock` makes blocking easy and unblocking hard. Run one command and the site is gone. To get it back, you have to manually edit `/etc/hosts` with sudo, find the right lines, and remove them. That's enough friction to make you think twice.

## How it works

`lock` adds entries to `/etc/hosts` pointing the domain to `0.0.0.0`, then flushes your DNS cache. The site becomes unreachable immediately - no browser restart needed. Simple, effective, and OS-level (works across all browsers).

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

## License

MIT

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
lock x.com reddit.com youtube.com
lock -n "finish PR before scrolling" x.com reddit.com
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
- `python3` (optional; used by `-j`/`--disable-js`, state metadata, and duration calculations for `--status`)

**Note:** DNS cache flushing is handled automatically on macOS and most Linux systems. On other systems, you may need to restart your browser for changes to take effect.

## Flags

- `-n`, `--note`: Add a shared note for this lock command. The note is written into `/etc/hosts` lock comments.
- `-s`, `--status`: Show currently active locks, including duration (`locked_for`) and saved note metadata.
- `-t`, `--kill-terminal`: Close the current terminal session after locking the domain.
- `-j`, `--disable-js`: Best-effort: disable JavaScript in Chrome preferences for each provided domain and close Chrome if running.

## State File

`lock --status` reads active locks from `/etc/hosts` and enriches duration/note data from a state file:

- macOS: `~/Library/Application Support/lock/state.json`
- Linux/WSL: `~/.local/state/lock/state.json`

State format:

```json
{
  "state_version": 1,
  "domains": {
    "x.com": {
      "locked_at": "2026-02-15T18:40:00Z",
      "note": "finish checkout work"
    }
  }
}
```

If `state_version` mismatches or the file is corrupt, lock rebuilds state from active `/etc/hosts` lock entries and writes `state_version: 1`.

## Credits

Built with [Claude](https://claude.ai) (Anthropic) and Codex (OpenAI).

## License

MIT

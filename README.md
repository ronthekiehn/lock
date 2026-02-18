# lock

A quick CLI I use to lock Twitter (and other sites) across all browsers.

The idea: make locking frictionless, but unlocking hard. One command locks a site instantly via `/etc/hosts`. To unlock, you have to manually edit the file with sudo. Just enough friction to break the habit.

There's intentionally no `unlock` command.

## Install

From latest GitHub release:

```bash
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/install.sh | sh
```

For local development (installs current working tree binary to `/usr/local/bin/lock`):

```bash
make install
```

`make install` defaults to `/usr/local/bin`; override with `PREFIX=/some/bin make install`.

## Usage

Lock any distracting website:

```bash
lock x.com
lock reddit.com
lock youtube.com
lock instagram.com
lock x.com reddit.com
lock -n "finish PR before scrolling" x.com reddit.com
lock -j x.com
lock --status
lock --version
```

To unlock, you'll need to manually edit `/etc/hosts`:

```bash
sudo nano /etc/hosts
# Find and delete the lines with "# lock: domain.com"
```

The friction is the point - if it was easy to unlock, it wouldn't work.

## Requirements

- macOS
- `sudo` access (required to modify `/etc/hosts`)
- Go 1.22+ (only needed for local builds/install with `make`)

## Flags

- `-n`, `--note`: Add a shared note for this lock command. The note is written into `/etc/hosts` lock comments.
- `-s`, `--status`: Show currently active locks, including duration (`locked_for`) and saved note metadata.
- `-v`, `--version`: Show binary version, installed binary path, and lock state file path.
- `-t`, `--kill-terminal`: Close the current terminal session after locking the domain.
- `-j`, `--disable-js`: Best-effort: disable JavaScript in Chrome preferences for each provided domain and close Chrome if running.

Version format:
- release binaries: release tag version
  - automatic main releases use `v0.0.0-main.<UTC timestamp>.<shortsha>`
  - manual tagged releases can still use semantic versions like `v1.2.3`
- local builds: `0.0.0-dev+<shortsha>` (and `.dirty` when working tree has uncommitted changes)

## State File

`lock --status` reads active locks from `/etc/hosts` and enriches duration/note data from a state file:

- macOS: `~/Library/Application Support/lock/state.json`

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

## Development

Build:

```bash
make build
```

Install locally to your bin directory:

```bash
make install
```

## Release

Pushes to `main` automatically create and publish a GitHub Release via GoReleaser.

Create local release artifacts only (without publishing):

```bash
make release-local
```

`install.sh` is still kept for one-line install from GitHub release binaries.

## Credits

Built with [Claude](https://claude.ai) (Anthropic) and Codex (OpenAI).

## License

MIT

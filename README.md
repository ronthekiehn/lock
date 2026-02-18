# lock

A quick CLI I use to lock Twitter (and other sites) across all browsers.

The idea: make locking frictionless, but unlocking hard. One command locks a site instantly via `/etc/hosts`. 

To unlock, you have to manually edit the file with sudo. I find this is enough friction to make me a bit more intentional about my time.

To this end, there is intentionally no `unlock` command.


## Credits

Built entirely with Claude in Cursor and Codex. I did not touch a line of code here. Codex for some reason does not attribute itself on github.

## Install

From latest GitHub release:

```bash
curl -fsSL https://raw.githubusercontent.com/ronthekiehn/lock/main/install.sh | sh
```

## Usage

Lock any distracting website:

```bash
lock x.com
```

Lock multiple sites at once:
```bash
lock x.com reddit.com
```

Add notes
```bash
lock -n "finish PR" x.com instagram.com
```

Check what is locked and for how long
```bash
lock --status
```

To unlock, you'll need to manually edit `/etc/hosts`:

```bash
sudo nano /etc/hosts
# Find and delete the lines with "# lock: domain.com"
```

The friction is the point - if it was easy to unlock, it wouldn't work.

## Flags

- `-n`, `--note`: Add a note for this lock command. The note is written into `/etc/hosts` lock comments.
- `-s`, `--status`: Show currently active locks, including duration and saved note metadata.
- `-t`, `--kill-terminal`: Close the current terminal session after locking the domain.
- `-j`, `--disable-js`: Best-effort: disable JavaScript in Chrome preferences for each provided domain and close Chrome if running.
- `-v`, `--version`: Show binary version, installed binary path, and lock state file path.


Version format:
- Automatic main releases use incrementing major tags: `v0.0.0`, `v1.0.0`, `v2.0.0`, ...
- I'm too lazy for semver.

## Development

### Requirements

- macOS
- `sudo` access (required to modify `/etc/hosts`)
- Go 1.22+ (only needed for local builds/install with `make`)

### Build

```bash
make build
```

Install locally to your bin directory:

```bash
make install
```

Local builds will have `0.0.0-dev+<shortsha>` (and `.dirty` when working tree has uncommitted changes) for the version number.

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
      "note": "finish work"
    }
  }
}
```

## Release

Pushes to `main` automatically create and publish a GitHub Release via GoReleaser.

Create local release artifacts only (without publishing):

```bash
make release-local
```

`install.sh` is still kept for one-line install from GitHub release binaries.


## License

MIT

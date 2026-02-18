package cli

import (
	"fmt"
	"strings"

	"github.com/ronthekiehn/lock/internal/domain"
)

type Options struct {
	Note         string
	ShowStatus   bool
	ShowVersion  bool
	KillTerminal bool
	DisableJS    bool
	Domains      []string
}

type ParseError struct {
	Message   string
	ShowUsage bool
}

func (e *ParseError) Error() string {
	return e.Message
}

func Usage() string {
	return `Usage: lock [options] <domain> [domain ...]

Options:
  -n, --note <text>     Add a note to lock comments in /etc/hosts
  -s, --status          Show currently locked domains and durations
  -v, --version         Show lock version and key local paths
  -t, --kill-terminal   Close the current terminal session after locking
  -j, --disable-js      Best-effort: disable JavaScript in Chrome for each domain
  -h, --help            Show this help message

Examples:
  lock x.com reddit.com
  lock -n "ship checkout" x.com reddit.com
  lock --kill-terminal x.com
  lock -j x.com
  lock --status
  lock --version

Unlock:
  sudo nano /etc/hosts (remove the lock entries)
`
}

func Parse(args []string) (Options, bool, *ParseError) {
	opts := Options{}
	var domainInputs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-t", "--kill-terminal":
			opts.KillTerminal = true
		case "-j", "--disable-js":
			opts.DisableJS = true
		case "-s", "--status":
			opts.ShowStatus = true
		case "-v", "--version":
			opts.ShowVersion = true
		case "-n", "--note":
			i++
			if i >= len(args) || args[i] == "" {
				return Options{}, false, &ParseError{Message: "Missing value for --note", ShowUsage: true}
			}
			opts.Note = args[i]
		case "-h", "--help":
			return Options{}, true, nil
		case "--":
			domainInputs = append(domainInputs, args[i+1:]...)
			i = len(args)
		default:
			if strings.HasPrefix(arg, "-") {
				return Options{}, false, &ParseError{Message: fmt.Sprintf("Unknown option: %s", arg), ShowUsage: true}
			}
			domainInputs = append(domainInputs, arg)
		}
	}

	if opts.ShowStatus && opts.ShowVersion {
		return Options{}, false, &ParseError{Message: "--status and --version cannot be used together", ShowUsage: true}
	}

	if opts.ShowStatus {
		if len(domainInputs) > 0 {
			return Options{}, false, &ParseError{Message: "--status does not accept domains", ShowUsage: true}
		}
		return opts, false, nil
	}

	if opts.ShowVersion {
		if len(domainInputs) > 0 {
			return Options{}, false, &ParseError{Message: "--version does not accept domains", ShowUsage: true}
		}
		return opts, false, nil
	}

	if len(domainInputs) == 0 {
		return Options{}, false, &ParseError{ShowUsage: true}
	}

	valid, invalid := domain.NormalizeAndClassify(domainInputs)
	if len(invalid) > 0 {
		var b strings.Builder
		b.WriteString("Invalid domain input:\n")
		for _, item := range invalid {
			b.WriteString("  - ")
			b.WriteString(item)
			b.WriteByte('\n')
		}
		return Options{}, false, &ParseError{Message: strings.TrimRight(b.String(), "\n")}
	}

	if len(valid) == 0 {
		return Options{}, false, &ParseError{Message: "No valid domains provided."}
	}

	opts.Domains = valid
	return opts, false, nil
}

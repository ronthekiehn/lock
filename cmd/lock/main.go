package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ronthekiehn/lock/internal/chrome"
	"github.com/ronthekiehn/lock/internal/cli"
	"github.com/ronthekiehn/lock/internal/domain"
	"github.com/ronthekiehn/lock/internal/hosts"
	"github.com/ronthekiehn/lock/internal/state"
	"github.com/ronthekiehn/lock/internal/system"
)

func main() {
	opts, showHelp, parseErr := cli.Parse(os.Args[1:])
	if showHelp {
		fmt.Print(cli.Usage())
		return
	}
	if parseErr != nil {
		if parseErr.Message != "" {
			fmt.Println(parseErr.Message)
		}
		if parseErr.ShowUsage {
			fmt.Print(cli.Usage())
		}
		os.Exit(1)
	}

	if opts.ShowStatus {
		runStatus()
		return
	}

	if err := runLock(opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runStatus() {
	records, err := hosts.LockRecords(hosts.DefaultPath)
	if err != nil {
		fmt.Println("No active locks found.")
		return
	}
	if len(records) == 0 {
		fmt.Println("No active locks found.")
		return
	}

	lines := state.StatusLines(records, time.Now())
	for _, line := range lines {
		fmt.Println(line)
	}
}

func runLock(opts cli.Options) error {
	fmt.Println("🔒 Locking domains...")

	runNow := time.Now().UTC()
	touched := make([]string, 0, len(opts.Domains))
	anyAdded := false

	hasNote := opts.Note != ""
	safeNote := ""
	if hasNote {
		safeNote = domain.SanitizeNote(opts.Note)
	}

	for _, lockDomain := range opts.Domains {
		locked, err := hosts.IsDomainLocked(hosts.DefaultPath, lockDomain)
		if err != nil {
			return err
		}
		if locked {
			fmt.Printf("ℹ️  %s is already locked\n", lockDomain)
			if opts.DisableJS {
				runDisableJS(lockDomain)
			}
			continue
		}

		if err := hosts.AppendLockEntry(hosts.DefaultPath, lockDomain, hasNote, safeNote); err != nil {
			return err
		}
		fmt.Printf("✅ Added %s to /etc/hosts\n", lockDomain)

		anyAdded = true
		touched = append(touched, lockDomain)

		if opts.DisableJS {
			runDisableJS(lockDomain)
		}
	}

	if anyAdded {
		if err := system.FlushDNSCache(); err != nil {
			return err
		}
		fmt.Println("✅ DNS cache flushed")
	}

	records, _ := hosts.LockRecords(hosts.DefaultPath)
	if err := state.SyncAfterLock(touched, opts.Note, runNow, records); err != nil {
		fmt.Println("⚠️  Failed to update lock state file")
	}

	fmt.Println()
	fmt.Println("🔒 Lock command completed")

	if opts.KillTerminal {
		fmt.Println("👋 Closing this terminal session...")
		system.CloseParentTerminal()
	}

	return nil
}

func runDisableJS(lockDomain string) {
	err := chrome.DisableJavaScript(lockDomain)
	if err == nil {
		fmt.Printf("✅ JavaScript disabled for %s in Chrome preferences\n", lockDomain)
		return
	}
	if errors.Is(err, chrome.ErrPreferencesNotFound) {
		fmt.Printf("⚠️  Chrome preferences not found; skipping JavaScript disable for %s\n", lockDomain)
		return
	}
	fmt.Printf("⚠️  Failed to update Chrome preferences; skipping JavaScript disable for %s\n", lockDomain)
}

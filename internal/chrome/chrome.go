package chrome

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

var ErrPreferencesNotFound = errors.New("chrome preferences not found")

func DisableJavaScript(domain string) error {
	prefsPath, err := preferencesPath()
	if err != nil {
		return err
	}

	quitChromeIfRunning()

	content, err := os.ReadFile(prefsPath)
	if err != nil {
		return err
	}

	var data map[string]any
	if err := json.Unmarshal(content, &data); err != nil {
		return err
	}

	profile := ensureMap(data, "profile")
	contentSettings := ensureMap(profile, "content_settings")
	exceptions := ensureMap(contentSettings, "exceptions")
	javascript := ensureMap(exceptions, "javascript")

	setJSLock(javascript, domain)
	setJSLock(javascript, "www."+domain)

	encoded, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	mode := os.FileMode(0o644)
	if info, err := os.Stat(prefsPath); err == nil {
		mode = info.Mode().Perm()
	}

	return os.WriteFile(prefsPath, encoded, mode)
}

func preferencesPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, "Library", "Application Support", "Google", "Chrome", "Default", "Preferences")
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return "", ErrPreferencesNotFound
		}
		return "", err
	}
	return path, nil
}

func quitChromeIfRunning() {
	pgrep := exec.Command("pgrep", "-x", "Google Chrome")
	if err := pgrep.Run(); err != nil {
		return
	}

	osascript := exec.Command("osascript", "-e", `quit app "Google Chrome"`)
	_ = osascript.Run()
}

func ensureMap(parent map[string]any, key string) map[string]any {
	if value, ok := parent[key]; ok {
		if casted, ok := value.(map[string]any); ok {
			return casted
		}
	}
	created := map[string]any{}
	parent[key] = created
	return created
}

func setJSLock(javascript map[string]any, host string) {
	javascript["https://"+host+":443,*"] = map[string]any{"setting": 2}
	javascript["http://"+host+":80,*"] = map[string]any{"setting": 2}
}

package store

import (
	"errors"
	"os"
	"path"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rapidloop/skv"
)

// Store struct
type Store struct {
	Name string
	Path string
}

// New config
func New(name string) (db *skv.KVStore, err error) {
	p, err := Path(name)
	if err != nil {
		return db, err
	}

	// create the directory
	err = os.MkdirAll(p, os.ModePerm)
	if err != nil {
		return db, err
	}

	return skv.Open(path.Join(p, "config.db"))
}

// Path to the storage
func Path(paths ...string) (p string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return p, err
	}

	switch runtime.GOOS {
	case "darwin":
		ps := append([]string{home, "Library", "Preferences"}, paths...)
		return path.Join(ps...), err
	case "linux":
		base := os.Getenv("XDG_CONFIG_HOME")
		if base == "" {
			base = path.Join(home, ".config")
		}
		ps := append([]string{base}, paths...)
		return path.Join(ps...), err
	case "windows":
		appdata := os.Getenv("LOCALAPPDATA")
		if appdata == "" {
			appdata = path.Join(home, "AppData", "Local")
		}
		ps := append([]string{appdata}, paths...)
		ps = append(ps, "Config")
		return path.Join(ps...), err
	default:
		return p, errors.New("store does not yet support " + runtime.GOOS + ". Please open a pull request!")
	}
}

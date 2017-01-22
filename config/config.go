package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	configDir string
	cacheDir  string
)

// GetConfigFolder returns the folder used to store configurations.
func GetConfigDir() string {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		// Create folder
		if err = os.MkdirAll(configDir, os.ModeDir|0700); err != nil {
			fmt.Println("Error when creating config folder:", err)
			return ""
		}
	}
	return configDir
}

// GetCacheFolder returns the folder used to cache.
func GetCacheDir() string {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		// Create cache folder.
		if err = os.MkdirAll(cacheDir, os.ModeDir|0700); err != nil {
			fmt.Println("Error when creating cache folder:", err)
			return ""
		}
	}
	return cacheDir
}

func init() {
	xdgConf := os.Getenv("XDG_CONFIG_HOME")
	if xdgConf == "" {
		xdgConf = os.Getenv("HOME")
		configDir = filepath.Join(xdgConf, ".config", "clever")
	} else {
		configDir = filepath.Join(xdgConf, "clever")
	}
	if xdgConf == "" {
		panic("can't locate user's xdg config folder.")
	}

	xdgCache := os.Getenv("XDG_CACHE_HOME")
	if xdgCache == "" {
		xdgCache = os.Getenv("HOME")
		cacheDir = filepath.Join(xdgCache, ".cache", "clever")
	} else {
		cacheDir = filepath.Join(xdgCache, "clever")
	}
	if xdgCache == "" {
		panic("can't locate user's xdg cache folder.")
	}
}

// Package config_helpers provides helper functions to locate configuration files.
package config_helpers

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/common/file_helpers"
)

func ConfigDir() string {
	if dir := bluemix.EnvConfigDir.Get(); dir != "" {
		return dir
	}
	// TODO: switched to the new default config after all plugin has bumped SDK
	if new := defaultConfigDirNew(); file_helpers.FileExists(new) {
		return new
	}
	return defaultConfigDirOld()
}

// func MigrateFromOldConfig() error {
// 	new := defaultConfigDirNew()
// 	if file_helpers.FileExists(new) {
// 		return nil
// 	}

// 	old := defaultConfigDirOld()
// 	if !file_helpers.FileExists(old) {
// 		return nil
// 	}

// 	if err := file_helpers.CopyDir(old, new); err != nil {
// 		return err
// 	}
// 	return os.RemoveAll(old)
// }

func defaultConfigDirNew() string {
	return filepath.Join(homeDir(), ".ibmcloud")
}

func defaultConfigDirOld() string {
	return filepath.Join(homeDir(), ".bluemix")
}

func homeDir() string {
	if homeDir := bluemix.EnvConfigHome.Get(); homeDir != "" {
		return homeDir
	}
	return UserHomeDir()
}

func TempDir() string {
	return filepath.Join(ConfigDir(), "tmp")
}

func ConfigFile() string {
	return filepath.Join(ConfigDir(), "config.json")
}

func PluginsInstallationDir() string {
	if dir := bluemix.EnvPluginsDir.Get(); dir != "" {
		return dir
	}
	return filepath.Join(ConfigDir(), "plugins")
}

func PluginsInstallationCacheDir() string {
	return filepath.Join(PluginsInstallationDir(), ".cache")
}

func PluginsInstallationConfigFile() string {
	return filepath.Join(PluginsInstallationDir(), "config.json")
}

func PluginBinaryLocation(pluginName string) string {
	binary := filepath.Join(PluginsInstallationDir(), pluginName, pluginName)
	if runtime.GOOS == "windows" {
		binary = binary + ".exe"
	}
	return binary
}

func PluginDataDir(pluginName string) string {
	return filepath.Join(ConfigDir(), "plugins", pluginName)
}

func PluginConfigFile(pluginName string) string {
	return filepath.Join(PluginDataDir(pluginName), "config.json")
}

func CFHome() string {
	return ConfigDir()
}

func CFConfigDir() string {
	return filepath.Join(CFHome(), ".cf")
}

func CFConfigFilePath() string {
	return filepath.Join(CFConfigDir(), "config.json")
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	return os.Getenv("HOME")
}

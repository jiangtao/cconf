package config

import (
	"os"
	"path/filepath"

	"github.com/jiangtao/ccconfig/pkg/i18n"
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Repo        string       `mapstructure:"repo"`
	Projects    []string     `mapstructure:"projects"`
	ScanDirs    []string     `mapstructure:"scanDirs"`
	Lang        string       `mapstructure:"lang"`
	Git         GitConfig    `mapstructure:"git"`
	Backup      BackupConfig `mapstructure:"backup"`
	cliLanguage string       // Language from command line flag
}

// GitConfig holds git-related configuration
type GitConfig struct {
	AutoCommit bool   `mapstructure:"autoCommit"`
	AutoPull   bool   `mapstructure:"autoPull"`
	PushRemote string `mapstructure:"pushRemote"`
}

// BackupConfig holds backup-related configuration
type BackupConfig struct {
	IncludeSettings bool `mapstructure:"includeSettings"`
	IncludeCommands bool `mapstructure:"includeCommands"`
	IncludeSkills   bool `mapstructure:"includeSkills"`
}

var cfg *Config

// Init initializes the configuration
func Init() error {
	cfg = &Config{}

	// Set default values
	setDefaults()

	// Read config file
	if err := readConfigFile(); err != nil {
		// Config file is optional, don't fail if not found
		return nil
	}

	// Initialize i18n
	lang := cfg.Lang
	if cfg.cliLanguage != "" {
		lang = cfg.cliLanguage
	}
	if lang == "" {
		if err := i18n.InitFromEnv(); err != nil {
			return err
		}
	} else {
		if err := i18n.Init(lang); err != nil {
			return err
		}
	}

	return nil
}

// setDefaults sets default configuration values
func setDefaults() {
	viper.SetDefault("repo", "~/cc-config")
	viper.SetDefault("lang", "en")
	viper.SetDefault("scanDirs", []string{
		"~/Places/work",
		"~/Places/personal",
		"~/work",
		"~/projects",
		"~/dev",
	})
	viper.SetDefault("git.autoCommit", false)
	viper.SetDefault("git.autoPull", false)
	viper.SetDefault("git.pushRemote", "origin")
	viper.SetDefault("backup.includeSettings", true)
	viper.SetDefault("backup.includeCommands", true)
	viper.SetDefault("backup.includeSkills", true)
}

// readConfigFile reads the configuration file
func readConfigFile() error {
	// Set config file name and locations
	viper.SetConfigName(".ccconfig")
	viper.SetConfigType("yaml")

	// Search config in home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	viper.AddConfigPath(home)

	// Read config file (optional)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// Unmarshal config
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	return nil
}

// Get returns the current configuration
func Get() *Config {
	return cfg
}

// SetLang sets the language from command line flag
func SetLang(lang string) {
	if cfg == nil {
		cfg = &Config{}
	}
	cfg.cliLanguage = lang
}

// SetRepo sets the repository path from command line flag
func SetRepo(repo string) {
	if cfg == nil {
		cfg = &Config{}
	}
	cfg.Repo = repo
}

// ExpandPath expands a path with ~ to the full home directory path
func ExpandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[1:]), nil
	}
	return path, nil
}

// GetRepoPath returns the expanded repository path
func GetRepoPath() (string, error) {
	return ExpandPath(cfg.Repo)
}

// GetClaudeDir returns the Claude config directory
func GetClaudeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".claude"), nil
}

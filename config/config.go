// Package config
package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"go.yaml.in/yaml/v4"
)

const (
	DirName  = ".config/time-cmd"
	FileName = "time.cfg.yaml"
)

type Config struct {
	Time struct {
		Format string `yaml:"format" env:"TIME_FORMAT"`
	} `yaml:"time"`
}

func New(timeFormat string) Config {
	var cfg Config
	cfg.Time.Format = timeFormat
	return cfg
}

func MustDir() string {
	cfgDir, err := Dir()
	if err != nil {
		panic(fmt.Sprintf("must get config dir: %s", err.Error()))
	}
	return cfgDir
}

func MustParse(fpath string) Config {
	cfg, err := Parse(fpath)
	if err != nil {
		panic(fmt.Sprintf("must parse config: %s", err.Error()))
	}
	return cfg
}

func Dir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("config: dir: user home directory: %w", err)
	}
	cfgDir := filepath.Join(homeDir, DirName)
	return cfgDir, nil
}

func Parse(fpath string) (Config, error) {
	fpath = strings.TrimSpace(fpath)
	fpath = filepath.Clean(fpath)

	f, err := os.Open(fpath)
	if err != nil {
		return Config{}, fmt.Errorf("config: parse %s file: open: %w", fpath, err)
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		if errors.Is(err, io.EOF) {
			// Handle empty configuration files gracefully
			return Config{}, nil
		}
		return Config{}, fmt.Errorf("config: parse %s file: decode yaml: %w", fpath, err)
	}

	return cfg, nil
}

package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		desc   string
		fmtStr string
	}{
		{
			desc:   "date time",
			fmtStr: time.DateTime,
		},
		{
			desc:   "date only",
			fmtStr: time.DateOnly,
		},
		{
			desc:   "time only",
			fmtStr: time.TimeOnly,
		},
		{
			desc:   "RFC3339",
			fmtStr: time.RFC3339,
		},
		{
			desc:   "Kitchen",
			fmtStr: time.Kitchen,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			cfg := New(tt.fmtStr)

			if cfg.Time.Format != tt.fmtStr {
				t.Errorf("expected format %q, got %q", tt.fmtStr, cfg.Time.Format)
			}
		})
	}
}

func TestParse_Valid(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "time-config-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "time.cfg.yaml")
	content := []byte("time:\n  format: \"rfc3339\"\n")
	if err := os.WriteFile(tempFile, content, 0o600); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	cfg, err := Parse(tempFile)
	if err != nil {
		t.Fatalf("unexpected error parsing valid config: %v", err)
	}

	if cfg.Time.Format != "rfc3339" {
		t.Errorf("expected format %q, got %q", "rfc3339", cfg.Time.Format)
	}
}

func TestParse_EmptyFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "time-config-test-empty")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "time.cfg.yaml")
	// Create an empty file
	f, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("failed to create empty file: %v", err)
	}
	f.Close()

	cfg, err := Parse(tempFile)
	if err != nil {
		t.Fatalf("unexpected error parsing empty config: %v", err)
	}

	if cfg.Time.Format != "" {
		t.Errorf("expected format %q, got %q", "", cfg.Time.Format)
	}
}

func TestDir(t *testing.T) {
	dir, err := Dir()
	if err != nil {
		t.Fatalf("unexpected error resolving config directory: %v", err)
	}

	if dir == "" {
		t.Errorf("expected non-empty config directory path")
	}
}

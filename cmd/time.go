package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/Pshimaf-Git/time/config"
)

var (
	timeFormatFl = ""
	initFl       = false
)

var (
	cfgFileName = config.FileName
	cfgDir      = config.MustDir()
	cfgFilePath = filepath.Join(cfgDir, cfgFileName)
)

var matchCodesToTime = map[string]string{
	"default": defaultTimeFormat,

	"ansic":       time.ANSIC,
	"unixdate":    time.UnixDate,
	"rubydate":    time.RubyDate,
	"rfc822":      time.RFC822,
	"rfc822z":     time.RFC822Z,
	"rfc850":      time.RFC850,
	"rfc1123":     time.RFC1123,
	"rfc1123z":    time.RFC1123Z,
	"rfc3339":     time.RFC3339,
	"rfc3339nano": time.RFC3339Nano,
	"kitchen":     time.Kitchen,
	"stamp":       time.Stamp,
	"stampmilli":  time.StampMilli,
	"stampmicro":  time.StampMicro,
	"stampnano":   time.StampNano,
	"datetime":    time.DateTime,
	"dateonly":    time.DateOnly,
	"timeonly":    time.TimeOnly,
}

const defaultTimeFormat = "2006/01/02 15:04 Monday"

func init() {
	flag.StringVar(&timeFormatFl, "format", "", "time out format")
	flag.StringVar(&timeFormatFl, "f", "", "time out format")
	flag.BoolVar(&initFl, "init", false, "create config file")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `time is a command that print current time.

Usage:
  time [flags]

Example:
  time                 ---- print current time with default format. 
  time -f 2006-01-01   ---- print current time with provided format.

Flags:
  -f, --format string   format string
      --init            create config file
  -h, --help            print this help message`)
	}

	flag.Parse()
}

func toCode(s string) string {
	s = strings.ToLower(strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s))

	return s
}

func timeFormat(cfg config.Config) string {
	format := timeFormatFl
	if format == "" {
		format = cfg.Time.Format
	}

	code := toCode(format)
	if tfmt, ok := matchCodesToTime[code]; ok {
		return tfmt
	}

	return format
}

func initConfig() {
	if _, err := os.Stat(cfgFilePath); !os.IsNotExist(err) {
		fmt.Println("Config file already exists. Path:", cfgFilePath)
		return
	}

	err := os.MkdirAll(cfgDir, 0o755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating directories %q: %v\n", cfgDir, err)
		return
	}

	f, err := os.Create(cfgFilePath)
	f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating %q file error: %v\n", cfgFilePath, err)
		return
	}

	fmt.Println("Config file created. Path:", cfgFilePath)
}

func main() {
	if initFl {
		initConfig()
		return
	}

	cfgExists := false
	if _, err := os.Stat(cfgFilePath); !os.IsNotExist(err) {
		cfgExists = true
	}

	cfg := config.New(defaultTimeFormat)

	if cfgExists {
		cfg = config.MustParse(cfgFilePath)
	}

	format := timeFormat(cfg)

	currTime := time.Now()
	s := currTime.Format(format)

	fmt.Println(s)
}

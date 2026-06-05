# Time

`time` is a simple and flexible command-line utility written in Go that helps you print the current time exactly how you want it. Whether you need a quick timestamp, a standard RFC format, or a completely custom layout, `time` has you covered.

## Features

- **Preset Formats:** Comes with a bunch of built-in presets like `RFC3339`, `kitchen`, `datetime`, and more.
- **Custom Layouts:** Use standard Go-style time layout strings (e.g., `"2006-01-02 15:04:05"`) for full control.
- **Persistent Config:** Save your favorite format in a YAML configuration file so you don't have to type it every time.
- **One-Off Overrides:** Quickly override your default settings using the `-f` or `--format` flags.
- **Simple Setup:** Generate a starter configuration file in seconds with the `--init` flag.

## Installation

To build the executable from source, make sure you have [Go](https://go.dev/) installed (version 1.24.4 or later is recommended), then run:

```powershell
go build -o time.exe ./cmd/time.go
```

You can then move `time.exe` to a directory in your PATH to use it globally.

Or you can use `go install`:

```bash
go install github.com/Pshimaf-Git/time/time
```

## Usage

Running the command without any arguments prints the current time using the default format:

```powershell
time
# Output: 2026/06/05 15:04 Friday
```

### Using Presets
You can use one of the many built-in presets:

```powershell
time -f rfc3339
# Output: 2026-06-05T15:04:05Z
```

### Custom Layouts
Provide your own Go-style format string:

```powershell
time -f "15:04:05 (Monday)"
# Output: 15:04:05 (Friday)
```

### Configuration Setup
To create a configuration file:

```powershell
time --init
```

## Configuration

The tool looks for a configuration file at:
`~/.config/time-cmd/time.cfg.yaml`

The structure is simple YAML:

```yaml
time:
  format: "rfc3339"
```

If the file exists, the `format` specified there will be used as the default whenever you run the command without the `-f` flag.

## Supported Preset Codes

The tool is smart enough to ignore casing and spaces (e.g., `RFC 3339` works just like `rfc3339`).

| Code | Layout String / Description |
|------|-----------------------------|
| `ansic` | `Mon Jan _2 15:04:05 2006` |
| `unixdate` | `Mon Jan _2 15:04:05 MST 2006` |
| `rubydate` | `Mon Jan 02 15:04:05 -0700 2006` |
| `rfc822` | `02 Jan 06 15:04 MST` |
| `rfc822z` | `02 Jan 06 15:04 -0700` |
| `rfc850` | `Monday, 02-Jan-06 15:04:05 MST` |
| `rfc1123` | `Mon, 02 Jan 2006 15:04:05 MST` |
| `rfc1123z` | `Mon, 02 Jan 2006 15:04:05 -0700` |
| `rfc3339` | `2006-01-02T15:04:05Z07:00` |
| `rfc3339nano` | `2006-01-02T15:04:05.999999999Z07:00` |
| `kitchen` | `3:04PM` |
| `stamp` | `Jan _2 15:04:05` |
| `stampmilli` | `Jan _2 15:04:05.000` |
| `stampmicro` | `Jan _2 15:04:05.000000` |
| `stampnano` | `Jan _2 15:04:05.000000000` |
| `datetime` | `2006-01-02 15:04:05` |
| `dateonly` | `2006-01-02` |
| `timeonly` | `15:04:05` |
| `default` | `2006/01/02 15:04 Monday` |

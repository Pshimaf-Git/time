# Project Instructions

## Project Overview
`time` is a simple, lightweight command-line utility written in Go for printing the current time with flexible and customizable formats. 

### Key Features
* **Custom & Preset Formats:** Supports a wide variety of standard pre-defined layouts (e.g., `rfc3339`, `kitchen`, `datetime`, `dateonly`) as well as custom Go style time layout strings (e.g., `"2006-01-02 15:04:05"`).
* **Configuration Support:** Supports configuration via a YAML file.
* **CLI Overrides:** Allows overriding the default format on the fly with `-f` or `--format`.
* **Config Initialization:** Features a simple `--init` CLI flag to generate a skeleton configuration file at the default path.

### Tech Stack
* **Language:** Go (1.24.4)
* **Configuration:** YAML format managed by the `go.yaml.in/yaml/v4` library.

### Architecture
The project is structured into two main packages:
1. **`main` (located in `cmd/time.go`):** The program entrypoint. Configures the command-line flags, maps preset codes to layout constants, controls flow execution, and handles user input/outputs.
2. **`config` (located in `config/config.go`):** Handles OS home directory resolution (defining `$HOME/.config/time-cmd/time.cfg.yaml` as the config location) and handles configuration file parsing and decoding.

---

## Building and Running

### Key Commands
* **Build executable:**
  ```powershell
  go build -o time.exe ./cmd/time.go
  ```
* **Run directly from source:**
  ```powershell
  go run ./cmd/time.go
  ```
* **Run tests:**
  *(Note: No tests currently exist; use this command when adding test suites)*
  ```powershell
  go test ./...
  ```

### Usage Examples
* **Print current time using default format:**
  ```powershell
  go run ./cmd/time.go
  # Output: 2026/06/05 15:04 Friday
  ```
* **Print using a preset code:**
  ```powershell
  go run ./cmd/time.go -f rfc3339
  # Output: 2026-06-05T15:04:05Z
  ```
* **Print using a custom layout:**
  ```powershell
  go run ./cmd/time.go -f "15:04:05 (Monday)"
  ```
* **Initialize configuration file:**
  ```powershell
  go run ./cmd/time.go --init
  ```

---

## Configuration Details
* **Config File Path:** `~/.config/time-cmd/time.cfg.yaml` (where `~` is resolved using standard user home directory functions).
* **Format Structure:**
  ```yaml
  time:
    format: "rfc3339"
  ```

### Supported Preset Layout Codes
The tool is case-insensitive and ignores spaces in format codes (e.g., `RFC 3339` is normalized to `rfc3339`):
* `ansic` &rarr; `Mon Jan _2 15:04:05 2006`
* `unixdate` &rarr; `Mon Jan _2 15:04:05 MST 2006`
* `rubydate` &rarr; `Mon Jan 02 15:04:05 -0700 2006`
* `rfc822` &rarr; `02 Jan 06 15:04 MST`
* `rfc822z` &rarr; `02 Jan 06 15:04 -0700`
* `rfc850` &rarr; `Monday, 02-Jan-06 15:04:05 MST`
* `rfc1123` &rarr; `Mon, 02 Jan 2006 15:04:05 MST`
* `rfc1123z` &rarr; `Mon, 02 Jan 2006 15:04:05 -0700`
* `rfc3339` &rarr; `2006-01-02T15:04:05Z07:00`
* `rfc3339nano` &rarr; `2006-01-02T15:04:05.999999999Z07:00`
* `kitchen` &rarr; `3:04PM`
* `stamp` &rarr; `Jan _2 15:04:05`
* `stampmilli` &rarr; `Jan _2 15:04:05.000`
* `stampmicro` &rarr; `Jan _2 15:04:05.000000`
* `stampnano` &rarr; `Jan _2 15:04:05.000000000`
* `datetime` &rarr; `2006-01-02 15:04:05`
* `dateonly` &rarr; `2006-01-02`
* `timeonly` &rarr; `15:04:05`
* `default` &rarr; `2006/01/02 15:04 Monday`

---

## Development Conventions

* **Formatting & Quality:**
  * Always format Go files with `go fmt` prior to checking in.
  * Use standard variable and package naming structures (camelCase, clear package scope).
* **Error Handling:**
  * Follow idiomatic Go error checking.
  * Use `Must` prefix functions (e.g., `MustDir`, `MustParse`) carefully, only when a failure prevents the application from booting up/recovering, and panic is intended.
* **Dependencies:**
  * Kept to a minimum. Run `go mod tidy` to clean up any modifications to the module dependencies.
* **Testing:**
  * Keep unit tests next to source files, following the naming pattern `*_test.go`.
  * Ensure tests are clean and do not leave behind configuration files in user directories (use mock/temp directory mechanisms for file parsing tests).

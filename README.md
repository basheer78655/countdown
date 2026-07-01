# Countdown ⏱️

A terminal countdown timer with large ASCII-art digits, built in Go.

Inspired by [antonmedv/countdown](https://github.com/antonmedv/countdown) — adapted for Linux.

![demo](demo.gif)

## Install

### From source

```bash
go install github.com/basheer-shaik/countdown@latest
```

### From releases

Download a prebuilt binary from [Releases](https://github.com/basheer-shaik/countdown/releases).

### Build from source

```bash
git clone https://github.com/basheer-shaik/countdown.git
cd countdown
go build -o countdown .
```

## Usage

Specify duration in Go format `1h2m3s` or a target time like `02:15pm`, `14:15`.

```bash
countdown 25s
countdown 1m30s
countdown 11:32
countdown 02:15PM
```

### Count up from zero

```bash
countdown -up 30s
```

### Display a title

```bash
countdown -title "Coffee Break" 5m
```

### Voice announcement (last 10 seconds)

Uses `spd-say`, `espeak`, or `espeak-ng` (whichever is available on your system).

```bash
countdown -say 10s
```

### Chain commands

Run a command after the countdown finishes:

```bash
countdown 1m30s && notify-send "Time's up!"
```

## Key Bindings

| Key | Action |
|---|---|
| `Space` | Pause / Resume |
| `Esc` | Quit |
| `Ctrl+C` | Quit |

## Credits

Based on [antonmedv/countdown](https://github.com/antonmedv/countdown) by Anton Medvedev.

## License

[MIT](LICENSE)

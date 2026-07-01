# countdownclilinux ⏱️🎉

A terminal countdown timer with large ASCII-art digits and a **colorful confetti celebration** when time's up! Built in Go.

Inspired by [antonmedv/countdown](https://github.com/antonmedv/countdown) — adapted and enhanced for Linux.

## ✨ Features

- 🔢 **Large ASCII-art digits** centered in your terminal
- 🎉 **Confetti party pop animation** with colorful particles on completion
- ⏸️ **Pause / Resume** with Space key
- ⬆️ **Count up** mode
- 🔊 **Voice announcement** of last 10 seconds (Linux TTS)
- 📝 **Custom title** display below the timer
- 🖥️ **Responsive** — adapts to terminal resize

---

## 📦 Install

### Homebrew (macOS & Linux)

```bash
brew tap basheer78655/countdownclilinux
brew install countdownclilinux
```

### One-liner install (Linux / macOS)

```bash
curl -sSL https://raw.githubusercontent.com/basheer78655/countdown/main/install.sh | sudo bash
```

### Go install

```bash
go install github.com/basheer78655/countdown@latest
```

### From releases

Download a prebuilt binary (`.tar.gz`, `.deb`, `.rpm`) from [Releases](https://github.com/basheer78655/countdown/releases).

### Build from source

```bash
git clone https://github.com/basheer78655/countdown.git
cd countdown
go build -o countdownclilinux .
sudo mv countdownclilinux /usr/local/bin/
```

---

## 🚀 Usage

```bash
# Simple countdown
countdownclilinux 25s
countdownclilinux 1m30s
countdownclilinux 1h2m3s

# Target a specific time
countdownclilinux 14:15
countdownclilinux 02:15PM

# Count up from zero
countdownclilinux -up 30s

# Display a title
countdownclilinux -title "Coffee Break" 5m

# Voice announcement (last 10 seconds)
countdownclilinux -say 10s

# Disable confetti animation
countdownclilinux -noconfetti 10s

# Chain commands
countdownclilinux 1m30s && notify-send "Time's up!"
```

---

## ⌨️ Key Bindings

| Key | Action |
|---|---|
| `Space` | Pause / Resume |
| `Esc` | Quit (no confetti) |
| `Ctrl+C` | Quit (no confetti) |
| Any key during confetti | Skip animation |

---

## 🎉 Confetti Animation

When the countdown reaches zero, a **colorful particle explosion** fills your terminal:
- 80+ particles with physics (velocity + gravity)
- Random Unicode confetti characters: ★ ● ◆ ✦ ♦ ▲ ■ ♥ ✿ ❖
- Color-cycling "TIME'S UP!" banner
- Runs for ~3 seconds, skippable with any key
- Disable with `-noconfetti` flag

---

## Credits

Based on [antonmedv/countdown](https://github.com/antonmedv/countdown) by Anton Medvedev.

## License

[MIT](LICENSE)

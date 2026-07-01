package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

// Confetti characters — a mix of Unicode symbols for a festive look
var confettiChars = []rune{'★', '●', '◆', '✦', '♦', '▲', '■', '♥', '✿', '❖', '✸', '◈', '♠', '♣', '△', '○', '◇', '⬟', '⬡'}

// All available bright termbox colors for maximum vibrancy
var confettiColors = []termbox.Attribute{
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorBlue,
	termbox.ColorMagenta,
	termbox.ColorCyan,
	termbox.ColorWhite,
	termbox.ColorRed | termbox.AttrBold,
	termbox.ColorGreen | termbox.AttrBold,
	termbox.ColorYellow | termbox.AttrBold,
	termbox.ColorBlue | termbox.AttrBold,
	termbox.ColorMagenta | termbox.AttrBold,
	termbox.ColorCyan | termbox.AttrBold,
	termbox.ColorWhite | termbox.AttrBold,
}

// particle represents a single confetti piece with physics
type particle struct {
	x, y   float64 // position
	vx, vy float64 // velocity
	char   rune
	color  termbox.Attribute
	life   int // frames remaining
}

// timesUpBanner is the large "TIME'S UP!" text displayed during confetti
var timesUpBanner = []string{
	"████████╗██╗███╗   ███╗███████╗ ██╗███████╗    ██╗   ██╗██████╗ ██╗",
	"╚══██╔══╝██║████╗ ████║██╔════╝ ╚═╝██╔════╝    ██║   ██║██╔══██╗██║",
	"   ██║   ██║██╔████╔██║█████╗       ███████╗    ██║   ██║██████╔╝██║",
	"   ██║   ██║██║╚██╔╝██║██╔══╝       ╚════██║    ██║   ██║██╔═══╝ ╚═╝",
	"   ██║   ██║██║ ╚═╝ ██║███████╗     ███████║    ╚██████╔╝██║     ██╗",
	"   ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝     ╚══════╝     ╚═════╝ ╚═╝     ╚═╝",
}

// playConfetti runs a ~3 second confetti particle animation
func playConfetti(screenW, screenH int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Spawn particles from center with random radial velocity
	centerX := float64(screenW) / 2.0
	centerY := float64(screenH) / 2.0
	numParticles := 80

	particles := make([]particle, numParticles)
	for i := range particles {
		// Random angle in radians (full 360°)
		angle := rng.Float64() * 2 * math.Pi
		// Random speed — varied for natural spread
		speed := 0.5 + rng.Float64()*2.5

		particles[i] = particle{
			x:     centerX,
			y:     centerY,
			vx:    math.Cos(angle) * speed * 1.8, // wider horizontal spread (terminal chars are taller than wide)
			vy:    math.Sin(angle) * speed,
			char:  confettiChars[rng.Intn(len(confettiChars))],
			color: confettiColors[rng.Intn(len(confettiColors))],
			life:  50 + rng.Intn(50), // 50–100 frames of life
		}
	}

	// Animation loop: ~3 seconds at 30fps = 90 frames
	frameDuration := time.Millisecond * 33 // ~30fps
	totalFrames := 90
	gravity := 0.04 // subtle downward pull

	for frame := 0; frame < totalFrames; frame++ {
		frameStart := time.Now()

		// Clear screen
		_ = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		// Draw the "TIME'S UP!" banner centered
		if len(timesUpBanner) > 0 {
			bannerW := 0
			for _, r := range timesUpBanner[0] {
				_ = r
				bannerW++
			}
			bannerH := len(timesUpBanner)
			bx := screenW/2 - bannerW/2
			by := screenH/2 - bannerH/2

			// Cycle banner color for a pulsing effect
			bannerColor := confettiColors[frame%len(confettiColors)]

			for row, line := range timesUpBanner {
				col := 0
				for _, r := range line {
					if bx+col >= 0 && bx+col < screenW && by+row >= 0 && by+row < screenH {
						termbox.SetCell(bx+col, by+row, r, bannerColor, termbox.ColorDefault)
					}
					col++
				}
			}
		}

		// Update and draw particles
		for i := range particles {
			p := &particles[i]
			if p.life <= 0 {
				continue
			}

			// Physics update
			p.x += p.vx
			p.vy += gravity // gravity pulls down
			p.y += p.vy

			// Decrease life
			p.life--

			// Draw if on screen
			px, py := int(math.Round(p.x)), int(math.Round(p.y))
			if px >= 0 && px < screenW && py >= 0 && py < screenH {
				termbox.SetCell(px, py, p.char, p.color, termbox.ColorDefault)
			}
		}

		_ = termbox.Flush()

		// Check for key press to skip animation
		select {
		case ev := <-queues:
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter {
				return
			}
		default:
			// no input, continue animation
		}

		// Frame timing
		elapsed := time.Since(frameStart)
		if elapsed < frameDuration {
			time.Sleep(frameDuration - elapsed)
		}
	}
}

package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/alecthomas/kong"
	cliutils "github.com/thewizardplusplus/motivator/cli-utils"
)

type cli struct {
	Start      startCommand         `kong:"cmd,help='Start (or restart) showing notifications in background.'"` // nolint: lll
	Status     statusCommand        `kong:"cmd,help='Check that notifications are being show in background.'"`  // nolint: lll
	Stop       cliutils.StopCommand `kong:"cmd,help='Stop showing notifications in background.'"`               // nolint: lll
	Foreground foregroundCommand    `kong:"cmd,help='Start showing notifications in foreground.'"`              // nolint: lll
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, err := kong.Must(&cli{}).Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("unable to parse a CLI: %v", err)
	}

	if err := ctx.Run(); err != nil {
		log.Fatalf("unable to process a CLI: %v", err)
	}
}

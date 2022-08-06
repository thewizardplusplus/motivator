package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/alecthomas/kong"
	cliutils "github.com/thewizardplusplus/motivator/cli-utils"
)

// nolint: lll
type cli struct {
	Start      startCommandWrapper    `kong:"cmd,help='Start (or restart) displaying notifications in the background.'"`
	Status     cliutils.StatusCommand `kong:"cmd,help='Check that notifications are being display in the background.'"`
	Stop       cliutils.StopCommand   `kong:"cmd,help='Stop displaying notifications in the background.'"`
	Foreground foregroundCommand      `kong:"cmd,help='Start displaying notifications in the foreground.'"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, err := kong.Must(&cli{}).Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("unable to parse the arguments of the CLI: %s", err)
	}

	if err := ctx.Run(); err != nil {
		log.Fatalf("unable to run the selected command: %s", err)
	}
}

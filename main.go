package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/alecthomas/kong"
)

type cli struct {
	Foreground foregroundCommand `kong:"cmd,help='Start showing notifications in foreground.'"` // nolint: lll
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

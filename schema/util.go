package schema

import (
	"log"
	"time"
)

// TimeTrack verifies  how loog takes a function
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

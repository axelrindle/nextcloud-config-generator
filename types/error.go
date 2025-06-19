package types

import "log"

type Error struct {
	Err     error
	Context string
}

func (e Error) Throw() {
	log.Printf("Failed task: %s\n", e.Context)
	log.Fatal(e.Err)
}

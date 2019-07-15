package main

import (
	log "github.com/pysta3515/logace"
)

func main() {
	// use the default golang log
	log.Print("test Print:", 1, 2, 3)
	log.Println("test Println", 4, 5, 6)
	log.Printf("test Printf: %d, %d, %d", 7, 8, 9)
}

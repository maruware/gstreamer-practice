package main

import (
	"log"

	practice "github.com/maruware/gstreamer-practice"
)

func main() {
	cmd := practice.MosaicCmd(4, 3, 300, 200)
	log.Printf("cmd: %v", cmd.String())
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

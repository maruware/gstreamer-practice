package main

import (
	"log"
	"os"

	practice "github.com/maruware/gstreamer-practice"
)

func main() {
	u := os.Getenv("RTMP_URL")
	if len(u) <= 0 {
		log.Fatal("undefined RTMP_URL")
	}
	cmd := practice.Rtmp(u)
	log.Printf("cmd: %v", cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

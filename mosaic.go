package practice

import (
	"fmt"
	"os/exec"
)

func MosaicCmd(nx int, ny int, width int, height int) *exec.Cmd {
	args := []string{"-e", "videomixer", "name=mix"}
	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			n := x + y*nx
			args = append(args, fmt.Sprintf("sink_%v::xpos=%v", n, x*width))
			args = append(args, fmt.Sprintf("sink_%v::ypos=%v", n, y*height))
		}
	}

	args = append(args, "!", "autovideosink")

	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			n := x + y*nx
			pattern := n % 24
			args = append(args, "videotestsrc", fmt.Sprintf("pattern=%v", pattern))
			args = append(args, "!", fmt.Sprintf("video/x-raw,width=%v,height=%v", width, height))
			args = append(args, "!", fmt.Sprintf("mix.sink_%v", n))
		}
	}

	return exec.Command("gst-launch-1.0", args...)
}

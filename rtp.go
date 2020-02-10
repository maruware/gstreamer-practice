package practice

import (
	"fmt"
	"os/exec"
)

func Publisher(port int) *exec.Cmd {
	args := []string{
		"videotestsrc",
		"!", "x264enc",
		"!", "rtph264pay",
		"!", "udpsink", "host=127.0.0.1", fmt.Sprintf("port=%v", port), "sync=false"}
	return exec.Command(exe, args...)
}

func Receiver(port int) *exec.Cmd {
	args := []string{
		fmt.Sprintf("udpsrc port=%v", port),
		"!", "application/x-rtp,media=video,encoding-name=H264",
		"!", "queue",
		"!", "rtph264depay",
		"!", "avdec_h264",
		"!", "videoconverter",
		"!", "autovideosink",
	}
	return exec.Command(exe, args...)
}

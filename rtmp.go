package practice

import (
	"fmt"
	"os/exec"
)

func Rtmp(dst string) *exec.Cmd {
	args := []string{
		"videotestsrc",
		"!", "videoconvert",
		"!", "video/x-raw,width=1280,height=720,framerate=25/1",
		"!", "queue",
		"!", "x264enc", "bitrate=2000", "byte-stream=false", "key-int-max=60", "bframes=0", "aud=true", "tune=zerolatency",
		"!", "video/x-h264,profile=main",
		"!", "mux.",
		"audiotestsrc",
		"!", "audioconvert",
		"!", "faac",
		"!", "mux.",
		"flvmux", "streamable=true", "name=mux",
		"!", "rtmpsink", fmt.Sprintf("location=%s", dst),
	}
	return exec.Command(exe, args...)
}

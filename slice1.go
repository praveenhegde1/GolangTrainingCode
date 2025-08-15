package main

import (
	"fmt"
	"time"
)

func main() {

	videoFrames := []string{
		"Frame 1",
		"Frame 2",
		"Frame 3",
		"Frame 4",
		"Frame 5",
		"Frame 6",
	}

	playVideoFrames(videoFrames)

	pauseVideo(3)

	playVideoFrames(videoFrames)
}

func playVideoFrames(frames []string) {
	for _, frame := range frames {
		fmt.Println("Playing:", frame)
		time.Sleep(1 * time.Second)
	}
}

func pauseVideo(seconds int) {
	fmt.Println("Pausing video...")
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Resuming video...")
}

package main

import (
	"github.com/gustavomello-21/poc-face-recognition-go/src/detection"
	"github.com/gustavomello-21/poc-face-recognition-go/src/recognitions"
)

func main() {
	recognitions.FaceDetection()

	detection.DetectionFace()
}

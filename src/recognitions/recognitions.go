package recognitions

import (
	"log"
	"path/filepath"

	"github.com/leandroveronezi/go-recognizer"
)

const fotosDir = "./img/"
const dataDir = "models"

func FaceDetection() {
	rec := recognizer.Recognizer{}
	err := rec.Init(dataDir)
	if err != nil {
		log.Fatal(err)
		return
	}

	rec.Tolerance = 0.4
	rec.UseGray = true
	rec.UseCNN = false
	defer rec.Close()

	faces, err := rec.RecognizeMultiples(filepath.Join(fotosDir, "detection_face.png"))
	if err != nil {
		log.Fatal(err)
		return
	}

	img, err := rec.DrawFaces2(filepath.Join(fotosDir, "detection_face.png"), faces)
	if err != nil {
		log.Fatal(err)
		return
	}

	rec.SaveImage("./img/result.jpeg", img)
}

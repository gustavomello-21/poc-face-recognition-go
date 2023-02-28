package detection

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/leandroveronezi/go-recognizer"
)

const fotosDir = "img"
const dataDir = "models"

func addFile(rec *recognizer.Recognizer, Path, Id string) {

	err := rec.AddImageToDataset(Path, Id)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func DetectionFace() {
	rec := recognizer.Recognizer{}
	err := rec.Init(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	rec.Tolerance = 0.4
	rec.UseCNN = false
	rec.UseGray = true
	defer rec.Close()

	addFile(&rec, filepath.Join(fotosDir, "base_picture.png"), "Mabppe")

	rec.SetSamples()

	face, err := rec.ClassifyMultiples(filepath.Join(fotosDir, "detection_face.png"))
	if err != nil {
		log.Fatal(err)
		return
	}

	img, err := rec.DrawFaces(filepath.Join(fotosDir, "detection_face.png"), face)
	if err != nil {
		log.Fatal(err)
		return
	}

	rec.SaveImage("./img/final_result.jpg", img)
}

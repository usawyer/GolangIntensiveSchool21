package main

import (
	img "day06/internal/img_generator"
	"log"
)

func main() {
	image := img.NewImage(300, 300)
	err := image.GenerateAmazingLogo()
	if err != nil {
		log.Println(err)
	}

	log.Println("Amazing logo has been generated successfully!")
}

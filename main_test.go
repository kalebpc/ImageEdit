package main

import (
	"fmt"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	imageedit := Imageedit{pixels: 33}
	for j := 0; j < b.N; j += 1 {
		// for i := 0; i < reflect.ValueOf(&imageedit).NumMethod()-1; i += 1 {
		// function := reflect.TypeOf(&imageedit).Method(i).Name
		function := "PIX"
		arguments := []string{"./assets/dino.png", "./new.png", function}
		processImage(arguments, imageedit)
		fmt.Println("function: ", arguments[2])
		fmt.Print("\n")
		// }
	}
}
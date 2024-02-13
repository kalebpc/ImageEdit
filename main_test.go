package main

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	imageedit := Imageedit{pixels: 33}
	for j := 0; j < b.N; j += 1 {
		for i := 0; i < reflect.ValueOf(&imageedit).NumMethod(); i += 1 {
			function := reflect.TypeOf(&imageedit).Method(i).Name
			arguments := []string{"./assets/dino.png", "./new.png", function}
			processImage(arguments, imageedit)
			fmt.Println("function: ", arguments[2])
			fmt.Print("\n")
		}
	}
}

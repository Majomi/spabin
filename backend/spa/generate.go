// +build ignore

package main

import (
	"net/http"
	"log"
	"github.com/shurcooL/vfsgen"
)

func main() {
	
	assets := http.Dir("static")
	err := vfsgen.Generate(assets, vfsgen.Options{
		PackageName:  "spa",
		Filename:     "spa/assets.go",
		VariableName: "assets",
	})

	if err != nil {
		log.Fatal(err)
	}

}
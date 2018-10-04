package main

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/loader"
)

func main() {
	var conf loader.Config
	conf.FromArgs([]string{"./thepackage"}, false)
	program, err := conf.Load()
	if err != nil {
		panic(err)
	}

	for _, pkg := range program.Imported {
		for name, def := range pkg.Info.Defs {
			if name.Name == "SuperInterface" {
				ms := types.NewMethodSet(def.Type())
				for i := 0; i < ms.Len(); i++ {
					fmt.Println(ms.At(i).Obj())
				}
			}
		}
	}
}

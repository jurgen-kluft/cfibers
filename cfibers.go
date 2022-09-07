package main

import (
	ccode "github.com/jurgen-kluft/ccode"
	cpkg "github.com/jurgen-kluft/cfibers/package"
)

func main() {
	ccode.Init()
	ccode.Generate(cpkg.GetPackage())
}

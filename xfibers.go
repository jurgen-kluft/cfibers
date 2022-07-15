package main

import (
	"github.com/jurgen-kluft/xcode"
	xpackage "github.com/jurgen-kluft/xfibers/package"
)

func main() {
	xcode.Init()
	xcode.Generate(xpackage.GetPackage())
}

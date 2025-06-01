package cfibers

import (
	cbase "github.com/jurgen-kluft/cbase/package"
	"github.com/jurgen-kluft/ccode/denv"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

// GetPackage returns the package object of 'cfibers'
func GetPackage() *denv.Package {
	// Dependencies
	unittestpkg := cunittest.GetPackage()
	basepkg := cbase.GetPackage()

	// The main (cfibers) package
	mainpkg := denv.NewPackage("github.com\\jurgen-kluft", "cfibers")
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(basepkg)

	// 'cfibers' library
	mainlib := denv.SetupCppLibProject(mainpkg, "cfibers")
	mainlib.AddDependencies(basepkg.GetMainLib()...)

	// 'cfibers' unittest project
	maintest := denv.SetupCppTestProject(mainpkg, "cfibers_test")
	maintest.AddDependencies(unittestpkg.GetMainLib()...)
	maintest.AddDependencies(basepkg.GetMainLib()...)
	maintest.AddDependency(mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)

	return mainpkg
}

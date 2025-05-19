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
	mainpkg := denv.NewPackage("cfibers")
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(basepkg)

	// 'cfibers' library
	mainlib := denv.SetupCppLibProject("cfibers", "github.com\\jurgen-kluft\\cfibers")
	mainlib.AddDependencies(basepkg.GetMainLib()...)

	// 'cfibers' unittest project
	maintest := denv.SetupDefaultCppTestProject("cfibers_test", "github.com\\jurgen-kluft\\cfibers")
	maintest.AddDependencies(unittestpkg.GetMainLib()...)
	maintest.AddDependencies(basepkg.GetMainLib()...)
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)

	return mainpkg
}

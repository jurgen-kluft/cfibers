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
	mainlib := denv.SetupDefaultCppLibProject("cfibers", "github.com\\jurgen-kluft\\cfibers")
	mainlib.Dependencies = append(mainlib.Dependencies, basepkg.GetMainLib())

	// 'cfibers' unittest project
	maintest := denv.SetupDefaultCppTestProject("cfibers_test", "github.com\\jurgen-kluft\\cfibers")
	maintest.Dependencies = append(maintest.Dependencies, unittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, basepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)

	return mainpkg
}

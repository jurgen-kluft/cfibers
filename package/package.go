package xfibers

import (
	xbase "github.com/jurgen-kluft/xbase/package"
	"github.com/jurgen-kluft/xcode/denv"
	xentry "github.com/jurgen-kluft/xentry/package"
	xunittest "github.com/jurgen-kluft/xunittest/package"
)

// GetPackage returns the package object of 'xfibers'
func GetPackage() *denv.Package {
	// Dependencies
	unittestpkg := xunittest.GetPackage()
	entrypkg := xentry.GetPackage()
	basepkg := xbase.GetPackage()

	// The main (xfibers) package
	mainpkg := denv.NewPackage("xfibers")
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(entrypkg)
	mainpkg.AddPackage(basepkg)

	// 'xfibers' library
	mainlib := denv.SetupDefaultCppLibProject("xfibers", "github.com\\jurgen-kluft\\xfibers")
	mainlib.Dependencies = append(mainlib.Dependencies, basepkg.GetMainLib())

	// 'xfibers' unittest project
	maintest := denv.SetupDefaultCppTestProject("xfibers_test", "github.com\\jurgen-kluft\\xfibers")
	maintest.Dependencies = append(maintest.Dependencies, unittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, entrypkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, basepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)

	return mainpkg
}

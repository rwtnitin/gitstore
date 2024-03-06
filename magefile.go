//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

func Run() error {
	mg.Deps(Build)
	return sh.RunV("./gitstore")
}

func Build() error {
	return sh.RunV("go", "build", "-tags=mage", "-o", "gitstore", ".")
}

func Clean() {
	os.RemoveAll("gitstore")
}

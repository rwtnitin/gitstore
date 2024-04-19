//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

func Build() error {
	mg.Deps(TailwindBuild)
	mg.Deps(TemplGenerate)
	return sh.RunV("go", "build", "-o", "dist/gitstore", "./cmd/main.go")
}

func TailwindWatch() error {
	return sh.RunV("tailwindcss", "-i", "web/static/css/input.css", "-o", "web/static/css/styles.css", "--watch")
}

func TailwindBuild() error {
	return sh.RunV("tailwindcss", "-i", "web/static/css/input.css", "-o", "web/static/css/styles.css", "--minify")
}

func TemplWatch() error {
	return sh.RunV("templ", "generate", "--watch")
}

func TemplGenerate() error {
	return sh.RunV("templ", "generate")
}
func Clean() {
	os.RemoveAll("gitstore")
}

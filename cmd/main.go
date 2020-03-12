package main

import (
	"github.com/nemotoy/salomrn"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(salomrn.Analyzer) }

package main

import (
	"github.com/nemotoy/salomrn"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(salomrn.Analyzer) }

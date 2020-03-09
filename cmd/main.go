package main

import (
	"github.com/nemotoy/mreclen"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(mreclen.Analyzer) }

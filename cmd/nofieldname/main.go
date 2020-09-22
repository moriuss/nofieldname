package main

import (
	"github.com/moriuss/nofieldname"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(nofieldname.Analyzer) }

package main

import (
	"github.com/sonatard/noctx"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(noctx.Analyzer) }

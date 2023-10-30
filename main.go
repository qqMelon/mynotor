package main

import (
	"flag"
	"github.com/qqMelon/mynotor/call"
)

var (
	Url string
)

func init() {
	flag.StringVar(&Url, "url", "", "Check authentication url")
	flag.Parse()
}

func main() {
	call.CheckUrl(Url)
}

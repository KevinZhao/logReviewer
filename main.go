package main

import (
	"flag"

	"github.com/zhaokm/sleep/util"
)

func main() {

	folderPath := flag.String("o", "/Users/zhaokm/Workspace/sleep/", "folder path")
	flag.Parse()
	util.FolderIteration_new(*folderPath)
}

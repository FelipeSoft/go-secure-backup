package main

import (
	"fmt"
	"github.com/FelipeSoft/go-secure-backup/internal/agent/capture"
)

func main() {
	fmt.Println("Agent Process is running")
	// go func(){
		fmt.Println("capturing files and folders")
		capture.CaptureFilesAndFolders()
	// }()
}
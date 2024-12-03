package main

import (
	"fmt"
	"time"

	// "github.com/FelipeSoft/go-secure-backup/internal/agent/capture"
	// "github.com/FelipeSoft/go-secure-backup/internal/agent/storage"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {
	godotenv.Load("../../.env")
	fmt.Println("Agent Process is running")
	fmt.Println("capturing files and folders")
	// storageMethod := storage.NewBareMetalStorage()
	// captureFilesAndFolders := capture.NewCaptureFilesAndFolders(storageMethod)
	// captureFilesAndFolders.Execute()

	c := cron.New()
	c.AddFunc("@every 00h00m2s", task)
	c.Start()
	select {}
}

func fixZero() {

}

func task() {
	now := time.Now()
	fmt.Println(now.UnixNano())
}
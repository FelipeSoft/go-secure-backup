package capture

import (
	"os"
	"fmt"
	"runtime"
	"log"
	"github.com/FelipeSoft/go-secure-backup/internal/agent/platform"
)

func CaptureFilesAndFolders() {
	operationalSystem := runtime.GOOS
	path := "C:/Users/felip/OneDrive/Área de Trabalho"
	content := platform.PlatformFactory(operationalSystem).GetContentFromPath(path)
	for c := 0; c < len(content); c++ {
		fmt.Println(content[c])
	}
	file, err := os.OpenFile("C:/Users/felip/OneDrive/Área de Trabalho/its_ok.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error on logging: %s", err.Error())
	}
	file.Write([]byte("Hello World!"))
}
package platform

import (
	"github.com/FelipeSoft/go-secure-backup/internal/agent/platform/strategy"
)

type PlatformStrategy interface {
	GetContentFromPath(path string) []string
}

func PlatformFactory(os string) PlatformStrategy {
	if os == "windows" {
		return &platform.WindowsStrategy{}
	}
	return nil
}

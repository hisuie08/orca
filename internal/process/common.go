package process

import (
	"orca/internal/ostools"
	"os"
)

func findCompose(path string) string {
	supportedName := []string{
		"compose.yaml",
		"compose.yml",
		"docker-compose.yml",
		"docker-compose.yaml",
	}
	for _, name := range supportedName {
		composeFile := path + "/" + name
		if ostools.FileExisists(composeFile) {
			return composeFile
		}
	}
	return ""
}

func CommonProcess(path string) {
	if path == "" {
		wd, _ := os.Getwd()
		path = wd

	}
	directories := ostools.Directories(path)
	for _, directly := range directories {
		findCompose(directly) //TODO: implement
		//ParseCompose(composeFile)
	}
}

package utils

import (
	"errors"
	"os"

	"github.com/spf13/afero"
)

func GetConfigDirectory() (string, error) {
	appConfigDir, err := os.UserHomeDir()
	if err != nil {
		panic("could not get user config directory")
	}

	if _, err := os.Stat(appConfigDir + "/vgtracker"); os.IsNotExist(err) {
		dirErr := os.Mkdir(appConfigDir+"/vgtracker", 0777)
		if dirErr != nil {
			panic("could not create config directory")
		}
	}

	return appConfigDir + "/vgtracker", nil
}

func FileExists(appFS afero.Fs, nameWithPath string) bool {
	_, err := appFS.Stat(nameWithPath)
	if err != nil {
		return false
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func Patch[T any](incoming *T, defaultValue T) *T {
	if incoming == nil {
		return &defaultValue
	}

	return incoming
}

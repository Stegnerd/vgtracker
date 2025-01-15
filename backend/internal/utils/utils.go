package utils

import "os"

func GetConfigDirectory() (string, error) {
	appConfigDir, err := os.UserConfigDir()
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

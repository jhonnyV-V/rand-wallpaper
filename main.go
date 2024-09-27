package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
)

func setImage(image, path string) error {
	cmd := exec.Command("hsetroot", "-cover", path+"/"+image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Failed to get config directory: %s\n", err)
		return
	}

	configPath := userConfigPath + "/rand-wallpaper"

	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(configPath, 0766)
			if err != nil {
				fmt.Printf("Failed to create rand-wallpaper directory: %s\n", err)
				return
			}
		} else {
			fmt.Printf("Failed to get rand-wallpaper directory: %s\n", err)
			return
		}
	}

	images, err := os.ReadDir(configPath)
	if err != nil {
		fmt.Printf("Failed to read rand-wallpaper directory: %s\n", err)
		return
	}

	if len(images) == 0 {
		fmt.Printf("no images in %s directory\n", configPath)
		return
	}

	if len(images) == 1 {
		image := images[0].Name()
		err = setImage(image, configPath)
		if err != nil {
			fmt.Printf("Failed to set image %s: %s\n", image, err)
			return
		}
		return
	}

	randImage := rand.IntN(len(images))
	image := images[randImage].Name()
	err = setImage(image, configPath)
	if err != nil {
		fmt.Printf("Failed to set image %s: %s\n", image, err)
		return
	}
}

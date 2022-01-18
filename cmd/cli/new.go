package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

var appURL string

func doNew(appName string) {
	appName = strings.ToLower(appName)
	appURL = appName

	// sanitize the application name (convert url to single word)
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[len(exploded)-1]
	}

	log.Println("App name is", appName)

	// git clone the skeleton application
	color.Green("\tCloning repository...")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/ck3g/gwf-template.git",
		Progress: os.Stdout,
		Depth:    1,
	})

	if err != nil {
		exitGracefully(err)
	}

	// remove .git directory
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		exitGracefully(err)
	}

	// create a ready to go .env file
	color.Yellow("\tCreating .env file...")
	data, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		exitGracefully(err)
	}

	env := string(data)
	env = strings.ReplaceAll(env, "${APP_NAME}", appName)
	env = strings.ReplaceAll(env, "${KEY}", g.RandomString(32))

	err = copyDataToFile([]byte(env), fmt.Sprintf("./%s/.env", appName))
	if err != nil {
		exitGracefully(err)
	}

	// create a Makefile
	// Optional: Create one for windows, by checking `runtime.GOOS == "windows"`
	source, err := os.Open(fmt.Sprintf("./%s/Makefile", appName))
	if err != nil {
		exitGracefully(err)
	}
	defer source.Close()

	destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
	if err != nil {
		exitGracefully(err)
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		exitGracefully(err)
	}
	// If using separate Makefiles
	// os.Remove("./" + appName + "/Makefile.windows")

	// update the go.mod file
	color.Yellow("\tCreating go.mod file...")
	data, err = os.ReadFile("./" + appName + "/go.mod")
	if err != nil {
		exitGracefully(err)
	}

	mod := string(data)
	mod = strings.ReplaceAll(mod, "test", appURL)
	err = os.WriteFile("./"+appName+"/go.mod", []byte(mod), 0)
	if err != nil {
		exitGracefully(err)
	}

	// update the existing .go files with correct name and imports
	color.Yellow("\tUpdating source files...")
	os.Chdir("./" + appName)
	updateSource()

	// run go mod tidy in the project directory
}

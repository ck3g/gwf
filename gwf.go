package gwf

import (
	"fmt"

	"github.com/joho/godotenv"
)

const version = "0.0.1"

type GWF struct {
	AppName string
	Debug   bool
	Version string
}

func (g *GWF) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env file
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (g *GWF) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := g.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GWF) checkDotEnv(path string) error {
	err := g.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

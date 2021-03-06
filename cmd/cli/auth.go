package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

func doAuth() error {
	// migrations
	dbType := g.DB.DatabaseType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := g.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := g.RootPath + "/migrations/" + fileName + ".down.sql"

	log.Println(dbType, upFile, downFile)

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("DROP TABLE IF EXISTS users CASCADE; DROP TABLE IF EXISTS tokens CASCADE; DROP TABLE IF EXISTS remember_tokens;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy files over
	err = copyFileFromTemplate("templates/data/user.go.txt", g.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", g.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/remember_token.go.txt", g.RootPath+"/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over middleware
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", g.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", g.RootPath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/remember.go.txt", g.RootPath+"/middleware/remember.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy handlers over
	err = copyFileFromTemplate("templates/handlers/auth-handlers.go.txt", g.RootPath+"/handlers/auth-handlers.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy views over
	err = copyFileFromTemplate("templates/mailer/password-reset.html.tmpl", g.RootPath+"/mail/password-reset.html.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/mailer/password-reset.plain.tmpl", g.RootPath+"/mail/password-reset.plain.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/login.jet", g.RootPath+"/views/login.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/forgot.jet", g.RootPath+"/views/forgot.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/reset-password.jet", g.RootPath+"/views/reset-password.jet")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow("	- users, tokens, and remember_tokens migrations created and executed")
	color.Yellow("	- users, and tokens models created")
	color.Yellow("	- auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.go, and to appropriate middleware to your routes!")

	return nil
}

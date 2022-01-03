package main

import (
	"fmt"
	"time"
)

func doSessionTable() error {
	dbType := g.DB.DatabaseType

	if dbType == "mariadb" {
		dbType = "mysql"
	}

	if dbType == "postgresql" {
		dbType = "postgres"
	}

	fileName := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro())
	upFile := g.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := g.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/"+dbType+"_sessions.sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("DROP TABLE sessions;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

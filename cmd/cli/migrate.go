package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// run the migration command
	switch arg2 {
	case "up":
		err := g.MigrateUp(dsn)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" {
			err := g.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := g.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}

	case "reset":
		err := g.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = g.MigrateUp(dsn)
		if err != nil {
			return err
		}

	case "default":
		showHelp()
	}

	return nil
}

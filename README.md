1. create the dbconfig.yml in config folder and set the db information.
2. run `sql-migrate new -config=config/dbconfig.yml your_desired_migration_name`.
3. put the sql statement under the "up" section, fill in the "down" section as well.
4. run `go run cmd/dbmigrate/main.go` to execute the migrations.
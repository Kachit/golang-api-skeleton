package commands_database

const (
	DatabaseMigrationsMigrate  string = "database:migrations:migrate"
	DatabaseMigrationsRollback string = "database:migrations:rollback"
	DatabaseSeedersSeed        string = "database:seeders:seed"
	DatabaseSeedersClear       string = "database:seeders:clear"
)

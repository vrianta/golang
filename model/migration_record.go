package model

var MigrationName string = Date.string() + "_" + Time.string() + "_" + "migration.sql"

// migrationName is the file name
func recordMigrationSchema(queries string) {
	// check if the migrationName file exists
	if !FileExists(MigrationName) {
		// create the file
		CreateFile(MigrationName)
	}
	// append the queries to the file
	AppendToFile(MigrationName, queries)
}

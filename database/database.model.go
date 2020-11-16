package database

// Database ...
type Database interface {
	Set(key, value string) (string, error)
	Get(key string) (string, error)
	Delete(key string) (string, error)
}

// Factory looks up acording to the databaseName the database implementation
func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "redis":
		return createRedisDatabase()
	default:
		return nil, &NotImplementedDatabaseError{databaseName}
	}
}

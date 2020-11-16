package database

// Database ...
type Database interface {
	Set(key, value string) (string, error)
	Get(key string) (string, error)
	Delete(key string) (string, error)
}

// Factory takes the name of a database implementaion to try and create, or errors
// if one has not been implemented. We can easily plug mongo or postgres in here
// as long as they satisfy the interface above.
func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "redis":
		return createRedisDatabase()
	default:
		return nil, &NotImplementedDatabaseError{databaseName}
	}
}

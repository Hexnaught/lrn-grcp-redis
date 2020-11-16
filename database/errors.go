package database

// OperationError when can not perform given operation on DB (SET, GET, DELETE)
type OperationError struct {
	operation string
}

func (err *OperationError) Error() string {
	return "Could not perform the " + err.operation + " operation."
}

// DownError when it is not a redis.Nil response or DB is down
type DownError struct{}

func (dbe *DownError) Error() string {
	return "Database is down"
}

// CreateDatabaseError when cannot perform set on DB
type CreateDatabaseError struct{}

func (err *CreateDatabaseError) Error() string {
	return "Could not create database"
}

// NotImplementedDatabaseError when user tries to create a not implemented DB
type NotImplementedDatabaseError struct {
	database string
}

func (err *NotImplementedDatabaseError) Error() string {
	return err.database + " not implemented"
}

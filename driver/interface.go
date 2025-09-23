package driver

type Driver interface {
	Open(connStr string) (Driver, error)
}

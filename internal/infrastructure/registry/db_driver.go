package registry

import (
	"crystal-cli/internal/domain"
	"crystal-cli/internal/infrastructure/db/sqlite"
)

var dbDrivers = map[string]domain.DatabaseDriver{
	"sqlite": &sqlite.SQLiteDriver{},
}

func GetDBDriver(name string) (domain.DatabaseDriver, error) {
	drv, ok := dbDrivers[name]
	if !ok {
		return nil, domain.ErrUnsupportedDB
	}
	return drv, nil
}

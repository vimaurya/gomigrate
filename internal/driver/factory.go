package driver

import (
	"fmt"
	"strings"
)

func GetDriver(connURL string) (Driver, error) {
	if strings.HasPrefix(connURL, "postgres://") || strings.HasPrefix(connURL, "postgresql://") {
		return NewPostgresDriver(connURL)
	}

	if strings.HasPrefix(connURL, "mysql://") {
		dsn := strings.TrimPrefix(connURL, "mysql://")
		return NewMySQLDriver(dsn)
	}

	return nil, fmt.Errorf("unsupported database scheme name in url : %s", connURL)
}

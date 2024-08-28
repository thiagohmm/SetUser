// infrastructure/database/oracle.go
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	"github.com/spf13/viper"
)

func ConectarBanco() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s/%s@%s", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_CONNECTSTRING"))
	return sql.Open("godror", dsn)
}

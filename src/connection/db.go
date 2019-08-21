package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func GetDBConnection(driver string) (*sql.DB, error) {

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		viper.GetString("APP.DB.USERNAME"),
		viper.GetString("APP.DB.PASSWORD"),
		viper.GetString("APP.DB.HOST"),
		viper.GetString("APP.DB.PORT"),
		viper.GetString("APP.DB.DATABASE"),
	)

	db, err := sql.Open(driver, dbUrl)

	db.SetMaxIdleConns(0)

	return db, err
}

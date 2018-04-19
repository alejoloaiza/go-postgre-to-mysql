package extra

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var MySqldb *sql.DB
var Postgredb *sql.DB
var err error

func DBConnectPostgres() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Localconfig.PostgresDBHost, Localconfig.PostgresDBPort, Localconfig.PostgresDBUser, Localconfig.PostgresDBPass, Localconfig.PostgresDBName)
	//fmt.Println(psqlInfo)
	Postgredb, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//fmt.Println(">>>>>>>>>>>>>>>>> Successfully connected to Database <<<<<<<<<<<<<<<<<")
}
func DBConnectMySQL() {
	psqlInfo := fmt.Sprintf("%s:%s@/%s", Localconfig.MySQLDBUser, Localconfig.MySQLDBPass, Localconfig.MySQLDBName)
	//fmt.Println(psqlInfo)
	MySqldb, err = sql.Open("mysql", psqlInfo)

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	err = MySqldb.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
func DBClosePostgres() {
	Postgredb.Close()
}
func DBCloseMySQL() {
	MySqldb.Close()
}

package main

import (
	"go-postgre-to-mysql/extra"
	"go-postgre-to-mysql/migrationdb"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	arg := "./extra/config.json"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	_ = extra.GetConfig(arg)
	rand.Seed(time.Now().UTC().UnixNano())
	extra.DBConnectPostgres()
	extra.DBConnectMySQL()
	migrationdb.SelectAndInsert()
	extra.DBClosePostgres()
	extra.DBCloseMySQL()
}

package extra

import "encoding/json"
import "os"
import "fmt"

type Configuration struct {
	MySQLDBUser    string
	MySQLDBPass    string
	MySQLDBName    string
	PostgresDBHost string
	PostgresDBPort string
	PostgresDBUser string
	PostgresDBPass string
	PostgresDBName string
}

var Localconfig *Configuration

func GetConfig(configpath string) *Configuration {
	file, _ := os.Open(configpath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	Localconfig = &configuration
	return &configuration
}

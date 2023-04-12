package flags

import (
	"flag"
	"fmt"
	"github.com/peter-yen/p-builder/pkg/global"
)

func ParseFlags() (driver, dir, folderName string) {

	// default postgres.go
	flag.StringVar(&driver, "driver", "postgres.go", "driver name (default: postgres.go)")

	// required
	flag.StringVar(&dir, "dir", "", "database connection dir")

	// default model
	flag.StringVar(&folderName, "folder", "model", "folder name (default: model)")

	flag.Parse()

	if dir == "" {
		global.Log.Fatalln("error: dir is required!")
		return
	}

	// validate driver name

	if // github.com/lib/pq
	driver != "postgres.go" &&
		// github.com/go-sql-driver/mysql
		driver != "mysql" &&
		// github.com/mattn/go-sqlite3
		driver != "sqlite3" &&
		// github.com/denisenkom/go-mssqldb
		driver != "mssql" &&
		// github.com/mattn/go-oci8
		driver != "oci8" &&
		// github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql
		driver != "cloudsqlmysql" &&
		// github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres.go
		driver != "cloudsqlpostgres" {
		global.Log.Fatalln("error: driver name is invalid!")
		return
	}

	fmt.Println(driver, dir, folderName)

	return
}

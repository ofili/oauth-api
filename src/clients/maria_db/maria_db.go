package maria_db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)


func init() {
	// Connect to MariaDB cluster
	db, err := sql.Open("mysql", "1Chi@goziem:@/oauth")
	if err != nil {
		panic(err)
	}

	fmt.Println("maria db connection established successfully")
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

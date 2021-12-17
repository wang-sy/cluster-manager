package main

import (
	"log"
	"time"

	"github.com/wuhuizuo/sqlm"

	_ "github.com/go-sql-driver/mysql"
)

// Table table
//
// swagger:model table
type Table struct {

	// id
	ID uint64 `json:"id,omitempty" db:"id,type=INT,auto_increment,primary,key"`

	// n cols
	NCols uint64 `json:"n_cols,omitempty" db:"n_cols,type=INT"`

	// name
	Name string `json:"name,omitempty" db:"name,type=VARCHAR(32),not_null"`
}

// TableRawModel return Table obj for sqlm framework.
func TableRawModel() interface{} {
	return &Table{}
}

func main() {
	dbConn := &sqlm.Database{
		Driver: "mysql",
		DSN:    "root:123456@tcp(mysql-service:3306)/MYSQL_PROJ",
	}

	if err := dbConn.Create(); err != nil {
		log.Println("can't conn", err)
	}

	foobarTable := &sqlm.Table{
		Database:  dbConn,
		TableName: "foobar",
	}
	foobarTable.SetRowModel(TableRawModel)
	log.Println(foobarTable.Create())

	a, b := foobarTable.Insert(&Table{NCols: 2, Name: "test"})
	log.Println("insert: ", a, b)

	// sleep.
	time.Sleep(time.Hour)
}

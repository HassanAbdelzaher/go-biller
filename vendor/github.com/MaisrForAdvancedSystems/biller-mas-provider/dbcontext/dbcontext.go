package dbcontext

import (
	"github.com/HassanAbdelzaher/lama"
	_ "github.com/HassanAbdelzaher/lama/dialects/mssql"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbpool"
)

var DbConnPool *lama.Lama

func init() {
	conn,err:=dbpool.GetConnection()
	if err!=nil{
		panic(err)
	}
	DbConnPool=conn
}

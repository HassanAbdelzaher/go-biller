package dbpool

import (
	"context"
	"errors"
	"fmt"
	"github.com/MaisrForAdvancedSystems/mas-db-models/config"
	"github.com/MaisrForAdvancedSystems/mas-db-models/migration"
	"github.com/spf13/viper"
	"log"
	"time"

	"github.com/HassanAbdelzaher/lama"
	_ "github.com/HassanAbdelzaher/lama/dialects/mssql"
	_ "github.com/MaisrForAdvancedSystems/mas-db-models/config" //important to activate config
)


var _dbConnectionPool *lama.Lama


func GetConnection() (*lama.Lama,error){
	if _dbConnectionPool==nil{
		err:=_init();
		if err!=nil{
			return nil,err
		}
		if _dbConnectionPool==nil{
			return nil,errors.New("no database connection pool avilable");
		}
	}
	return _dbConnectionPool,nil;
}

func _init() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			err=errors.New(fmt.Sprintf("error:%v",r))
		}
	}()
	str:=viper.GetString(config.MasDbConnName)
	debug:=viper.GetBool("debug")
	if str==""{
		return errors.New("invalied connection string")
	}
	if debug{
		str=str+ ";log=63"
	}
	_dbConnectionPool, err = lama.Connect("sqlserver", str)
	if err!=nil{
		return err
	}
	if _dbConnectionPool==nil{
		return errors.New("error while creating db connection pool")
	}
	_dbConnectionPool.Debug = false
	log.Println("connecting...")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_dbConnectionPool.DB.PingContext(ctx)
	poolSize:=viper.GetInt(config.MaxDbConnections)
	if poolSize<1{
		poolSize=100
	}
	poolIdleSize:=viper.GetInt(config.MaxDbIdleConnections)
	if poolIdleSize<1{
		poolIdleSize=10
	}
	_dbConnectionPool.DB.SetMaxOpenConns(poolSize)
	_dbConnectionPool.DB.SetMaxIdleConns(poolIdleSize)
	log.Println("connectd to database")
	err=migration.Migrate(_dbConnectionPool)
	return err
}


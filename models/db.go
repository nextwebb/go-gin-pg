package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //PostgreSQL Driver
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
// boilerplate configurations
func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=pgpass dbname=sampledb host=postgres_docker sslmode=disable")
	orm.RegisterModel(new(Users))
	ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() orm.Ormer {
	return ormObject
}

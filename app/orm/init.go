package orm 

import (
	"test_sql/infra"
)

func InitOrmInstances() {
	User = InitUserOrm(infra.GetDB().DB)
	//orm.Order = orm.Order.InitWithDB(infra.GetDB().DB)
	//orm.Point = orm.Point.InitWithDB(infra.GetDB().DB)
	// ....
}
package main

import (
	"github.com/foxsuagr-sanse/go-gobang_game/app/model"
	"github.com/foxsuagr-sanse/go-gobang_game/common/db"
)


func main() {
	var d db.DB = &db.SetData{}
	dblink := d.MySqlInit()
	defer dblink.Close()
	var UserGroup []*model.UserGroups
	dblink.Where("uid = 2003 AND group_rank = 1 AND user_group = '损友'").Find(&UserGroup)
	println(len(UserGroup))
}
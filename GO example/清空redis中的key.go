package main

import (
	"github.com/foxsuagr-sanse/go-gobang_game/common/db"
)

func main() {
	var d db.DB = &db.SetData{}
	dblink := d.RedisInit(0)
	defer dblink.Close()
	_,err := dblink.FlushAll().Result()
	if err != nil {
		panic(err)
	}
}
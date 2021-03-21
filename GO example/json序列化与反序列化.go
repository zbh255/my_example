package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//将json字符串 反序列化 到map集合
	jsonMapStr := `{"1":{"goodsName":"华为P40","price":"ok","color":"蓝色"},"2":{"goodsName":"华为P40","price":"no","color":"蓝色"}}`
	ResMap := map[string]map[string]string{}
	err2 := json.Unmarshal([]byte(jsonMapStr), &ResMap)
	if err2 != nil {
		fmt.Println("Unmarshal jsonMap err!", err2)
		return
	}
	fmt.Println("↓map集合↓")
	fmt.Println(ResMap)
	fmt.Println(ResMap["1"]["price"])
	fmt.Println()

	mapJSONRes, err2 := json.Marshal(ResMap)
	if err2 != nil {
		fmt.Println("map集合序列化为json失败...", err2)
		return
	}
	fmt.Println("↓map集合序列化为json后:↓")
	fmt.Println(string(mapJSONRes))
	fmt.Println()
}

package hello

import "fmt"

func main05(){
	var map0 = make(map[string]string)
	map0["key01"] = "value01"
	map0["key02"] = "value02"
	map0["key03"] = "value03"
	map0["key04"] = "value04"
	map0["key05"] = "value05"

	for key := range map0 {
		var value = map0[key]
		fmt.Println("key: ", key, ", value: ", value)
		if(key == "key03"){
			delete(map0, "key03")
		}
	}

	//启动一个新的线程
	go hello()
}

func hello(){
	fmt.Println("hello go")
}
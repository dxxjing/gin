package main

import (
	"encoding/json"
	esclient "gin-test/es/esClient"
)

type User struct {
	Name string `json:"name"`
	Age int	`json:"age"`
}

func main() {

	bulk()
	//add()

}

func bulk(){
	//批量插入 一行操作 一行数据
	/*data := `
		{"create":{ "_index":"jdx1","_type":"cc"}}
		{"name":"jdx2","age" : 30}

		{"create":{ "_index":"jdx1","_type":"cc"}}
		{"name":"jdx3","age" : 40}
`*/
	data := []User{
		{Name: "test1",Age: 13},
		{Name: "test2", Age: 14},
	}
	query := ""
	for _,v := range data {
		query += "{\"create\":{ \"_index\":\"jdx\",\"_type\":\"cc\"}}\n"
		str, _ := json.Marshal(v)
		query += string(str) + "\n"
	}
	//fmt.Println(query)
	esclient.HttpPost(query, "http://localhost:9200/_bulk")
}

func add() {
	//单个插入
	/*data := `
			{
				"name":"jdx",
				"age" : 20
			}
	`*/
		data := User{
			Name: "tom",
			Age: 12,
		}
		query,_ := json.Marshal(data)
		esclient.HttpPost(string(query), "http://localhost:9200/jdx/cc")
}
package main

import esclient "gin-test/es/esClient"

func main() {
	query := `
			{
	    "query":{
	        "bool":{
	            "must":{
	                "wildcard":{
	                    "name":{
	                        "value":"*j*"
	                    }
	                }
	            }
	        }
	    }
	}
	`
		esclient.HttpPost(query, "http://localhost:9200/jdx/cc/_search")
}

package esclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HttpPost(query, url string) {
	cli := http.DefaultClient

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	rsp , err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	defer rsp.Body.Close()
	data, _ := ioutil.ReadAll(rsp.Body)
	if rsp.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		os.Exit(-3)
	}
	fmt.Println(string(data))
}


func HttpGet() {
	cli := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9200/jdx/cc/_search", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	rsp , err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	defer rsp.Body.Close()
	data, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(data))
}
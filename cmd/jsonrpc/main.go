package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service struct {
	Port uint32
	// HandlerList []Foo `json:"handler_list"`
	RewriteList []struct {
		TestString         string // `json:"test_string"`
		ConditionPattern   string // `json:"condition_pattern"`
		AddServiceIpToRule bool   // `json:"add_service_ip_to_rule"`
		HandlerUrl         string // `json:"handler_url"`
		RedirectUrl        string // `json:"redirect_url"`
	} // `json:"rewrite_list"`
	RedirectList []struct {
		HandlerUrl  string // `json:"handler_url"`
		RedirectUrl string // `json:"redirect_url"`
	} // `json:"redirect_list"`
}

func main() {
	payload := map[string]interface{}{
		".oid":    "ClusterManager",
		".method": "get_service_list",
		".kwargs": map[string]interface{}{},
	}

	client := &http.Client{}
	buf, err := json.Marshal(payload)
	// err handling, etc.
	res, err := client.Post("http://localhost:2100/jsonrpc", "application/json",
		bytes.NewReader(buf))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer res.Body.Close()

	var respData struct {
		Response []Service `json:".return"`
	}
	response, _ := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(response, &respData); err != nil {
		fmt.Printf("Error: %v: %s\n", err, string(response))
		return
	}
	fmt.Printf("Successful response: %+v from {{%s}}\n", respData.Response,
		string(response))
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TEST = "http://10.20.6.79:9088"
	PROD = ""
	UAT = "http://10.20.5.102:9088"
	CONTENT_TYPE = "application/json;charset=UTF-8"
)
const (
	GET_CODE = "/system/getCode"
	LOGIN = "/system/login"
	QUERY_ACTION_ALL = "/action/queryByUser"
	DEL_ACTION = "/action/delete"
)

func delAction(TEST string, body delBody) {
	var v = make(map[string]interface{}, 1)
	v["data"] = map[string]string{
		"id": body.Id,
	}
	fmt.Println(v)
	jsonBody, _ := json.Marshal(v)
	fmt.Println("string(jsonBody)", string(jsonBody))
	ctx, err := http.NewRequest("POST", formUrl + DEL_ACTION, bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println(err)
	}
	ctx.Header.Set("Content-Type", CONTENT_TYPE)
	ctx.Header["token"] = []string{TOKEN}
	ctx.Header["userId"] = []string{USER_ID}
	resp, err := (&http.Client{}).Do(ctx)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请求错误")
	}
	var res Response3
	if err := json.Unmarshal([]byte(content), &res); err == nil {
		fmt.Println(res)
	}
}

func queryMenuAll() {
	var v = make(map[string]interface{}, 1)
	v["data"] = map[string]string {
		"userId": USER_ID,
	}
	jsonBody, _ := json.Marshal(v)
	ctx, err := http.NewRequest("POST", formUrl + QUERY_ACTION_ALL, bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println("请求错误")
	}
	ctx.Header.Set("Content-Type", CONTENT_TYPE)
	ctx.Header["token"] = []string{TOKEN}
	ctx.Header["userId"] = []string{USER_ID}
	resp, err := (&http.Client{}).Do(ctx)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请求错误")
	}
	var res Response1
	if err := json.Unmarshal([]byte(content), &res); err == nil {
		if res.Result == 200 {
			parse(res.Data)
		}
	} else {
		fmt.Println(err)
	}
}

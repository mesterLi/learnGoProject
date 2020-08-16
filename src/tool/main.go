package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"
)

var formUrl, toUrl, TOKEN, USER_ID string

func toMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	fmt.Println(data)
	return data
}

func getRandomString(length int) string {
	tempArr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var token string
	for i := 0; i < length; i++ {
		randNum := math.Floor(rand.Float64() * float64(length))
		token += tempArr[int(randNum)]
	}
	return token
}

func createCode(acc *Account) {
	token := getRandomString(32)
	body := Any{}
	body["data"] = Any {
		"token": token,
	}
	jsonAccount, _ := json.Marshal(body)
	reader := bytes.NewReader(jsonAccount)
	response, err := http.Post(formUrl + GET_CODE, CONTENT_TYPE, reader)
	if err != nil {
		return
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	var res Response
	if err := json.Unmarshal([]byte(content), &res); err == nil {
		//fmt.Println(res)
		acc.Code = res.Data
		acc.Token = token
	}
}

func login(account *Account) int {
	//var body = Any{}
	//body["data"] = toMap(*account)
	//jsonBody, _ := json.Marshal(body)
	//var ccc interface{}
	//if err := json.Unmarshal([]byte(jsonBody), &ccc); err == nil {
	//	fmt.Println(ccc)
	//}
	//reader := bytes.NewReader(jsonBody)
	//response, err := http.Post(formUrl + LOGIN, CONTENT_TYPE, reader)
	//if err != nil {
	//	fmt.Println(err)
	//	return 500
	//}
	//defer response.Body.Close()
	//content, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return 500
	//}
	//var res Response1
	//if err := json.Unmarshal([]byte(content), &res); err == nil {
	//	fmt.Println(res.Data["token"])
	//	fmt.Println(res.Data["id"])
	//} else {
	//	fmt.Println(err)
	//}
	return 200
}

func parse(ml MenuList) {
	for _, v := range ml {
		if v.Id == "F06AE7C175454C5D3C84499BB12FE562" {
			fmt.Println(v.Id)
		}
	}
}

func getEnv()  {
	var env string
	if len(os.Args) > 1 {
		env = strings.ToUpper(os.Args[len(os.Args) - 1])
	}
	switch env {
		case "":
			formUrl = TEST
		case "TEST":
			formUrl = TEST
		case "UAT":
			formUrl = UAT
		case "PROD":
			formUrl = PROD
	}
	fmt.Println(formUrl)
}

func startServe()  {
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("server running at port 23333")
}
func checkFileIsExist(filename string) bool {
	exit := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exit = false
	}
	return exit
}
func writeAuth()  {
	//var f *os.File
	fmt.Println("Token", TOKEN == "")
	if TOKEN == "" {
		return
	}
	UserHomeDir, _ := os.UserHomeDir()
	authInfo := map[string]string {
		"Token": TOKEN,
		"UserId": USER_ID,
	}
	filePath := UserHomeDir + "/.web_openservice"
	stringInfo, _ := json.Marshal(authInfo)
	fmt.Println(string(stringInfo))
	if isExist := checkFileIsExist(filePath); isExist {
		ioutil.WriteFile(filePath, []byte(stringInfo), 0777)
		fmt.Println("写入成功～")
	} else {
		fmt.Println("文件不存在～")
		if _, err := os.Create(filePath); err == nil {
			ioutil.WriteFile(filePath, []byte(stringInfo), 0777)
		}  else {
			fmt.Println("err", err)
		}
	}
}
func main() {
	//var id, password string
	fmt.Println(os.Args)
	//getEnv()
	//http.HandleFunc("/", func(w http.ResponseWriter, f *http.Request) {
	//	//info := f.Header.Get("Access-Control-Request-Headers")
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	w.Header().Set("Access-Control-Allow-Headers", "userid, token")
	//	Token, Userid := f.Header.Get("Token"), f.Header.Get("Userid")
	//	TOKEN = Token
	//	USER_ID = Userid
	//	writeAuth()
	//	fmt.Println(os.UserHomeDir())
	//	fmt.Println(Token, Userid)
	//	w.Write([]byte("hello world"))
	//})
	//err := http.ListenAndServe(":7777", nil)
	//if err != nil {
	//	fmt.Println("err", err)
	//	return
	//}
	//exec.Command("open", "http://localhost:8080").Start()
}

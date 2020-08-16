package main

type Account struct {
	Id string `json:"id"`
	Code string `json:"code"`
	Password string `json:"password"`
	Token string `json:"token"`
}

type Any map[string]interface{}

type Response struct {
	Data string `json:"data"`
	Result int `json:"result"`
	Total int `json:"total"`
}

type Menu struct {
	Url string `json:"url"`
	Name string `json:"name"`
	Target string `json:"target"`
	delBody
}

type MenuList []Menu

type Response1 struct {
	Data MenuList `json:"data"`
	Result int `json:"result"`
	Total int `json:"total"`
}

type delBody struct {
	Id string `json:"id"`
}

type Response3 struct {
	Data bool `json:"data"`
	Result int `json:"result"`
	Message string `json:"message"`
}

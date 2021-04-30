package model

type Users struct {
	Id       int
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Age      int    `json:"age"`
	IsMale   bool   `json:"is_male"`
}

type Table struct {
	ID 		int
	Name 	string 	`json:"name"`
	Age	 	int		`json:"age"`
	Job		string	`json:"job"`
}

type LocalCache struct{
	Cache	interface{} `json:"cache"`
	IsActual bool `json:"is_actual"`
}
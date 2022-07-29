package models

type Response struct {
	Error 	bool			`json:"error"`
	Data 	interface{}		`json:"data"`
	Msg		string			`json:"msg"`
	Errors	error			`json:"errors"`
}
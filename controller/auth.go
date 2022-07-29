package controller

import (
	"net/http"
	"encoding/json"
)

func Signup(w http.ResponseWriter, r *http.Request){
	output := make(map[string]int)
	output["ade"]=1
	output["bisi"]=2
	json.NewEncoder(w).Encode(output)
}

func Signin(w http.ResponseWriter, r *http.Request){
	output := make(map[string]int)
	output["ade"]=1
	output["bisi"]=2
	json.NewEncoder(w).Encode(output)
}
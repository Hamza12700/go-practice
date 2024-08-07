package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type UserData struct {
	A string `json:"a"`
	B string `json:"b"`
}

type Result struct {
	Total int
}

func Calculate(w http.ResponseWriter, r *http.Request, operator string) {
	decoded := json.NewDecoder(r.Body)
	userReq := &UserData{}
	if err := decoded.Decode(userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	num1 := userReq.A
	num2 := userReq.B

	parsedInt, err := strconv.Atoi(num1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	parsedInt2, err := strconv.Atoi(num2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var data = Result{}
	switch operator {
	case "+":
		data.Total = parsedInt + parsedInt2
	case "-":
		data.Total = parsedInt - parsedInt2
	case "*":
		data.Total = parsedInt * parsedInt2
	case "/":
		if parsedInt == 0 {
			http.Error(w, "Can't divide by 0", http.StatusBadRequest)
			return
		}
		data.Total = parsedInt / parsedInt2
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(jsonData)
}

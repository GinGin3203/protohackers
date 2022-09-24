package main

import "encoding/json"

type MyNumber struct {
	json.Number
	Set bool
}

func (m *MyNumber) UnmarshalJSON(data []byte) error {
	m.Set = true
	return json.Unmarshal(data, &m.Number)
}

type Request struct {
	Method string   `json:"method"`
	Number MyNumber `json:"number"`
}

func (r *Request) isWellFormed() bool {
	if r.Method != "isPrime" || !r.Number.Set {
		return false
	}
	return true
}

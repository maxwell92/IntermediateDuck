package main

import (
	"contract"
	"router"
)

const (
	CONTRACT_FILE = "./contract.json"
)

func main() {
	c := contract.NewContract(CONTRACT_FILE)
	r := router.NewRouter(c)
	r.RegistAndRun()
}

// Prepare:
// 1. Router: read and regist URL from contract file
// 2. Get Data: read from contract file and validate the json data
// 3. Post Data: read from contract file and validate the json data

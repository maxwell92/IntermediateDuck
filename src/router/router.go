package router

import (
	"contract"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Router struct {
	ContractList *contract.ContractList
}

func NewRouter(c *contract.ContractList) *Router {
	r := new(Router)
	r.ContractList = c
	return r
}

func (r Router) RegistAndRun() {
	for _, c := range r.ContractList.Contracts {
		ctr := new(contract.Contract)
		ctr.URL = c.URL
		ctr.Get = c.Get
		ctr.Post = c.Post

		http.HandleFunc(ctr.URL, func(w http.ResponseWriter, req *http.Request) {
			switch req.Method {
			case "GET":
				func(w http.ResponseWriter, req *http.Request) {
					fmt.Fprintln(w, ctr.Get.String())
				}(w, req)
			case "POST":
				func(w http.ResponseWriter, req *http.Request) {
					body, err := ioutil.ReadAll(req.Body)
					if err != nil {
						panic(err)
					}
					expected := ctr.Post
					fmt.Fprintln(w, string(body) == expected)
				}(w, req)
			case "OPTIONS":
				func(w http.ResponseWriter, req *http.Request) {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					fmt.Fprintln(w, "")
				}(w, req)
			}
		})
	}

	fmt.Println("Running...")
	http.ListenAndServe(":8080", nil)
}

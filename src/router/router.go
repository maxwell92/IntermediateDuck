package router

import (
	"contract"
	"fmt"
	"net/http"
)

type Router struct {
	Contract *contract.Contract
}

func NewRouter(c *contract.Contract) *Router {
	r := new(Router)
	r.Contract = c
	return r
}

func (r Router) RegistAndRun() {
	http.HandleFunc(r.Contract.URL, r.Handle)
	fmt.Println("Running...")
	http.ListenAndServe(":8080", nil)
}

func (r Router) Handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		r.Get(w)
	case "POST":
		r.Post(w)
	}
}

func (r Router) Get(w http.ResponseWriter) {
	// fmt.Fprintf(w, "Hello Get\n")
	fmt.Fprintln(w, r.Contract.Get.String())
}

// Need to validate the output and expected output
func (r Router) Post(w http.ResponseWriter) {
	// fmt.Fprintf(w, "Hello Post\n")
	fmt.Fprintln(w, r.Contract.Post.String())
}

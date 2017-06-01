package router

import (
	"contract"
	"fmt"
	"io/ioutil"
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
		r.Get(w, req)
	case "POST":
		r.Post(w, req)
	}
}

func (r Router) Get(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Hello Get\n")
	fmt.Fprintln(w, r.Contract.Get.String())
}

// Need to validate the output and expected output
func (r Router) Post(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Hello Post\n")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	//expected := r.Contract.Post.Encode()
	expected := r.Contract.Post
	// fmt.Fprintln(w, r.Contract.Post.String())
	fmt.Fprintln(w, string(body) == expected)
	fmt.Printf("%s\n%s\n", string(body), expected)
}

package handlers

import(
  "fmt"
  "net/http"
)

func (hnd *Handler) Blank(res http.ResponseWriter, req *http.Request){
  hnd.Tpl.ExecuteTemplate(res,"blank.html",nil)
  fmt.Println("Running blank")
  return
}

func (hnd *Handler) Test(res http.ResponseWriter, req *http.Request) {
  http.Error(res, "An error occurred", http.StatusInternalServerError)
	/*fmt.Println("HND WORKS")
	res.Write([]byte("Handler Test Function Works"))*/
}

/*
So I want to serve a bunch of templates by combining them using text templates:
  Header
    Notifications
  Sidebar
  Body
  Footer
*/

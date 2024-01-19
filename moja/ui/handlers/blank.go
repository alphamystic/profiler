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
  //http.Error(res, "An error occurred", http.StatusInternalServerError)
  data := map[string]interface{}{
    "header":"header.tmpl",
    "sidebar":"sidebar.tmpl",
    "body":"body.tmpl",
    "footer":"footer.tmpl",
  }
	hnd.Tpl.ExecuteTemplate(res,"base.tmpl",data)
}

/*
So I want to serve a bunch of templates by combining them using text templates:
  Header
    Notifications
  Sidebar
  Body
  Footer


given I have 3 files
file 1
This is File 1
{{.DataForFile1}}

file 2
this is file 2
{{.DataForFIle2}}

file 3
This is file 3
{{.DataForFile3}}

how can I add data rather combine them to the base Html File
<html>
<head>
</head>
<body>
file 1
file 2
file 3
</body>
</html>
*/

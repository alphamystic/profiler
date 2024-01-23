package handlers

import(
  "fmt"
  "net/http"
  "github.com/alphamystic/profiler/libgo/utils"
)


func (hnd *Handler) Profiler(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetDash("profiler","dash_profiler.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"profiler",nil)
}


func (hnd *Handler) Auto(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetDash("profiler","dash_auto.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"profiler",nil)
}

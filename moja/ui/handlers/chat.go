package handlers

import(
  "net/http"
)

func (hnd *Handler) Chat(res http.ResponseWriter, req *http.Request){
  //hnd.Tpl.ExecuteTemplate(res,"chat.html",nil)
  http.Error(res, "An error occurred", http.StatusInternalServerError)
  return
}

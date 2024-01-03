package router

import(
  "os"
  "log"
  "fmt"
  "time"
  "context"
  "syscall"
  "net/http"
  "os/signal"
  "github.com/alphamystic/profiler/libgo/utils"
  ent"github.com/alphamystic/profiler/libgo/entities"
  "github.com/alphamystic/profiler/moja/ui/handlers"
)

type Router struct {
  Mux *http.ServeMux
  HTTPSvr *http.Server
  HTTPSSvr *http.Server
}

// should probably receive a server
func NewRouter(httpsSvr,httpSvr *http.Server) *Router {
  return &Router {
    Mux: http.NewServeMux(),
    HTTPSvr: httpSvr,
    HTTPSSvr: httpsSvr,
  }
}


func (rtr *Router) Run(reg bool){
  // create a file server for the static files
  fs := http.FileServer(http.Dir("./static"))
  // Cache static files for 1 hour (adjust as needed)
  rtr.Mux.Handle("/static/", http.StripPrefix("/static", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Cache-Control", "max-age=3600")
    fs.ServeHTTP(res,req)
  })))

  // create a file server for the downloadable files
  downloads_dir := http.FileServer(http.Dir("./downloads"))
  rtr.Mux.Handle("/downloads/", http.StripPrefix("/downloads", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Cache-Control", "max-age=3600")
    downloads_dir.ServeHTTP(res,req)
  })))

  // create a file server for the uploaded files
  uploads := http.FileServer(http.Dir("./uploads"))
  rtr.Mux.Handle("/uploads/", http.StripPrefix("/uploads", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Cache-Control", "max-age=3600")
    uploads.ServeHTTP(res,req)
  })))

  // create your db connection
  dbConfig := ent.IntitializeConnector("root","","localhost","odin")
  dbConn,err := ent.NewMySQLConnector(dbConfig)
  if err != nil {
    utils.Warning(fmt.Sprintf("Error connecting to the DB. \n[-]   ERROR: %s",err))
    return
  }

  // create a request logger
  rl := utils.NewRequestLogger("./.data/logs/requests/",066)
  // create shutdown channels
  ShutdownCh := make(chan bool)
  DoneCh := make(chan bool)
  //create  your handler
  hnd,err := handlers.NewHandler(dbConn, ShutdownCh, DoneCh,rl)
  if err != nil {
    utils.Logerror(err)
    return
  }

  //rtr.Mux.HandleFunc("/",hnd.Home)
  rtr.Mux.HandleFunc("/pcapanalyzer",hnd.PcapAnalyzer)

  // Start the server on the background
   go func(){
     if err := rtr.HTTPSvr.ListenAndServe(); err != http.ErrServerClosed {
       log.Fatalf("[-] Error starting server: %s\n",err.Error())
     }
   }()
   go func(){
     // we need to find a better way of supplying this
     if err := rtr.HTTPSSvr.ListenAndServeTLS("../../../certs/server.crt", "../../../certs/server.key"); err != http.ErrServerClosed {
       log.Fatalf("[-] Error starting HTTPS server: %s\n",err.Error())
     }
   }()
   interruptChan := make(chan os.Signal,1)
   signal.Notify(interruptChan,os.Interrupt, syscall.SIGTERM)
   //sedn a close channel to the handler
   hnd.ShutdownChan <- true
   // wait for the receiver to finish writing all logs
   <-hnd.DoneChan
   // read from the interrupt chan and shutdown
   <-interruptChan
   shutdownCtx,shutdownCancel := context.WithTimeout(context.Background(),5 * time.Second)
   defer shutdownCancel()
   err = rtr.HTTPSvr.Shutdown(shutdownCtx)
   if err != nil {
     log.Fatalf("[-] Server shutdown error: %s\n",err.Error())
   }
   err = rtr.HTTPSSvr.Shutdown(shutdownCtx)
   if err != nil {
     log.Fatalf("[-] Server shutdown error: %s\n",err.Error())
   }
   log.Println("[+] Server gracefully stopped.")
}

package main

import (
    "fmt"
    "log"
   // "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/gocql/gocql"
)
type Server struct {
    config *Config
  }
  
  func New(config *Config) *Server {
    return &Server{
      config: config,
    }
  }

  type timerel struct{
      counterval int
      creationtime int
      steptime int
  }
func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}


_
func main() {

    cluster := gocql.NewCluster("PublicIP", "PublicIP", "PublicIP") //replace PublicIP with the IP addresses used by your cluster.
cluster.Consistency = gocql.Quorum
cluster.ProtoVersion = 4
cluster.ConnectTimeout = time.Second * 10
cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
session, err := cluster.CreateSession()
        if err != nil {
log.Println(err)
return
}
defer session.Close()


    r := mux.NewRouter()

    api := r.PathPrefix("/api/v1").Subrouter()
    //api.HandleFunc("", _create).Methods(http.MethodPost)
    api.HandleFunc("/CHECK", _check).Methods(http.MethodGet)
    api.HandleFunc("/RENDER", _render).Methods(http.MethodGet)
   
    //api.HandleFunc("/RENDER", put).Methods(http.MethodPut)
    api.HandleFunc("clear", _clear).Methods(http.MethodDelete)


	
    api.HandleFunc("/CREATE", _create).Methods(http.MethodPost)

    log.Fatal(http.ListenAndServe(":8080", r))
}
func _create(w http.ResponseWriter, r *http.Request)*rand.Rand {
   
    startval, ok := r.URL.Query()["startval"]
    steptime, ok := r.URL.Query()["steptime"]
   
    s1 := rand.NewSource(steptime)
    r1 := rand.New(s1)
    go count(){
        return r1 
    }

    return r1 
        
    
}

func _check(w http.ResponseWriter, r *http.Request)*rand.Rand {
   
    startval, ok := r.URL.Query()["startval"]
    steptime, ok := r.URL.Query()["steptime"]
   
    
    go current(){
        return startval+ steptime
    }()

    return startval+ steptime
        
    
}

func  __render() []*temerel {
    rows, _ := repository.database.Query(
      `SELECT counterval, creationtime, steptime FROM timerel;`
    )
    defer rows.Close()
  
    Timerel := []*timerel{}
  
    for rows.Next() {
      var (
        counterval   int
        creationtime string
        steptime  int
      )
  
      rows.Scan(&counterval, &creationtime, &steptime)
  
      timerel = append(timerel, &timerel{
        counterval:   counterval,
        creationtime: creationtime,
        steptime:  steptime,
      })
    }
  
    return timerel
  }
  
  

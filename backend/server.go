package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os/exec"
	"github.com/rs/cors"
	"bytes"
	"time"
)

func LeerCpu(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	go fmt.Println("Módulo CPU obtenido correctamente")
  output := string(out[:])

  fmt.Fprintf(w, output)
}


func LeerRam(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "cat /proc/mem_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	go fmt.Println("Módulo RAM obtenido correctamente")
  output := string(out[:])
  fmt.Fprintf(w, output)
}




func killProcess(w http.ResponseWriter, r *http.Request){
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "kill "+newStr
	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
//		log.Fatal(err)
go fmt.Println("error")
	}
    go fmt.Println(out)
	fmt.Fprintf(w, "OK")
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	go fmt.Println("Server Running on port: 8080")
	go router.HandleFunc("/leerram", LeerRam).Methods("GET")
	go router.HandleFunc("/leercpu", LeerCpu).Methods("GET")
	go router.HandleFunc("/killprocess", killProcess).Methods("POST")
	time.Sleep(time.Second)
    handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
	 
}

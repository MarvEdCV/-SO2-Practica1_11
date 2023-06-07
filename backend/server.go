package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os/exec"
	"strings"
	"strconv"
	"github.com/rs/cors"
	"bytes"
	"time"
)

func getCPU(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("sh", "-c", "ps -eo pcpu | sort -k 1 -r | head -50")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    go fmt.Println("CPU obtenido correctamente")

    output := string(out[:])
	//fmt.Fprintf(w, output)

    s := strings.Split(output, "\n")
    cpuUsado := 0.0
    for i := 1; i < 51; i++ {

    	valor, err := strconv.ParseFloat(strings.Trim(s[i], " "), 64)
    	if err != nil {
    		go fmt.Println("valorError ->" + s[i] + "<-" + strconv.Itoa(i))
    		go fmt.Println(err)
		}

		cpuUsado += valor 
		//go fmt.Println("valor ->" + s[i] + "<-" + strconv.Itoa(i))

	}


    fmt.Fprintf(w, "%f", cpuUsado)
}


func LeerCpu(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	go fmt.Println("Ram obtenida correctamente")
  output := string(out[:])

  fmt.Fprintf(w, output)
}


func LeerRam(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sh", "-c", "cat /proc/mem_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

  //go fmt.Println("Ram obtenida correctamente")
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
	go router.HandleFunc("/cpu", getCPU).Methods("GET")
	//go router.HandleFunc("/statistics", Statistics).Methods("GET")
	//go router.HandleFunc("/allprocess", getAllProcess).Methods("GET")
	go router.HandleFunc("/killprocess", killProcess).Methods("POST")
	//go router.HandleFunc("/treeprocess", getTreeProcess).Methods("GET")
	// cors.Default() setup the middleware with default options being
    // all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	time.Sleep(time.Second)
    handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
	 
}

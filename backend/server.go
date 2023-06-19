package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"github.com/gorilla/mux"   // Importar paquete para enrutamiento HTTP
	"github.com/rs/cors"       // Importar paquete para configurar CORS (Cross-Origin Resource Sharing)
	"bytes"
	"time"
)

// Función para leer el módulo de CPU
func LeerCpu(w http.ResponseWriter, r *http.Request){
	// Ejecutar el comando "cat /proc/cpu_grupo11" en el sistema operativo
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	go fmt.Println("Módulo CPU obtenido correctamente")
	output := string(out[:])

	fmt.Fprintf(w, output)
}

// Función para leer el módulo de RAM
func LeerRam(w http.ResponseWriter, r *http.Request){
	// Ejecutar el comando "cat /proc/mem_grupo11" en el sistema operativo
	cmd := exec.Command("sh", "-c", "cat /proc/mem_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	go fmt.Println("Módulo RAM obtenido correctamente")
	output := string(out[:])
	fmt.Fprintf(w, output)
}

// Función para matar un proceso
func killProcess(w http.ResponseWriter, r *http.Request){
	// Leer el cuerpo de la solicitud HTTP para obtener el ID del proceso a matar
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "kill "+newStr

	// Ejecutar el comando "kill <pid>" en el sistema operativo
	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		go fmt.Println("error")
	}
	go fmt.Println(out)
	fmt.Fprintf(w, "OK")
}

func main() {
	// Crear un enrutador utilizando el paquete gorilla/mux
	router := mux.NewRouter().StrictSlash(true)

	go fmt.Println("Server Running on port: 8080")
	
	// Definir las rutas y las funciones de controlador correspondientes
	go router.HandleFunc("/leerram", LeerRam).Methods("GET")         // Ruta para leer el módulo de RAM
	go router.HandleFunc("/leercpu", LeerCpu).Methods("GET")         // Ruta para leer el módulo de CPU
	go router.HandleFunc("/killprocess", killProcess).Methods("POST")  // Ruta para matar un proceso

	time.Sleep(time.Second)

	// Configurar CORS (Cross-Origin Resource Sharing) para permitir acceso a recursos desde diferentes dominios
	handler := cors.Default().Handler(router)

	// Iniciar el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

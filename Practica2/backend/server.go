package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"encoding/json"
	"strings"

	"github.com/gorilla/mux" // Importar paquete para enrutamiento HTTP
	"github.com/rs/cors"     // Importar paquete para configurar CORS (Cross-Origin Resource Sharing)
)

func Inicio(w http.ResponseWriter, r *http.Request) {
	// Ejecutar el comando "cat /proc/cpu_grupo11" en el sistema operativo
	go fmt.Println("Grupo 11 n_n")
	output := "Hola te saluda grupo 11 de Sopes2"

	fmt.Fprintf(w, output)
}

/*
** Función para leer el módulo de CPU
** @param w
** @param r
 */
func LeerCpu(w http.ResponseWriter, r *http.Request) {
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
func LeerRam(w http.ResponseWriter, r *http.Request) {
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
func killProcess(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud HTTP para obtener el ID del proceso a matar
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "kill " + newStr

	// Ejecutar el comando "kill <pid>" en el sistema operativo
	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		go fmt.Println("error")
	}
	go fmt.Println(out)
	fmt.Fprintf(w, "OK")
}

// Estructura para representar una entrada en el mapa de memoria
type MemoryMapEntry struct {
	AddressRange string `json:"addressRange"`
	Permissions  string `json:"permissions"`
	Offset       string `json:"offset"`
	Device       string `json:"device"`
	Inode        string `json:"inode"`
	Path         string `json:"path"`
	Size         uint64 // Nuevo campo para el tamaño de la región de memoria
}

// Estructura para el resultado de LeerMaps
type MapsResult struct {
	PID   string           `json:"pid"`
	Maps  []MemoryMapEntry `json:"maps"`
	Error string           `json:"error,omitempty"`
}

// Función para acceder a maps
func LeerMaps(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud HTTP para obtener el PID de una tarea
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "cat /proc/" + newStr + "/maps"

	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		response := MapsResult{
			PID:   newStr,
			Error: err.Error(),
		}
		sendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	maps := parseMemoryMapsOutput(string(out))
	response := MapsResult{
		PID:  newStr,
		Maps: maps,
	}
	sendJSONResponse(w, http.StatusOK, response)
}

// Función para calcular el tamaño de una región de memoria
func CalculateMemorySize(entry MemoryMapEntry) (uint64, error) {
	addressRange := entry.AddressRange
	rangeParts := strings.Split(addressRange, "-")

	startAddress, err := strconv.ParseUint(rangeParts[0], 16, 64)
	if err != nil {
		return 0, err
	}

	endAddress, err := strconv.ParseUint(rangeParts[1], 16, 64)
	if err != nil {
		return 0, err
	}

	size := (endAddress - startAddress) / 1024
	return size, nil
}

// Función para analizar la salida del archivo maps
func parseMemoryMapsOutput(output string) []MemoryMapEntry {
	lines := strings.Split(output, "\n")
	entries := make([]MemoryMapEntry, 0)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 5 {
			entry := MemoryMapEntry{
				AddressRange: fields[0],
				Permissions:  fields[1],
				Offset:       fields[2],
				Device:       fields[3],
				Inode:        fields[4],
			}

			if len(fields) > 5 {
				entry.Path = fields[5]
				size, err := CalculateMemorySize(entry)
				if err != nil {
					log.Println("Error calculating memory size:", err)
				}
				entry.Size = size
			}

			entries = append(entries, entry)
		}
	}

	return entries
}

// Función para enviar una respuesta HTTP con formato JSON
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Crear un enrutador utilizando el paquete gorilla/mux
	router := mux.NewRouter().StrictSlash(true)

	go fmt.Println("Server Running on port: 8080")

	// Definir las rutas y las funciones de controlador correspondientes
	go router.HandleFunc("/", Inicio).Methods("GET")
	go router.HandleFunc("/leerram", LeerRam).Methods("GET")          // Ruta para leer el módulo de RAM
	go router.HandleFunc("/leercpu", LeerCpu).Methods("GET")          // Ruta para leer el módulo de CPU
	go router.HandleFunc("/killprocess", killProcess).Methods("POST") // Ruta para matar un proceso
	go router.HandleFunc("/leermaps", LeerMaps).Methods("POST")       // Ruta ver maps

	time.Sleep(time.Second)

	// Configurar CORS (Cross-Origin Resource Sharing) para permitir acceso a recursos desde diferentes dominios
	handler := cors.Default().Handler(router)

	// Iniciar el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

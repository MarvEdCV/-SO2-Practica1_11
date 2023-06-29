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

/*
** Función para de terminar que la api esta en funcionamiento
** @param w respuesta del endpoint
** @param r peticion del endpoint
 */
func Inicio(w http.ResponseWriter, r *http.Request) {
	go fmt.Println("Grupo 11 n_n")
	output := "Hola te saluda grupo 11 de Sopes2"

	fmt.Fprintf(w, output)
}

/*
** Función para leer el módulo de CPU
** @param w respuesta del endpoint
** @param r peticion del endpoint
 */
func LeerCpu(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	go fmt.Println("Módulo CPU obtenido correctamente")
	output := string(out[:])

	fmt.Fprintf(w, output)
}

/*
** Función para leer el módulo de RAM
** @param w respuesta del endpoint
** @param r peticion del endpoint
 */
func LeerRam(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "-c", "cat /proc/mem_grupo11")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	go fmt.Println("Módulo RAM obtenido correctamente")
	output := string(out[:])
	fmt.Fprintf(w, output)
}

/*
** Función para matar un proceso
** @param w respuesta del endpoint
** @param r peticion del endpoint
 */
func killProcess(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "kill " + newStr

	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		go fmt.Println("error")
	}
	go fmt.Println(out)
	fmt.Fprintf(w, "OK")
}

/*
** Estructura para representar una entrada en el mapa de memoria

 */

type MemoryMapEntry struct {
	AddressRange string `json:"addressRange"`
	Permissions  string `json:"permissions"`
	Offset       string `json:"offset"`
	Device       string `json:"device"`
	Inode        string `json:"inode"`
	Path         string `json:"path"`
	Size         uint64 // Nuevo campo para el tamaño de la región de memoria
}

type MemoryMapEntryB struct {
	AddressRange string `json:"addressRange"`
	Permissions  string `json:"permissions"`
	Offset       string `json:"offset"`
	Device       string `json:"device"`
	Inode        string `json:"inode"`
	Path         string `json:"path"`
	Size         uint64 // Nuevo campo para el tamaño de la región de memoria
}

/*
** Estructura para el resultado de LeerMaps

 */

type MapsResultB struct {
	PID   string           `json:"pid"`
	Maps  []MemoryMapEntry `json:"maps"`
	Error string           `json:"error,omitempty"`
}

/*
** Función para acceder a maps
** @param w respuesta del endpoint
** @param r peticion del endpoint

 */

func LeerMaps(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "cat /proc/" + newStr + "/maps"

	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		response := MapsResultB{
			PID:   newStr,
			Error: err.Error(),
		}
		sendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	maps := parseMemoryMapsOutput(string(out))

	response := MapsResultB{
		PID:  newStr,
		Maps: maps,
	}
	sendJSONResponse(w, http.StatusOK, response)
}

/*
** Función para calcular el tamaño de una región de memoria
** @param entry recibe estructiura tipo MemoryMapEntry
** return valor tipo numerico y error

 */

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

/*
** Función para analizar la salida del archivo maps
** @param output recibe parametro tipo string
** return arreglo de structuras MemoryMapEntry

 */

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

/*
** Función para enviar una respuesta HTTP con formato JSON
** @param w respuesta del endpoint
** @param r peticion del endpoint

 */

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

/*
	type Block struct {
		Address        string `json:"address"`
		Permissions    string `json:"permissions"`
		Offset         string `json:"offset"`
		Dev            string `json:"dev"`
		Inode          string `json:"inode"`
		Pathname       string `json:"pathname"`
		Size           string `json:"Size"`
		KernelPageSize string `json:"KernelPageSize"`
		MMUPageSize    string `json:"MMUPageSize"`
		Rss            string `json:"Rss"`
		Pss            string `json:"Pss"`
		SharedClean    string `json:"Shared_Clean"`
		SharedDirty    string `json:"Shared_Dirty"`
		PrivateClean   string `json:"Private_Clean"`
		PrivateDirty   string `json:"Private_Dirty"`
		Referenced     string `json:"Referenced"`
		Anonymous      string `json:"Anonymous"`
		LazyFree       string `json:"LazyFree"`
		AnonHugePages  string `json:"AnonHugePages"`
		ShmemPmdMapped string `json:"ShmemPmdMapped"`
		FilePmdMapped  string `json:"FilePmdMapped"`
		SharedHugetlb  string `json:"Shared_Hugetlb"`
		PrivateHugetlb string `json:"Private_Hugetlb"`
		Swap           string `json:"Swap"`
		SwapPss        string `json:"SwapPss"`
		Locked         string `json:"Locked"`
		THPeligible    string `json:"THPeligible"`
		VmFlags        string `json:"VmFlags"`
	}

	func LeerSmaps(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newStr := buf.String()
		str := "cat /proc/" + newStr + "/smaps"

		cmd := exec.Command("sh", "-c", str)
		out, err := cmd.CombinedOutput()
		if err != nil {
			response := MapsResultB{
				PID:   newStr,
				Error: err.Error(),
			}
			sendJSONResponse(w, http.StatusInternalServerError, response)
			return
		}
		// Texto plano a convertir

		// Dividir el texto en bloques
		blocks := splitBlocks(string(out))
		//fmt.Println(string(out))
		// Crear una estructura para almacenar los bloques convertidos a JSON

		// Convertir los bloques a objetos JSON
		var jsonBlocks []Block
		for _, block := range blocks {
			jsonBlocks = append(jsonBlocks, convertBlockToJSON(block))
		}

		// Enviar la respuesta JSON
		sendJSONResponseB(w, http.StatusOK, jsonBlocks)
	}

	func splitBlocks(text string) []string {
		// Dividir el texto en bloques separados por líneas vacías
		return strings.Split(text, "VmFlags:")
	}

	func convertBlockToJSON(block string) Block {
		// Dividir el bloque en líneas y extraer los campos clave-valor
		lines := strings.Split(block, "\n")
		data := make(map[string]string)
		for _, line := range lines {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				data[key] = value
				if key == "VmFlags" {
					return Block{
						Address:        data["address"],
						Permissions:    data["permissions"],
						Offset:         data["offset"],
						Dev:            data["dev"],
						Inode:          data["inode"],
						Pathname:       data["pathname"],
						Size:           data["Size"],
						KernelPageSize: data["KernelPageSize"],
						MMUPageSize:    data["MMUPageSize"],
						Rss:            data["Rss"],
						Pss:            data["Pss"],
						SharedClean:    data["Shared_Clean"],
						SharedDirty:    data["Shared_Dirty"],
						PrivateClean:   data["Private_Clean"],
						PrivateDirty:   data["Private_Dirty"],
						Referenced:     data["Referenced"],
						Anonymous:      data["Anonymous"],
						LazyFree:       data["LazyFree"],
						AnonHugePages:  data["AnonHugePages"],
						ShmemPmdMapped: data["ShmemPmdMapped"],
						FilePmdMapped:  data["FilePmdMapped"],
						SharedHugetlb:  data["Shared_Hugetlb"],
						PrivateHugetlb: data["Private_Hugetlb"],
						Swap:           data["Swap"],
						SwapPss:        data["SwapPss"],
						Locked:         data["Locked"],
						THPeligible:    data["THPeligible"],
						VmFlags:        data["VmFlags"],
					}
				}
			}
		}

		// Crear un objeto Block a partir de los datos extraídos
		return Block{
			Address:        data["address"],
			Permissions:    data["permissions"],
			Offset:         data["offset"],
			Dev:            data["dev"],
			Inode:          data["inode"],
			Pathname:       data["pathname"],
			Size:           data["Size"],
			KernelPageSize: data["KernelPageSize"],
			MMUPageSize:    data["MMUPageSize"],
			Rss:            data["Rss"],
			Pss:            data["Pss"],
			SharedClean:    data["Shared_Clean"],
			SharedDirty:    data["Shared_Dirty"],
			PrivateClean:   data["Private_Clean"],
			PrivateDirty:   data["Private_Dirty"],
			Referenced:     data["Referenced"],
			Anonymous:      data["Anonymous"],
			LazyFree:       data["LazyFree"],
			AnonHugePages:  data["AnonHugePages"],
			ShmemPmdMapped: data["ShmemPmdMapped"],
			FilePmdMapped:  data["FilePmdMapped"],
			SharedHugetlb:  data["Shared_Hugetlb"],
			PrivateHugetlb: data["Private_Hugetlb"],
			Swap:           data["Swap"],
			SwapPss:        data["SwapPss"],
			Locked:         data["Locked"],
			THPeligible:    data["THPeligible"],
			VmFlags:        data["VmFlags"],
		}
	}
*/
func sendJSONResponseB(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	w.Write(jsonData)
}

type Block struct {
	Size string `json:"size"`
	Rss  string `json:"Rss"`
}

type Data struct {
	Blocks []Block `json:"data"`
}

func LeerSmaps(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "cat /proc/" + newStr + "/smaps"

	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	lines := strings.Split(string(out), "\n")
	blocks := []Block{}
	block := Block{}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 1 {
			continue
		} else {
			key := strings.TrimSuffix(fields[0], ":")
			value := strings.Join(fields[1:], " ")
			switch key {
			case "Size":
				block.Size = value
			case "Rss":
				block.Rss = value
			}
		}
		if strings.HasPrefix(line, "VmFlags") {
			blocks = append(blocks, block)
			block = Block{}
		}
	}

	data := Data{
		Blocks: blocks,
	}

	sendJSONResponseB(w, http.StatusOK, data)
}

type DataTotal struct {
	TotalSize float64 `json:"totalSize"`
	TotalRss  float64 `json:"totalRss"`
}

func LeerSmapsSizeRss(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	str := "cat /proc/" + newStr + "/smaps"

	cmd := exec.Command("sh", "-c", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	bloques := strings.Split(string(out), "\n")
	totalSize := 0
	totalRss := 0

	// Iterar sobre cada bloque y extraer los campos "Size" y "Rss"
	for _, bloque := range bloques {
		lines := strings.Split(bloque, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Size: ") {
				size := strings.TrimSpace(strings.TrimPrefix(line, "Size: "))
				// Convertir el tamaño a un número entero
				var sizeValue int
				fmt.Sscanf(size, "%d", &sizeValue)
				totalSize += sizeValue
			} else if strings.HasPrefix(line, "Rss: ") {
				rss := strings.TrimSpace(strings.TrimPrefix(line, "Rss: "))
				// Convertir el tamaño a un número entero
				var rssValue int
				fmt.Sscanf(rss, "%d", &rssValue)
				totalRss += rssValue
			}
		}
	}

	// Convertir los totales de KB a MB
	totalSizeMB := float64(totalSize) / 1024
	totalRssMB := float64(totalRss) / 1024

	// Crear la estructura de datos
	data := DataTotal{
		TotalSize: totalSizeMB,
		TotalRss:  totalRssMB,
	}

	sendJSONResponseB(w, http.StatusOK, data)
}

/*
** Función principal
 */
func main() {
	router := mux.NewRouter().StrictSlash(true)

	go fmt.Println("Server Running on port: 8080")

	go router.HandleFunc("/", Inicio).Methods("GET")
	go router.HandleFunc("/leerram", LeerRam).Methods("GET")          // Ruta para leer el módulo de RAM
	go router.HandleFunc("/leercpu", LeerCpu).Methods("GET")          // Ruta para leer el módulo de CPU
	go router.HandleFunc("/killprocess", killProcess).Methods("POST") // Ruta para matar un proceso
	go router.HandleFunc("/leermaps", LeerMaps).Methods("POST")       // Ruta ver maps
	go router.HandleFunc("/leersmaps", LeerSmaps).Methods("POST")
	go router.HandleFunc("/leersmapssizerss", LeerSmapsSizeRss).Methods("POST")

	time.Sleep(time.Second)

	// Configurar CORS (Cross-Origin Resource Sharing) para permitir acceso a recursos desde diferentes dominios
	handler := cors.Default().Handler(router)

	// Iniciar el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

#include <linux/module.h>      // Archivo de encabezado para la programación de módulos del kernel de Linux
#include <linux/init.h>        // Archivo de encabezado para macros de inicialización
#include <linux/seq_file.h>    // Archivo de encabezado para la escritura de archivos de secuencia
#include <linux/mm.h>          // Archivo de encabezado para operaciones de gestión de memoria
#include <linux/proc_fs.h>     // Archivo de encabezado para el sistema de archivos /proc
#include <linux/kernel.h>      // Archivo de encabezado para macros y funciones del kernel
#include <linux/sysinfo.h>     // Archivo de encabezado para obtener información del sistema

MODULE_LICENSE("GPL");                            // Licencia del módulo
MODULE_DESCRIPTION("Monitor modulo RAM");         // Descripción del módulo
MODULE_AUTHOR("Marvin Eduardo Catalán Véliz, Sara Paulina Medrano Cojulún, Wilson Eduardo Perez Echeverria"); // Autores del módulo

struct sysinfo inf;   // Estructura para almacenar información del sistema

// Función para escribir en el archivo de secuencia
static int write_file(struct seq_file *file, void *v){
    long total_mem, free_mem;
    si_meminfo(&inf);   // Obtener información de memoria del sistema y almacenarla en la estructura inf
    total_mem = (inf.totalram * 4 / 1024);   // Calcular la cantidad total de memoria en mb
    free_mem = (inf.freeram * 4 / 1024);     // Calcular la cantidad de memoria libre en mb
    seq_printf(file, "{\n");                                      // Escribir una cadena en el archivo de secuencia
    seq_printf(file, " \"MemoriaTotal\":%8lu,\n", total_mem);       // Escribir la cantidad total de memoria
    seq_printf(file, " \"MemoriaLibre\":%8lu,\n", free_mem);        // Escribir la cantidad de memoria libre
    seq_printf(file, " \"MemoriaUsada\":%i\n", 100 - (free_mem * 100) / total_mem);   // Escribir el porcentaje de memoria utilizada
    seq_printf(file, "}\n");                                      // Escribir una cadena en el archivo de secuencia
    return 0;
}

// Función para abrir el archivo
static int to_open(struct inode *inode, struct file *file){
    return single_open(file, write_file, NULL);   // Abrir el archivo de secuencia y llamar a la función write_file
}

// Si el kernel es 5.6 o superior, se usa la estructura proc_ops
static struct proc_ops operations =
{
    .proc_open = to_open,   // Puntero a la función que se ejecuta al abrir el archivo /proc
    .proc_read = seq_read   // Puntero a la función de lectura de secuencia
};

/* Si el kernel es anterior a 5.6, se utilizan file_operations
static struct file_operations operaciones =
{
    .open = al_abrir,   // Puntero a la función que se ejecuta al abrir el archivo /proc
    .read = seq_read    // Puntero a la función de lectura de secuencia
};
*/

// Función de montaje del módulo
static int mount_module(void){
    proc_create("mem_grupo11", 0, NULL, &operations);   // Crear una entrada en /proc para el archivo "mem_grupo11"
    printk(KERN_INFO "Hola mundo, somos el grupo 11 y este es el monitor de memoria\n");   // Imprimir un mensaje en el registro del kernel
    return 0;
}

// Función de desmontaje del módulo
static void disassemble_module(void){
    remove_proc_entry("mem_grupo11", NULL);   // Eliminar la entrada en /proc para el archivo "mem_grupo11"
    printk(KERN_INFO "Sayonara mundo, somos el grupo 11 y este fue el monitor de memoria\n");   // Imprimir un mensaje en el registro del kernel
}

// Especificar la función de inicialización del módulo
module_init(mount_module);

// Especificar la función de desmontaje del módulo
module_exit(disassemble_module);

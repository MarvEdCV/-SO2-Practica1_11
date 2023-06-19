#include <linux/proc_fs.h>          // Archivo de encabezado para el sistema de archivos /proc
#include <linux/seq_file.h>         // Archivo de encabezado para la escritura de archivos de secuencia
#include <asm/uaccess.h>            // Archivo de encabezado para operaciones de acceso a memoria del usuario
#include <linux/hugetlb.h>          // Archivo de encabezado para funciones relacionadas con páginas grandes
#include <linux/sched/signal.h>     // Archivo de encabezado para operaciones de señales
#include <linux/sched.h>            // Archivo de encabezado para operaciones de programación de tareas
#include <linux/module.h>           // Archivo de encabezado para la programación de módulos del kernel de Linux
#include <linux/init.h>             // Archivo de encabezado para macros de inicialización
#include <linux/kernel.h>           // Archivo de encabezado para macros y funciones del kernel
#include <linux/fs.h>               // Archivo de encabezado para operaciones del sistema de archivos

struct list_head *p;
struct task_struct *processes, ts, *tsk;

MODULE_LICENSE("GPL");                          // Licencia del módulo
MODULE_DESCRIPTION("Monitor modulo CPU");        // Descripción del módulo
MODULE_AUTHOR("Marvin Eduardo Catalán Véliz, Sara Paulina Medrano Cojulún, Wilson Eduardo Perez Echeverria");  // Autores del módulo

/*
** Función para escribir en el archivo de secuencia
** @param seq_file tipo file para tener control sonbre archivos
** @param v tipo void
 */
static int write_file(struct seq_file *file, void *v){

    seq_printf(file, "{\n");
    seq_printf(file, "\"procesos\":[\n");

    bool seconditerative=false;

    for_each_process(processes){

        if(seconditerative){
        seq_printf(file, ",");    
        }else{
            seconditerative=true;
        }

        seq_printf(file, "{\n");
        seq_printf(file, "\"PID\": %d, \n", processes->pid);               // Escribir el ID del proceso
        seq_printf(file, "\"Nombre\": \"%s\", \n", processes->comm);       // Escribir el nombre del proceso
        seq_printf(file, "\"Usuario\": %d, \n", (int) processes->sessionid);   // Escribir el ID de usuario del proceso
        seq_printf(file, "\"Memory\": \"%d\", \n",__kuid_val(processes->real_cred->uid));   // Escribir memoria del proceso
        seq_printf(file, "\"Estado\": %ld} \n", processes->state);       // Escribir el estado del proceso
    }

    seconditerative=false;
    seq_printf(file, "],\n");
    seq_printf(file, "\"arbol\":[\n");
    
    for_each_process(processes){
        if(seconditerative){
        seq_printf(file, ",");    
        }else{
            seconditerative=true;
        }

        seq_printf(file, "{\n");
        seq_printf(file, "\"id\": %d, \n", processes->pid);           // Escribir el ID del proceso
        seq_printf(file, "\"parentId\": %d ,\n",0);                  // Escribir el ID del proceso padre
        seq_printf(file, "\"label\": \"%s\", \n", processes->comm);  // Escribir el nombre del proceso
        seq_printf(file, "\"items\": %s \n", "[]");                  // Escribir los elementos del proceso (vacío en este caso)
        seq_printf(file, "}\n");

        // Iterar sobre los hijos de cada proceso
        list_for_each(p, &(processes->children)){
            seq_printf(file, ",{\n");
            /*Obtener el proceso (task_struct) desde el nodo de la lista enlazada
            Desenlazar el nodo de la lista y obtener un puntero al objeto task_struct correspondiente.
            La macro list_entry toma un puntero al nodo de la lista (p), el tipo de estructura (struct task_struct),
            y el nombre del miembro que se utiliza para enlazar los nodos en la lista (sibling).
            Devuelve un puntero al objeto task_struct que contiene el nodo de la lista actual (p).
            */

            ts = *list_entry(p, struct task_struct, sibling);
            seq_printf(file, "     \"id\":%d, \n", ts.pid);              // Escribir el ID del hijo
            seq_printf(file, "     \"label\":\"%s\",\n", ts.comm);       // Escribir el nombre del hijo
            seq_printf(file, "     \"parentId\":%d,\n", processes->pid); // Escribir el ID del proceso padre
            seq_printf(file, "\"items\": %s \n", "[]");                  // Escribir los elementos del hijo (vacío en este caso)
            seq_printf(file, "}\n");
        }             
    }
    
    seq_printf(file, "]\n");
    seq_printf(file, "}\n");

    return 0;
}

/*
** Función para abrir el archivo
** @param inode structura tipo inode
** @param file structura tipo file para tener control sonbre archivos
 */
static int to_open(struct inode *inode, struct file *file){
    return single_open(file, write_file, NULL); // Abrir el archivo de secuencia y llamar a la función write_file
}

//Verificar el Kernel con el comando uname -r

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
/*static struct proc_ops operations ={
    .proc_open = to_open, // Puntero a la función que se ejecuta al abrir el archivo /proc
    .proc_read = seq_read // Puntero a la función de lectura de secuencia
};*/

/*
** Si el kernel es menor al 5.6 usan file_operations
 */
static struct file_operations operations =
{
    .open = to_open,
    .read = seq_read
};

/*
** Función de montaje del módulo
 */
static int mount_module(void){
    proc_create("cpu_grupo11", 0, NULL, &operations);  // Crear una entrada en /proc para el archivo "mem_grupo11"
    printk(KERN_INFO "Hola mundo, somos el grupo 11 y este es el monitor de cpu\n"); // Imprimir un mensaje en el registro del kernel
    return 0;
}

/*
** Función de desmontaje del módulo
 */
static void disassemble_module(void){
    remove_proc_entry("cpu_grupo11", NULL); // Eliminar la entrada en /proc para el archivo "mem_grupo11"
    printk(KERN_INFO "Sayonara mundo, somos el grupo 11 y este fue el monitor de cpu\n"); // Imprimir un mensaje en el registro del kernel
}

/*
** Especificar la función de inicialización del módulo
 */
module_init(mount_module);
/*
** Especificar la función de desmontaje del módulo
 */
module_exit(disassemble_module);

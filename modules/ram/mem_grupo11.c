#include <linux/module.h>
#include <linux/init.h>
#include <linux/seq_file.h>
#include <linux/sysinfo.h>


MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Monitor modulo RAM");
MODULE_AUTHOR("Marvin Eduardo Catalán Véliz, Sara Paulina Medrano Cojulún");

struct sysinfo inf;

static int write_file(struct seq_file *file, void *v){
    si_meminfo(&inf);
    long total_mem = (inf.totalram * 4 / 1024);
    long free_mem = (inf.freeram * 4 / 1024);
    seq_file(file, "{\n");
    seq_printf(file, " \"MemoriaTotal\":%8lu,\n", total_mem);
    seq_printf(file, " \"MemoriaLibre\":%8lu,\n", free_mem);
    seq_printf(file, " \"MemoriaUsada\":%i\n", (free_mem / total_mem) * 100);
    seq_printf(archivo, "}\n");
    return 0;
}

static int to_open(struct *inode, struct file *file){
    return single_open(file, write_file, NULL);
}

//Verificar el Kernel con el comando uname -r

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operations =
{
    .proc_open = to_open,
    .proc_read = seq_read
};

/*Si el kernel es menor al 5.6 usan file_operations
static struct file_operations operaciones =
{
    .open = al_abrir,
    .read = seq_read
};
*/

static int mount_module(void){
    proc_create("mem_grupo11", 0, NULL, &operations);
    printk(KERN_INFO "Hola mundo, somos el grupo 11 y este es el monitor de memoria\n");
    return 0;
}

static void disassemble_module(void){
    remove_proc_entry("mem_grupo11", NULL);
    printk(KERN_INFO "Sayonara mundo, somos el grupo 11 y este fue el monitor de memoria\n");
}

module_init(mount_module);
module_exit(disassemble_module);
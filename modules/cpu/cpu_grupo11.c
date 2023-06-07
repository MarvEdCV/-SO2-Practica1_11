#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/hugetlb.h>
#include <linux/sched/signal.h>
#include <linux/sched.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/fs.h>

struct list_head *p;
struct task_struct *processes, ts, *tsk;

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Monitor modulo CPU");
MODULE_AUTHOR("Marvin Eduardo Catalán Véliz, Sara Paulina Medrano Cojulún");


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
        seq_printf(file, "\"PID\": %d, \n", processes->pid);
        seq_printf(file, "\"Nombre\": \"%s\", \n", processes->comm);
        seq_printf(file, "\"Usuario\": %d, \n", (int) processes->sessionid);
        seq_printf(file, "\"Memory\": \"%d\", \n",__kuid_val(processes->real_cred->uid));
        seq_printf(file, "\"Estado\": %ld} \n", processes->__state);

        list_for_each(p, &(processes->children)){
            seq_printf(file, ",{\n");
            ts = *list_entry(p, struct task_struct, sibling);
            seq_printf(file, "     \"processes padre\":%d,\n", processes->pid);
            seq_printf(file, "     \"PID\":%d, \n", ts.pid);
            seq_printf(file, "     \"Nombre\":\"%s\",\n", ts.comm);
            seq_printf(file, "     \"Usuario\": %d, \n", (int) processes->sessionid);
            seq_printf(file, "     \"Estado\":%ld \n", ts.__state);
            seq_printf(file, "}\n");
        }
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
        seq_printf(file, "\"id\": %d, \n", processes->pid);
        seq_printf(file, "\"parentId\": %d ,\n",0);
        seq_printf(file, "\"label\": \"%s\", \n", processes->comm);
        seq_printf(file, "\"items\": %s \n", "[]");
        seq_printf(file, "}\n");

        list_for_each(p, &(processes->children)){
            seq_printf(file, ",{\n");
            ts = *list_entry(p, struct task_struct, sibling);
            seq_printf(file, "     \"id\":%d, \n", ts.pid);
            seq_printf(file, "     \"label\":\"%s\",\n", ts.comm);
            seq_printf(file, "     \"parentId\":%d,\n", processes->pid);
            seq_printf(file, "\"items\": %s \n", "[]");
            seq_printf(file, "}\n");
        }             
    }
    
    seq_printf(file, "]\n");
    seq_printf(file, "}\n");

    return 0;
}

static int to_open(struct inode *inode, struct file *file){
    return single_open(file, write_file, NULL);
}

//Verificar el Kernel con el comando uname -r

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operations ={
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
    proc_create("cpu_grupo11", 0, NULL, &operations);
    printk(KERN_INFO "Hola mundo, somos el grupo 11 y este es el monitor de cpu\n");
    return 0;
}

static void disassemble_module(void){
    remove_proc_entry("cpu_grupo11", NULL);
    printk(KERN_INFO "Sayonara mundo, somos el grupo 11 y este fue el monitor de cpu\n");
}

module_init(mount_module);
module_exit(disassemble_module);
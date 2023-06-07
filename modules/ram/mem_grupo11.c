#include <linux/module.h>
#include <linux/init.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Monitor modulo RAM");
MODULE_AUTHOR("Marvin Eduardo Catalán Véliz, Sara Paulina Medrano Cojulún");

static int mount_module(void)
{
    //proc_create("memo_practica1", 0, NULL, &operaciones);
    printk(KERN_INFO "Hola mundo, somos el grupo 11 y este es el monitor de memoria\n");
    return 0;
}

static void disassemble_module(void)
{
    //remove_proc_entry("memo_practica1", NULL);
    printk(KERN_INFO "Sayonara mundo, somos el grupo 11 y este fue el monitor de memoria\n");
}

module_init(mount_module);
module_exit(disassemble_module);
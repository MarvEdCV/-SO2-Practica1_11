#include <linux/module.h>
#include <linux/init.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Monitor modulo CPU");
MODULE_AUTHOR("Marvin Eduardo Catalán Véliz, Sara Paulina Medrano Cojulún");

static int mount_module(void)
{
    printk(KERN_INFO "Hola mundo, somos el grupo 11 y este es el monitor de cpu\n");
    return 0;
}

static void disassemble_module(void)
{
    printk(KERN_INFO "Sayonara mundo, somos el grupo 11 y este fue el monitor de cpu\n");
}

module_init(mount_module);
module_exit(disassemble_module);
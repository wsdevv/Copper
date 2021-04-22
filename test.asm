section .text
log:
    mov ebx, 1 ;file descriptor (this is the terminal)
    mov eax, 4 ;sys_write
    int 0x80  ;calls the kernel

    ret ;

test6:
   push ebp
   mov esp, ebp
   
   mov ebp, esp
   pop ebp
   ret ;
.Const.0:
    DB "[PROGRAM]: hello, world!", 0xA, ""
.Const.0.length equ $-.Const.0
global _start
_start:
   
   xor ebx,ebx
   mov eax, 1
   int 0x80

section .text
;this file contains all the base functions used by the
;compiler by default

; syntax for logl
;
; sets the params for log
; mov ecx, msg
; mov edx, msg_len 
; call log
;
; gets rid of the values from here 
; xor ecx, ecx
; xor edx, edx
log:
    mov ebx, 1 ;file descriptor (this is the terminal)
    mov eax, 4 ;sys_write
    int 0x80  ;calls the kernel

    ret ;

ftestr:
   mov DWORD [esp+4], .Const.0;
   mov DWORD [esp+8], .Const.0.length;
   mov DWORD [esp+12], .Const.1;
   mov DWORD [esp+16], .Const.1.length;
   mov ecx,[esp+4];
   mov edx,[esp+8];
   call log       ;
   xor ecx, ecx   ;
   xor edx, edx   ;
   mov DWORD [esp+4],0;
   mov DWORD [esp+12],0;
   mov DWORD [esp+16],0;
   mov DWORD [esp+16],0;
   ret ;
.Const.0:
    DB "[PROGRAM]: hello, world!", 0xA, ""
.Const.0.length equ $-.Const.0
.Const.1:
    DB "This is a useless variable declaration! (WARNING, THERE ARE SOME FATAL BUGS WITH VARS)"
.Const.1.length equ $-.Const.1
global _start
_start:
   call ftestr
   xor ebx,ebx
   mov eax, 1
   int 0x80

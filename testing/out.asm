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

test66:
   push ebp
   mov  ebp, esp
   sub esp, 8
   mov DWORD [ebp-4], .Const.0;
   mov DWORD [ebp-8], .Const.0.length;
   mov ecx,[ebp-4];
   mov edx,[ebp-8];
   call log       ;
   xor ecx, ecx   ;
   xor edx, edx   ;
   mov DWORD [ebp-4],0;
   mov DWORD [ebp-8],0;
   mov esp, ebp
   pop ebp
   ret ;
.Const.0:
    DB "[PROGRAM]: hello, world!", 0xA, ""
.Const.0.length equ $-.Const.0
global _start
_start:
   push ebp
   mov  ebp, esp
   sub esp, 8
   mov DWORD [ebp-4], .Const.0;
   mov DWORD [ebp-8], .Const.0.length;
   call test66
   mov DWORD [ebp-4],0;
   mov DWORD [ebp-8],0;
   mov esp, ebp
   pop ebp
   xor ebx,ebx
   mov eax, 1
   int 0x80

.Const.0:
    DB "[PROGRAM]: test variable 2", 0xA, ""
.Const.0.length equ $-.Const.0
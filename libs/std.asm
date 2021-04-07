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

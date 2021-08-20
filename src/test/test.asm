section .text:
 global _main 

_main:

   push ebp
   mov esp, ebp
   sub esp, 8
   mov DWORD [ebp-4], .eIBx_ZABd_cZZaabag
   mov DWORD [ebp-8], ._khh_p__YYinSYeS
   mov ebp, esp
   pop ebp
   ret ;
.eIBx_ZABd_cZZaabag:
  DB "hello""
.eIBx_ZABd_cZZaabag.length equ $-eIBx_ZABd_cZZaabag
._khh_p__YYinSYeS:
  DB "hello""
._khh_p__YYinSYeS.length equ $-_khh_p__YYinSYeS
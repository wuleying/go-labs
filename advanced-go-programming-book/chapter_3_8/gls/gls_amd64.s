#include "go_asm.h"
#include "textflag.h"

// func Get() int64
TEXT ·Get(SB),NOSPLIT,$0-8
	MOVQ (TLS), R14
	MOVQ g_goid(R14), R13
	MOVQ R13, ret+0(FP)
	RET

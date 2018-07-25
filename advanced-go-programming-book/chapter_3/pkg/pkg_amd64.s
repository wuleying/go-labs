TEXT    ·Add+0(SB),$0-24
    MOVQ    a+0(FP),BX
    MOVQ    b+8(FP),BP
    ADDQ    BP,BX
    MOVQ    BX,res+16(FP)
    RET     ;

DATA Id<>+0x00(SB)/8, $0x3725000000000000
GLOBL Id<>(SB), 8, $8

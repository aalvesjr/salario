package salary

import (
	"github.com/aalvesjr/salario/impostos"
)

type Salary struct {
	Gross             float32
	Net               float32
	Discounts         float32
	INSSRate          float32
	INSSBase          float32
	INSS              float32
	IRBase            float32
	IRRate            float32
	IRDiscount        float32
	IRWithoutDiscount float32
	IR                float32
}

func NewSalary(value, discounts float32) Salary {
	s := Salary{Gross: value, Discounts: discounts}
	inss := impostos.NewINSS(value)
	ir := impostos.NewIR(value)

	s.INSSRate, s.INSSBase, s.INSS = inss.Aliquota, inss.Base, inss.Valor
	s.IRRate, s.IRBase, s.IRWithoutDiscount, s.IRDiscount, s.IR = ir.Aliquota, ir.Base, ir.ValorSemDesconto, ir.Desconto, ir.Valor

	s.Net = s.Gross - (s.INSS + s.IR + s.Discounts)
	return s
}

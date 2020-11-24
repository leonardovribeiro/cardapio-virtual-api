package models

// Currency representa uma quantidade de dinheiro em centavos
type Currency int64

// ToCurrency converte um float64 em dinheiro (BRL)
func ToCurrency(f float64) Currency {
	return Currency((f * 100) + 0.5)
}

// Float64 converte dinheiro em float64
func (m Currency) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply multiplica o dinheiro por um float64 de forma segura
// arredondando para o mais próximo de centavos
func (m Currency) Multiply(f float64) Currency {
	x := (float64(m) * f) + 0.5
	return Currency(x)
}

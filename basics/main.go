package main

import "fmt"

func main() {
	// Numeros
	entero := 10
	decimal := 3.1416
	suma := entero + int(decimal)
	fmt.Println(suma)

	// Texto
	mensaje := "Hola, "
	concat := mensaje + "Carlos"
	


}

// Data types
// bool => Flag o condicionaes true/false == false
// string => Cadena de caracteres | para representar texto == ""
// int, int8, int16 int64, int32 => Entero | Controlar el tamaño de los entero == 0
// float32, float64 => Representar valores numericos reales, con punto, depende del sistema == 0
// uin, uin16, uin32 | Entero sin signo, valor absoluto == 0
// byte => uint8 | Trabajar con datos binarios == 0
// rune => Cuando necesitas representar un solo caractar que ocupa mas de un byte == 0
// complex64, complex128 => Cuando tiene una parte real y una imaginaria == 0 + 0i

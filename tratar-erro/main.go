package main

import "fmt"

func divide(numerador, denominador int) (float64, error) {
	if denominador == 0 {
		return 0, fmt.Errorf("Falha ao dividir o valor")
		// panic(nil)
	}
	return float64(numerador) / float64(denominador), nil
}

func main() {
	numerador := 10
	denominador := 0

	resultado, err := divide(numerador, denominador)
	if err != nil {
		fmt.Printf("Erro: %s\n", err)
		return
	}
	fmt.Printf("%d / %d = %2.f\n", numerador, denominador, resultado)

}

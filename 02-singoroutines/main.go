package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Usamos el tiempo actual en nanosegundos para asegurar una semilla diferente en cada ejecución
	rand.Seed(time.Now().UnixNano())

	begin := time.Now() // Guardamos el tiempo de inicio para calcular la duración total
	fmt.Println("==== Construction begins ====")

	// Ejecutamos las tareas de los trabajadores de forma secuencial
	Worker("Bob", "installing a door")
	Worker("Alice", "moving the lawn")
	Worker("John", "painting the walls")

	// Imprimimos el tiempo total que tomó el proceso
	fmt.Printf("==== Construction ended in %.2f seconds ====\n", time.Since(begin).Seconds())
}

// Worker inicia un trabajo y toma un número aleatorio de segundos entre 3 y 5 para terminarlo
func Worker(name, job string) {
	fmt.Printf("Worker %s began to work on %s\n", name, job)

	// Simulamos el trabajo del trabajador, que toma entre 3 y 5 segundos
	for i := 1; i <= 10; i++ {
		fmt.Printf("Worker %s is working %d\n", name, i)
		time.Sleep(50 * time.Millisecond) // Simulamos tiempo de trabajo con Sleep
	}

	fmt.Printf("Worker %s finished working on %s\n", name, job)
}

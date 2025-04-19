package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Crear una nueva fuente aleatoria basada en el valor actual de UnixNano
	source := rand.NewSource(time.Now().UnixNano())
	// Crear un nuevo generador de números aleatorios usando esa fuente
	r := rand.New(source)

	begin := time.Now()
	fmt.Println("==== Construction begins ====")

	// Inicializar el WaitGroup
	wg := &sync.WaitGroup{}

	// Iniciar cada trabajador como una goroutine
	go Worker("Bob", "installing a door", r, wg)
	wg.Add(1)
	go Worker("Alice", "mowing the lawn", r, wg)
	wg.Add(1)
	go Worker("John", "painting the walls", r, wg)
	wg.Add(1)

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	// Imprimir el tiempo transcurrido
	fmt.Printf("==== Construction ended in %.2f seconds ====\n", time.Since(begin).Seconds())
}

// Worker inicia un trabajo y toma un número aleatorio de segundos entre 3 y 5 para finalizarlo
func Worker(name, job string, r *rand.Rand, wg *sync.WaitGroup) {
	fmt.Printf("Worker %s began to work on %s\n", name, job)

	for i := 1; i <= 10; i++ {
		fmt.Printf("Worker %s is working %d\n", name, i)
		time.Sleep(50 * time.Millisecond)
	}

	// Imprimir cuando el trabajador termina
	fmt.Printf("Worker %s finished working on %s\n", name, job)

	// Indicar que la goroutine ha terminado
	wg.Done()
}

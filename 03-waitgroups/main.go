package main

import (
	"fmt"       // Paquete para formatear la salida y mostrar en consola
	"math/rand" // Paquete para generar números aleatorios
	"sync"      // Paquete para sincronización, incluye WaitGroup
	"time"      // Paquete para manejar el tiempo y las fechas
)

func main() {
	// Crear una nueva fuente aleatoria basada en el valor actual de UnixNano
	source := rand.NewSource(time.Now().UnixNano()) // Usamos el tiempo actual (en nanosegundos) como semilla
	// Crear un nuevo generador de números aleatorios usando esa fuente
	r := rand.New(source) // Esto nos da un generador 'r' que podemos usar para generar números aleatorios

	// Guardamos el tiempo de inicio para calcular el tiempo total de ejecución
	begin := time.Now()
	fmt.Println("==== Construction begins ====")

	// Inicializamos el WaitGroup para sincronizar las goroutines
	wg := &sync.WaitGroup{}

	// Iniciar cada trabajador como una goroutine (función ejecutada concurrentemente)
	// Cada trabajador tiene un trabajo diferente y utiliza el generador de números aleatorios
	go Worker("Bob", "installing a door", r, wg) // La goroutine comienza el trabajo de "Bob"
	wg.Add(1)                                    // Añadimos un trabajo al WaitGroup

	go Worker("Alice", "mowing the lawn", r, wg) // La goroutine comienza el trabajo de "Alice"
	wg.Add(1)                                    // Añadimos un trabajo al WaitGroup

	go Worker("John", "painting the walls", r, wg) // La goroutine comienza el trabajo de "John"
	wg.Add(1)                                      // Añadimos un trabajo al WaitGroup

	// Esperamos a que todas las goroutines terminen antes de continuar
	wg.Wait() // Este comando bloquea la ejecución del main hasta que todas las goroutines terminen

	// Imprimir el tiempo total que tomó la "construcción"
	fmt.Printf("==== Construction ended in %.2f seconds ====\n", time.Since(begin).Seconds())
}

// Worker inicia un trabajo y toma un número aleatorio de segundos entre 3 y 5 para terminarlo
// Recibe el nombre, trabajo, generador de números aleatorios, y el WaitGroup para indicar cuándo ha terminado
func Worker(name, job string, r *rand.Rand, wg *sync.WaitGroup) {
	// Imprimir que el trabajador comenzó a trabajar
	fmt.Printf("Worker %s began to work on %s\n", name, job)

	// Generar un tiempo aleatorio de trabajo entre 3 y 5 segundos
	randomTime := r.Intn(3) + 3        // r.Intn(3) genera un número entre 0 y 2, y sumarle 3 da un rango entre 3 y 5
	for i := 1; i <= randomTime; i++ { // Simulamos el trabajo con un ciclo que dura entre 3 y 5 iteraciones
		fmt.Printf("Worker %s is working %d\n", name, i)
		time.Sleep(1 * time.Second) // Simula 1 segundo de trabajo por cada iteración
	}

	// Imprimir cuando el trabajador termina su tarea
	fmt.Printf("Worker %s finished working on %s\n", name, job)

	// Indicar que la goroutine ha terminado su tarea
	wg.Done() // Disminuye el contador de tareas en el WaitGroup
}

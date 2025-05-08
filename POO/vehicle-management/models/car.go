package models

// Estructura Car que "compone" a Vehicle
type Car struct {
	Vehicle
	Doors int
}

// Método específico para Car
func (c Car) Start() string {
	return c.Vehicle.Start() + " El coche tiene " + string(rune(c.Doors)) + " puertas."
}

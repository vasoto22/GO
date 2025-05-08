package models

// Estructura Truck que "compone" a Vehicle
type Truck struct {
	Vehicle
	PayloadCapacity int
}

// Método específico para Truck
func (t Truck) Start() string {
	return t.Vehicle.Start() + " El camión tiene una capacidad de carga de " + string(rune(t.PayloadCapacity)) + " kg."
}

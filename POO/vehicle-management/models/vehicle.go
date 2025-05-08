package models

// Estructura base Vehicle
type Vehicle struct {
	ID    string
	Make  string
	Model string
	Year  int
}

// Método común para todos los vehículos
func (v Vehicle) Start() string {
	return v.Make + " " + v.Model + " está arrancando."
}

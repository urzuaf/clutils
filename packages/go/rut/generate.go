package rut

import (
	"math/rand/v2"
	"strconv"
)

// Generate genera un RUT aleatorio y matemáticamente válido con formato estándar.
// Por defecto utiliza el rango seguro habitual de personas naturales.
//
// Ejemplos:
//
//	rut.Generate() // ej: "18.123.458-K"
//	rut.Generate(FormatOptions{Dots: false, Hyphen: true}) // ej: "18123458-K"
func Generate(opts ...FormatOptions) string {
	return GenerateBetween(DefaultGeneralMin, DefaultGeneralMax, opts...)
}

// GenerateBetween genera un RUT matemáticamente válido dentro de un rango numérico específico.
// Corrige de forma automática valores negativos o límites invertidos para garantizar seguridad total.
//
// Ejemplos:
//
//	rut.GenerateBetween(1_000_000, 10_000_000) // RUT antiguo/bajo
//	rut.GenerateBetween(50_000_000, 99_000_000) // RUT en rango de personas jurídicas
func GenerateBetween(min, max int, opts ...FormatOptions) string {
	if min < DefaultGeneralMin {
		min = DefaultGeneralMin
	}
	if max < DefaultGeneralMin {
		max = DefaultGeneralMax
	}

	if min > max {
		min, max = max, min
	}

	body := rand.IntN(max-min+1) + min
	dv := CalculateDv(body)

	rut := strconv.Itoa(body) + dv

	return Format(rut, opts...)
}

// GeneratePerson genera un RUT aleatorio representativo de personas naturales vivas en Chile (rango 5M a 28M).
//
// Ejemplo:
//
//	rut.GeneratePerson() // ej: "21.458.123-5"
func GeneratePerson(opts ...FormatOptions) string {
	return GenerateBetween(DefaultPersonMin, DefaultPersonMax, opts...)
}

// GenerateBusiness genera un RUT aleatorio representativo de personas jurídicas y empresas en Chile (rango 70M a 80M).
//
// Ejemplo:
//
//	rut.GenerateBusiness() // ej: "76.458.123-0"
func GenerateBusiness(opts ...FormatOptions) string {
	return GenerateBetween(DefaultBusinessMin, DefaultBusinessMax, opts...)
}

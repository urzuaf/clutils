package rut

// Rangos numéricos canónicos sincronizados con la fuente única de verdad data/rut/ranges.json.
const (
	// DefaultPersonMin define el límite inferior habitual para personas naturales vivas en Chile.
	DefaultPersonMin = 5_000_000
	// DefaultPersonMax define el límite superior habitual para personas naturales en Chile.
	DefaultPersonMax = 28_000_000

	// DefaultBusinessMin define el límite inferior habitual para personas jurídicas y empresas en Chile.
	DefaultBusinessMin = 70_000_000
	// DefaultBusinessMax define el límite superior habitual para personas jurídicas y empresas en Chile.
	DefaultBusinessMax = 80_000_000

	// DefaultGeneralMin define el límite inferior seguro para generación general.
	DefaultGeneralMin = 1_000_000
	// DefaultGeneralMax define el límite superior seguro para generación general.
	DefaultGeneralMax = 28_000_000
)

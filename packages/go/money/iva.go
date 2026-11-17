package money

import "math"

// DefaultIVARate es la tasa estandar legal del Impuesto al Valor Agregado en Chile (19%).
const DefaultIVARate = 0.19

// IVABreakdown contiene el desglose contable exacto de una transaccion
// garantizando la cuadratura del SII (Neto + IVA = Bruto).
type IVABreakdown struct {
	// Neto es el valor base antes de impuestos.
	Neto int64 `json:"neto"`

	// IVA es el monto correspondiente al impuesto.
	IVA int64 `json:"iva"`

	// Bruto es el total final que incluye el impuesto.
	Bruto int64 `json:"bruto"`
}

// IVAOptions permite personalizar la tasa de calculo del IVA.
type IVAOptions struct {
	// Rate es la tasa decimal del impuesto (por defecto 0.19 para 19%).
	Rate float64
}

// CalculateIVA desglosa un monto en Neto, IVA y Bruto aplicando la regla de cuadratura
// del Servicio de Impuestos Internos (SII).
//
// Si isBruto es true, se asume que amount es el total final (Bruto) y se calcula:
//
//	Neto = Round(Bruto / (1 + Rate))
//	IVA  = Bruto - Neto
//
// Si isBruto es false, se asume que amount es el valor base (Neto) y se calcula:
//
//	IVA   = Round(Neto * Rate)
//	Bruto = Neto + IVA
func CalculateIVA(amount int64, isBruto bool, opts ...IVAOptions) IVABreakdown {
	rate := DefaultIVARate
	if len(opts) > 0 && opts[0].Rate > 0 {
		rate = opts[0].Rate
	}

	if amount == 0 {
		return IVABreakdown{Neto: 0, IVA: 0, Bruto: 0}
	}

	if isBruto {
		bruto := amount
		neto := int64(math.Round(float64(bruto) / (1.0 + rate)))
		iva := bruto - neto
		return IVABreakdown{
			Neto:  neto,
			IVA:   iva,
			Bruto: bruto,
		}
	}

	neto := amount
	iva := int64(math.Round(float64(neto) * rate))
	bruto := neto + iva
	return IVABreakdown{
		Neto:  neto,
		IVA:   iva,
		Bruto: bruto,
	}
}

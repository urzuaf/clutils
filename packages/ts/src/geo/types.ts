/**
 * @file types.ts
 * @description Tipos de datos e interfaces para la División Político-Administrativa (DPA) oficial de Chile.
 */

/**
 * Representa una de las 16 regiones oficiales de Chile según SUBDERE e INE.
 */
export interface Region {
  /** Código numérico oficial INE de 2 dígitos (ej: "13", "16", "05") */
  code: string;
  /** Código estándar internacional ISO 3166-2:CL (ej: "CL-RM", "CL-NB") */
  isoCode: string;
  /** Nombre oficial completo de la región */
  name: string;
  /** Abreviatura oficial (ej: "RM", "NUBLE", "VALPO") */
  abbreviation: string;
  /** Ciudad capital regional */
  capital: string;
  /** Posición geográfica canónica de Norte a Sur (1 a 16) */
  order: number;
}

/**
 * Representa una de las 56 provincias oficiales de Chile según SUBDERE e INE.
 */
export interface Provincia {
  /** Código numérico oficial INE de 3 dígitos (ej: "131", "161") */
  code: string;
  /** Nombre oficial de la provincia */
  name: string;
  /** Código numérico oficial INE de la región a la que pertenece (ej: "13", "16") */
  regionCode: string;
}

/**
 * Representa una de las 346 comunas oficiales de Chile según SUBDERE e INE.
 */
export interface Comuna {
  /** Código Único Territorial oficial del INE de 5 dígitos (ej: "13120", "16101") */
  cut: string;
  /** Nombre oficial de la comuna */
  name: string;
  /** Código oficial de la provincia de 3 dígitos (ej: "131", "161") */
  provinciaCode: string;
  /** Código oficial de la región de 2 dígitos (ej: "13", "16") */
  regionCode: string;
}

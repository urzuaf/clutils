/**
 * @file formatting.ts
 * @description Formateo y parseo de fechas segun el estandar chileno (DD/MM/YYYY) y deteccion de fin de semana.
 */

export type DateFormatStyle = 'numeric' | 'numeric_hyphen' | 'formal' | 'full';

const spanishMonths = [
  'enero',
  'febrero',
  'marzo',
  'abril',
  'mayo',
  'junio',
  'julio',
  'agosto',
  'septiembre',
  'octubre',
  'noviembre',
  'diciembre',
];

const spanishWeekdays = [
  'Domingo',
  'Lunes',
  'Martes',
  'Miércoles',
  'Jueves',
  'Viernes',
  'Sábado',
];

function isLeapYear(year: number): boolean {
  return year % 4 === 0 && (year % 100 !== 0 || year % 400 === 0);
}

function daysInMonth(year: number, month: number): number {
  switch (month) {
    case 1:
    case 3:
    case 5:
    case 7:
    case 8:
    case 10:
    case 12:
      return 31;
    case 4:
    case 6:
    case 9:
    case 11:
      return 30;
    case 2:
      return isLeapYear(year) ? 29 : 28;
    default:
      return 0;
  }
}

/**
 * Formatea una fecha segun las convenciones y estilos locales chilenos.
 *
 * @param date - Objeto Date a formatear.
 * @param style - Estilo visual ('numeric', 'numeric_hyphen', 'formal', 'full'). Por defecto 'numeric'.
 * @returns Cadena formateada.
 *
 * @example
 * ```ts
 * formatDate(new Date(2026, 8, 18)); // "18/09/2026"
 * formatDate(new Date(2026, 8, 18), 'formal'); // "18 de septiembre de 2026"
 * formatDate(new Date(2026, 8, 18), 'full'); // "Viernes, 18 de septiembre de 2026"
 * ```
 */
function parseStrictDigits(s: string): number | null {
  if (s.length === 0) return null;
  let val = 0;
  for (let i = 0; i < s.length; i++) {
    const code = s.charCodeAt(i);
    if (code < 48 || code > 57) {
      return null;
    }
    val = val * 10 + (code - 48);
  }
  return val;
}

export function formatDate(date: Date, style: DateFormatStyle = 'numeric'): string {
  const day = date.getDate();
  const month = date.getMonth();
  const year = date.getFullYear();

  const dayPad = day.toString().padStart(2, '0');
  const monthPad = (month + 1).toString().padStart(2, '0');
  const yearPad = year.toString().padStart(4, '0');

  switch (style) {
    case 'numeric_hyphen':
      return `${dayPad}-${monthPad}-${yearPad}`;
    case 'formal':
      return `${day} de ${spanishMonths[month]} de ${yearPad}`;
    case 'full': {
      const weekday = spanishWeekdays[date.getDay()];
      return `${weekday}, ${day} de ${spanishMonths[month]} de ${yearPad}`;
    }
    case 'numeric':
    default:
      return `${dayPad}/${monthPad}/${yearPad}`;
  }
}

/**
 * Interpreta una cadena en formato chileno estandar ("DD/MM/YYYY" o "DD-MM-YYYY"),
 * asegurando que el primer termino corresponda al dia y el segundo al mes (evitando inversion anglosajona MM/DD/YYYY).
 *
 * Valida estrictamente dias por mes, anos bisiestos y rechaza formatos ISO o textos invalidos.
 *
 * @param text - Cadena de fecha (ej. "18/09/2026", "05-01-2026").
 * @returns Objeto Date configurado a las 12:00:00 para evitar desfaces de huso horario local.
 * @throws {Error} Si el formato es invalido o los dias/meses estan fuera de rango.
 */
export function parseDate(text: string): Date {
  const trimmed = text.trim();
  if (trimmed.length === 0) {
    throw new Error('El texto de fecha está vacío');
  }

  let sep: string;
  if (trimmed.includes('/')) {
    sep = '/';
  } else if (trimmed.includes('-')) {
    sep = '-';
  } else {
    throw new Error("Formato no reconocido: se requiere separador '/' o '-'");
  }

  const parts = trimmed.split(sep);
  if (parts.length !== 3) {
    throw new Error(`Formato de fecha inválido, se esperaba DD${sep}MM${sep}YYYY`);
  }

  const [dayStr, monthStr, yearStr] = parts;
  if (dayStr.length < 1 || dayStr.length > 2 || monthStr.length < 1 || monthStr.length > 2) {
    throw new Error('Componentes de día o mes inválidos');
  }

  if (yearStr.length !== 4) {
    throw new Error(`El año debe tener 4 dígitos: "${yearStr}"`);
  }

  const day = parseStrictDigits(dayStr);
  const month = parseStrictDigits(monthStr);
  const year = parseStrictDigits(yearStr);

  if (day === null || month === null || year === null) {
    throw new Error('Componentes numéricos inválidos');
  }

  if (year <= 0) {
    throw new Error(`Año inválido: ${year}`);
  }

  if (month < 1 || month > 12) {
    throw new Error(`Mes fuera de rango: ${month} (debe ser 1 a 12)`);
  }

  const maxDays = daysInMonth(year, month);
  if (day < 1 || day > maxDays) {
    throw new Error(`Día fuera de rango: ${day} para mes ${month} del año ${year}`);
  }

  // Usar setFullYear para evitar que JS interprete años 0..99 como 1900..1999
  const result = new Date(0);
  result.setFullYear(year, month - 1, day);
  result.setHours(12, 0, 0, 0);
  return result;
}

/**
 * Determina si una fecha cae en sabado o domingo.
 *
 * @param date - Fecha a evaluar.
 * @returns True si es sabado o domingo, false de lunes a viernes.
 */
export function isWeekend(date: Date): boolean {
  const day = date.getDay();
  return day === 0 || day === 6;
}

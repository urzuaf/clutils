/**
 * @file timezones.ts
 * @description Husos horarios oficiales de Chile y conversion de zonas IANA.
 */

/** Identificador IANA oficial para Chile Continental ("America/Santiago"). */
export const TZ_SANTIAGO = 'America/Santiago';

/** Identificador IANA oficial para la Region de Magallanes y Antartica Chilena ("America/Punta_Arenas"). */
export const TZ_PUNTA_ARENAS = 'America/Punta_Arenas';

/** Identificador IANA oficial para Chile Insular Occidental / Rapa Nui ("Pacific/Easter"). */
export const TZ_EASTER = 'Pacific/Easter';

/**
 * Convierte una fecha a un objeto Date cuyos componentes locales (horas, minutos, dia)
 * reflejan la hora de reloj de pared en la zona horaria especificada.
 */
export function toTimezoneDate(date: Date, timeZone: string): Date {
  const parts = new Intl.DateTimeFormat('en-US', {
    timeZone,
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
    second: 'numeric',
    hourCycle: 'h23',
  }).formatToParts(date);

  let y = 0, m = 0, d = 0, h = 0, min = 0, s = 0;
  for (const p of parts) {
    switch (p.type) {
      case 'year': y = parseInt(p.value, 10); break;
      case 'month': m = parseInt(p.value, 10); break;
      case 'day': d = parseInt(p.value, 10); break;
      case 'hour': h = parseInt(p.value, 10); break;
      case 'minute': min = parseInt(p.value, 10); break;
      case 'second': s = parseInt(p.value, 10); break;
    }
  }

  const result = new Date(0);
  result.setFullYear(y, m - 1, d);
  result.setHours(h, min, s, date.getMilliseconds());
  return result;
}

/**
 * Retorna la fecha y hora actual en la zona oficial de Chile Continental ("America/Santiago"),
 * aplicando automaticamente la compensacion de horario de invierno (UTC-4) o verano (UTC-3).
 */
export function nowChile(): Date {
  return toTimezoneDate(new Date(), TZ_SANTIAGO);
}

/**
 * Retorna la fecha y hora actual en la Region de Magallanes y de la Antartica Chilena,
 * la cual se mantiene permanentemente en UTC-3 por decreto supremo.
 */
export function nowMagallanes(): Date {
  return toTimezoneDate(new Date(), TZ_PUNTA_ARENAS);
}

/**
 * Retorna la fecha y hora actual en Rapa Nui / Isla de Pascua ("Pacific/Easter").
 */
export function nowEasterIsland(): Date {
  return toTimezoneDate(new Date(), TZ_EASTER);
}

/**
 * Convierte cualquier Date a la hora de reloj de pared de Chile Continental ("America/Santiago").
 */
export function toChileTime(date: Date): Date {
  return toTimezoneDate(date, TZ_SANTIAGO);
}

/**
 * Convierte cualquier Date a la hora de reloj de pared de Magallanes ("America/Punta_Arenas").
 */
export function toMagallanesTime(date: Date): Date {
  return toTimezoneDate(date, TZ_PUNTA_ARENAS);
}

/**
 * Convierte cualquier Date a la hora de reloj de pared de Rapa Nui ("Pacific/Easter").
 */
export function toEasterIslandTime(date: Date): Date {
  return toTimezoneDate(date, TZ_EASTER);
}

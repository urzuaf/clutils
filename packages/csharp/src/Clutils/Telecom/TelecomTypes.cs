namespace Clutils
{
    /// <summary>
    /// Clasificación oficial del número telefónico chileno según la SUBTEL.
    /// </summary>
    public enum PhoneType
    {
        /// <summary>Número inválido o no reconocido.</summary>
        Unknown = 0,

        /// <summary>Telefonía móvil celular (+56 9 XXXX XXXX).</summary>
        Mobile = 1,

        /// <summary>Telefonía fija (Santiago área 2 o regional de 2 dígitos).</summary>
        Landline = 2,

        /// <summary>Telefonía IP / VoIP (código de área virtual 44).</summary>
        VoIP = 3,

        /// <summary>Línea gratuita / cobro revertido (800 XXX XXX).</summary>
        TollFree = 4,

        /// <summary>Línea de costo compartido (600 XXX XXX / 600 XXX XXXX).</summary>
        SharedCost = 5,

        /// <summary>Número corto oficial de emergencia (130-149, 911, 1404-1488).</summary>
        Emergency = 6,
    }

    /// <summary>
    /// Formatos visuales soportados para teléfonos chilenos.
    /// </summary>
    public enum PhoneFormatStyle
    {
        /// <summary>Formato internacional visual (ej: "+56 9 1234 5678").</summary>
        International = 0,

        /// <summary>Formato nacional visual sin prefijo país (ej: "9 1234 5678").</summary>
        National = 1,

        /// <summary>Estándar compacto E.164 (ej: "+56912345678").</summary>
        E164 = 2,
    }
}

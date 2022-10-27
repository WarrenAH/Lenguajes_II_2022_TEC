namespace calculadora.excepcion
{
    public class OperacionNoValidaException : Exception
    {
        public OperacionNoValidaException(String mensaje)
            : base(mensaje) { }
    }
}

namespace calculadora.excepcion
{
    public class FormatoInvalidoException : Exception
    {
        public FormatoInvalidoException(String mensaje)
            : base(mensaje) { }
    }
}

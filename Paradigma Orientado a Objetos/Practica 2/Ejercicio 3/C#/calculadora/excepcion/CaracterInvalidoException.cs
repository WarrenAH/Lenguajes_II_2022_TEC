namespace calculadora.excepcion
{
    public class CaracterInvalidoException : Exception
    {
        public CaracterInvalidoException(String mensaje)
            : base(mensaje) { }
    }
}

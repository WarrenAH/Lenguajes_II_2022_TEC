package excepcion;

public class OperacionNoValidaException extends Exception{
    public OperacionNoValidaException(String mensaje){
        super(mensaje);
    }
}

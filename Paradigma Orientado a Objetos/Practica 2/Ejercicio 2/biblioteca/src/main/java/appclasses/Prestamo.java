package appclasses;

public class Prestamo {
    private int codigoPrestamo;
    private int codigoSocio;
    private int codigoLibro;
    private String fechaPrestamo;

    public Prestamo(int codigoPrestamo, int codigoSocio, int codigoLibro, String fechaPrestamo) {
        this.codigoPrestamo = codigoPrestamo;
        this.codigoSocio = codigoSocio;
        this.codigoLibro = codigoLibro;
        this.fechaPrestamo = fechaPrestamo;

    }

    public int getCodigoSocioPrestamo() {
        return codigoSocio;
    }

    public int getCodigoLibroPrestamo() {
        return codigoLibro;
    }

    public int getCodigoPrestamo() {
        return codigoPrestamo;
    }

    @Override
    public String toString() {
        return "Prestamo{" +
                "codigoPrestamo=" + codigoPrestamo +
                ", codigoSocio=" + codigoSocio +
                ", codigoLibro=" + codigoLibro +
                ", fechaPrestamo='" + fechaPrestamo + '\'' +
                '}';
    }
}
package appclasses;

import java.util.LinkedList;

public class Socio {
    private int codigoSocio;
    private String nombreSocio;
    private String direccionSocio;
    private int numeroLibroPrestadosSocio;

    public Socio(int codigoSocio, String nombreSocio, String direccionSocio, int numeroLibroPrestadosSocio) {
        this.codigoSocio = codigoSocio;
        this.nombreSocio = nombreSocio;
        this.direccionSocio = direccionSocio;
        this.numeroLibroPrestadosSocio = numeroLibroPrestadosSocio;
    }
    
    public void setNumeroLibroPrestadosSocio(int numeroLibroPrestadosSocio) {
        this.numeroLibroPrestadosSocio = numeroLibroPrestadosSocio;
    }

    public int getCodigoSocio() {
        return codigoSocio;
    }

    public int getNumeroLibroPrestadosSocio() {
        return numeroLibroPrestadosSocio;
    }

    @Override
    public String toString() {
        return "Socio{" +
                "codigoSocio=" + codigoSocio +
                ", nombreSocio='" + nombreSocio + '\'' +
                ", direccionSocio='" + direccionSocio + '\'' +
                ", numeroLibroPrestadosSocio=" + numeroLibroPrestadosSocio +
                '}';
    }

}


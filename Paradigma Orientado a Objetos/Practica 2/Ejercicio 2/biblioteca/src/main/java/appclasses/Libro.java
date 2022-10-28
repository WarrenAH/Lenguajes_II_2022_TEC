package appclasses;

public class Libro {
    private int codigoLibro;
    private String tituloLibro;
    private String autorLibro;
    private boolean disponibilidadLibro;
    private String localizacionLibro;
    private String signaturaLibro;


    public Libro(int codigoLibro, String tituloLibro, String autorLibro, boolean disponibilidadLibro, String localizacionLibro, String signaturaLibro) {
        this.codigoLibro = codigoLibro;
        this.tituloLibro = tituloLibro;
        this.autorLibro = autorLibro;
        this.disponibilidadLibro = disponibilidadLibro;
        this.localizacionLibro = localizacionLibro;
        this.signaturaLibro = signaturaLibro;
    }

    public int getCodigoLibro() {
        return codigoLibro;
    }

    public boolean getDisponibilidadLibro() {
        return disponibilidadLibro;
    }
    
    public void setDisponibilidadLibro(boolean disponibilidadLibro) {
        this.disponibilidadLibro = disponibilidadLibro;
    }

    @Override
    public String toString() {
        return "Libro{" +
                "codigoLibro=" + codigoLibro +
                ", tituloLibro='" + tituloLibro + '\'' +
                ", autorLibro='" + autorLibro + '\'' +
                ", disponibilidadLibro='" + disponibilidadLibro + '\'' +
                ", localizacionLibro='" + localizacionLibro + '\'' +
                ", signaturaLibro='" + signaturaLibro + '\'' +
                '}';
    }
}

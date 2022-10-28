package appclasses;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

public class Biblioteca {
    private String nombre;

    private LinkedList<Libro> crearLibro;

    private LinkedList<Socio> crearSocio;

    private LinkedList<Prestamo> prestamos;
    
    private int maximaCantidadLibrosSocios=3;


    public Biblioteca(String nombre) {
        this.nombre = nombre;

        this.crearLibro = new LinkedList<Libro>();
        this.crearSocio = new LinkedList<Socio>();
        this.prestamos = new LinkedList<Prestamo>();
    }

    public void agregarPrestamo(Prestamo crearPrestamo) {
        boolean revisarLibro=true;
        for (Socio s : this.crearSocio){
            if(crearPrestamo.getCodigoSocioPrestamo()==s.getCodigoSocio()){
            for (Libro l : this.crearLibro){
            if ((l.getCodigoLibro() == crearPrestamo.getCodigoLibroPrestamo()) && (revisarLibro == l.getDisponibilidadLibro())){
                this.prestamos.add(crearPrestamo);
                
                s.setNumeroLibroPrestadosSocio(s.getNumeroLibroPrestadosSocio()+1);
                
                l.setDisponibilidadLibro(false);
                
                System.out.println("El prestamo con el codPrestamo "+crearPrestamo.getCodigoPrestamo()+" se realizo exitosamente" +"\n");
                return;
            }
        }
                
            }
        }
        return;
    }

    public void agregarSocio(Socio crearSocio) {
        this.crearSocio.add(crearSocio);

    }

    public void agregarLibro(Libro crearLibro){
        this.crearLibro.add(crearLibro);

    }
    
    public String sociosMaximaCantidad(){
        List<Socio> socios = this.crearSocio.stream()
        .filter(socio -> socio.getNumeroLibroPrestadosSocio()> this.maximaCantidadLibrosSocios)
        .collect(Collectors.toList());
        
        //return socios.toString();
        
        final StringBuffer sb = new StringBuffer("\n");
        for (Socio s : socios) {
            sb.append(s.toString() + "\n");
        }
        return sb.toString();
    }

    public String printDatos() {
        final StringBuffer sb = new StringBuffer("\n");
        for (Socio crearSocio : this.crearSocio) {
            sb.append(crearSocio.toString() + "\n");
        }
        for (Libro crearLibro : this.crearLibro) {
            sb.append(crearLibro.toString() + "\n");
        }
        for (Prestamo crearPrestamo : this.prestamos) {
            sb.append(crearPrestamo.toString() + "\n");
        }
        return sb.toString();
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("Biblioteca=");
        sb.append(this.nombre);
        sb.append(printDatos());
        return sb.toString();
    }


}
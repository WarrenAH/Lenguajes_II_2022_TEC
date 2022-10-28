package classes;

import java.util.LinkedList;

public class Documento {
    private String nombre;
    private LinkedList<Hoja> hojas;
    
    public Documento(String nombre){
        this.nombre = nombre;
        
        this.hojas = new LinkedList<Hoja>();
    }
    
    public void agregarHoja(Hoja h){
        this.hojas.add(h);
    }
    
    public String printHojas(){
        final StringBuffer sb = new StringBuffer("ListaHojas{\n");
        for (Hoja h : this.hojas){
            sb.append(h.toString()+"\n");
        }
        sb.append('}');
        
        return sb.toString();
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("Documento{");
        sb.append("nombre='").append(nombre).append('\'');
        sb.append(", hojas=");
        sb.append(printHojas());
        sb.append('}');
        return sb.toString();

    }
    
    
}

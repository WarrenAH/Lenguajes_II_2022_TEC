package classes;

import java.util.LinkedList;

public class Hoja {
    private int ancho;
    private int alto;
    
    private LinkedList<ObjRepr> objetos;
    
    public Hoja(int ancho, int alto){
        this.ancho = ancho;
        this.alto = alto;
        
        this.objetos = new LinkedList<ObjRepr>();
    }
    
    public void agregarObjeto(ObjRepr o){
        this.objetos.add(o);
    }
    
    public String printHojas(){
        final StringBuffer sb = new StringBuffer("ListaObjetos{\n");
        for (ObjRepr o : this.objetos){
            sb.append(o.toString()+"\n");
        }
        sb.append('}');
        
        return sb.toString();
    }
    
    @Override
    public String toString(){
        final StringBuffer sb = new StringBuffer("Hoja{");
        sb.append("ancho='").append(ancho).append('\'');
        sb.append("alto='").append(alto).append('\'');
        sb.append(", objetos=");
        sb.append(printHojas());
        sb.append('}');
        return sb.toString();
    }
    
}

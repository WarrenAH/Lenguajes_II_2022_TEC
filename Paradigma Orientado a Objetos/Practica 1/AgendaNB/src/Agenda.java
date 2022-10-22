import java.util.LinkedList;

public class Agenda {
    private LinkedList<Contacto> contactos;
    private LinkedList<Evento> eventos;

    public Agenda() {
        this.contactos = new LinkedList<Contacto>();
        this.eventos = new LinkedList<Evento>();
    }
    
    public void agregarEvento(Evento e){
        this.eventos.add(e);
    }

    //Agregar Contacto
    public void agregarContacto(Contacto c){
        this.contactos.add(c);
    }
    //Eliminar Contacto
    public void eliminarContacto(int index){
        this.contactos.remove(index);
    }
    public void eliminarContacto(String nombre){
        for(Contacto c : this.contactos){
            if (c.getPersona().getNombre().equals(nombre))
                this.contactos.remove(c);
        }
    }
    //modificar Contacto

    //imprimirContactos
    public void imprimirContactos(){
        for(Contacto c : this.contactos)
            c.imprimir();
        
        for(Evento e : this.eventos)
            e.imprimir();
    }
}

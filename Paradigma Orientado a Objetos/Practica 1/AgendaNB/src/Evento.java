public abstract class Evento extends Object {
    private String lugar;
    private String fecha;

    public Evento(String lugar, String fecha) {
        this.lugar = lugar;
        this.fecha = fecha;
    }
    
    public String getLugar() {
        return lugar;
    }
    
    public String getFecha() {
        return fecha;
    }
    
    public abstract void imprimir();

    @Override
    public String toString() {
        return "Evento{" + 
                "lugar='" + lugar + '\'' +
                ", fecha='" + fecha + '\'' + '}';
    }
    
    
}

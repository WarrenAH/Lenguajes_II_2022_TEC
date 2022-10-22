public class EventoT2 extends Evento{ //Este tipo de evento sera para una fiesta familiar.
    private int cantidadInvitados;
    private String platilloComida;
    
    public EventoT2(String lugar, String fecha, int cantidadInvitados, String platilloComida) {
        super(lugar, fecha);
        this.cantidadInvitados = cantidadInvitados;
        this.platilloComida = platilloComida;
    }
    
    @Override
    public void imprimir(){
        System.out.println("EVENTO2 : " + this.toString());
        EventoT2Frame pantalla = new EventoT2Frame();
        pantalla.lugarTxt.setText(this.getLugar());
        pantalla.fechaTxt.setText(this.getFecha());
        pantalla.cantidadInvitadosTxt.setText(String.valueOf(this.cantidadInvitados));
        pantalla.platilloComidaTxt.setText(this.platilloComida);
        pantalla.setVisible(true);
    }

    @Override
    public String toString() {
        return "EventoT2{" + 
                super.toString()+
                "cantidadInvitados=" + cantidadInvitados +
                ", platilloComida='" + platilloComida + '\'' + '}';
    }


}

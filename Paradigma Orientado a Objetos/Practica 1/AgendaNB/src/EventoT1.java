public class EventoT1 extends Evento{ //Este tipo de evento sera para un curso de una universidad X.
    private String requisito;
    private String curso;
    
    public EventoT1(String lugar, String fecha, String requisito, String curso) {
        super(lugar, fecha);
        this.requisito = requisito;
        this.curso = curso;
    }
    
    @Override
    public void imprimir(){
        System.out.println("EVENTO1 : " + this.toString());
        EventoT1Frame pantalla = new EventoT1Frame();
        pantalla.lugarTxt.setText(this.getLugar());
        pantalla.fechaTxt.setText(this.getFecha());
        pantalla.requisitoTxt.setText(this.requisito);
        pantalla.cursoTxt.setText(this.curso);
        pantalla.setVisible(true);
    }

    @Override
    public String toString() {
        return "EventoT1{" + 
                super.toString()+
                "requisito='" + requisito +  '\'' +
                ", curso='" + curso + '\'' + '}';
    }

}

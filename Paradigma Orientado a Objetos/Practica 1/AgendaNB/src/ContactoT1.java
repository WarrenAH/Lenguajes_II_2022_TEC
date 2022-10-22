public class ContactoT1 extends Contacto{
    private String correo;

    public ContactoT1(Persona persona, String direccion, String telefono, String correo) {
        super(persona, direccion, telefono);
        this.correo = correo;
    }

    public ContactoT1(String nombre, int edad, boolean genero, String direccion, String telefono, String correo) {
        super(nombre, edad, genero, direccion, telefono);
        this.correo = correo;
    }


    @Override
    public void imprimir(){
        System.out.println("CONTACTO1 : " + this.toString());
        ContactoT1Frame pantalla = new ContactoT1Frame();
        pantalla.nombreTxt.setText(this.getPersona().getNombre());
        pantalla.edadTxt.setText(String.valueOf(this.getPersona().getEdad()));
        pantalla.generoTxt.setText(String.valueOf(this.getPersona().getGenero()));
        pantalla.direccionTxt.setText(this.getDireccion());
        pantalla.telefonoTxt.setText(this.getTelefono());
        pantalla.correoTxt.setText(this.correo);
        pantalla.setVisible(true);
    }

    @Override
    public String toString() {
        return "ContactoT1{" +
                super.toString() +
                "correo='" + correo + '\'' +
                '}';
    }
}

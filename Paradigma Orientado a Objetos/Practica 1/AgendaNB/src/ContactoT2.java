public class ContactoT2 extends Contacto{
    private String facebook;
    private String telegram;

    public ContactoT2(Persona persona, String direccion, String telefono, String facebook, String telegram) {
        super(persona, direccion, telefono);
        this.facebook = facebook;
        this.telegram = telegram;
    }

    public ContactoT2(String nombre, int edad, boolean genero, String direccion, String telefono, String facebook, String telegram) {
        super(nombre, edad, genero, direccion, telefono);
        this.facebook = facebook;
        this.telegram = telegram;
    }

    @Override
    public void imprimir(){
        System.out.println("CONTACTO2 : " + this.toString());
        ContactoT2Frame pantalla = new ContactoT2Frame();
        pantalla.nombreTxt.setText(this.getPersona().getNombre());
        pantalla.edadTxt.setText(String.valueOf(this.getPersona().getEdad()));
        pantalla.generoTxt.setText(String.valueOf(this.getPersona().getGenero()));
        pantalla.direccionTxt.setText(this.getDireccion());
        pantalla.telefonoTxt.setText(this.getTelefono());
        pantalla.facebookTxt.setText(this.facebook);
        pantalla.telegramTxt.setText(this.telegram);
        pantalla.setVisible(true);
    }

    @Override
    public String toString() {
        return  "ContactoT2{" + 
                super.toString() +
                "facebook='" + facebook + '\'' +
                ", telegram='" + telegram + '\'' +
                "} ";
    }
}

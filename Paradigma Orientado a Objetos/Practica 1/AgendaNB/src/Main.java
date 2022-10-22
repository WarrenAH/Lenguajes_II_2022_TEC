public class Main {
    public static void main(String[] args) {
        Agenda agendaPersonal = new Agenda();
        System.out.println("***AGREGANDO CONTACTOS Y EVENTOS***");
        //agendaPersonal.agregarContacto(new Contacto(new Persona("pablo",20,false),"no se donde","88888888"));
        agendaPersonal.agregarContacto(new ContactoT1("maria",23,true,"detras del mercado","83274245","maria@suCorreo.com" ));
        agendaPersonal.agregarContacto(new ContactoT1("pedro",24,false,"manantial","83728323","pedro@suCorreo.com" ));
        agendaPersonal.agregarContacto(new ContactoT2("luis",30,false,"detras del cementerio","76348342","luisDB","66348342"));
        
        agendaPersonal.agregarEvento(new EventoT1("Santa Clara", "21/10/2022", "Explicar proyecto","Lenguajes de programacion"));
        agendaPersonal.agregarEvento(new EventoT2("Buenos Aires", "22/10/2022", 20,"Arroz con pollo"));
        System.out.println("***IMPRIMIENDO CONTACTOS Y EVENTOS***");
        agendaPersonal.imprimirContactos();
    }
}

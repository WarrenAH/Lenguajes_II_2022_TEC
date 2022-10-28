import appclasses.*;

public class Main {
    public static void main(String[] args) {

        Biblioteca laBiblioteca = new Biblioteca("Biblioteca1");
        
        //AGREGANDO USUARIOS Y LIBROS.

        laBiblioteca.agregarLibro(new Libro(1001, "El Mundo", "Andrey Quesada", false, "Desconocida", "Desconocida"));
        laBiblioteca.agregarSocio(new Socio(101, "Warren", "Aguas Zarcas", 9));


        laBiblioteca.agregarLibro(new Libro(1002, "El Universo", "Enrique Castor", true, "Desconocida", "Desconocida"));
        laBiblioteca.agregarSocio(new Socio(102, "Alvaro", "Ciudad Quesada", 10));

        laBiblioteca.agregarLibro(new Libro(1003, "La Vida", "Carlos Alvarez", true, "Desconocida", "Desconocida"));
        laBiblioteca.agregarSocio(new Socio(103, "Ivan", "Rio Cuarto", 8));
        
        laBiblioteca.agregarSocio(new Socio(104, "Carlos", "Grecia", 3));//ESTE SOCIO NO TIENE MAS DE 3 LIBROS.

        System.out.println(laBiblioteca.toString());

        //AGREGANDO PRESTAMOS.
        laBiblioteca.agregarPrestamo(new Prestamo(1, 101, 1001, "27 de octubre 2022")); //EL LIBRO NO SE ENCUENTRA DISPONIBLE, POR LO TANTO NO HAY PRESTAMO.
        laBiblioteca.agregarPrestamo(new Prestamo(2, 1021, 1002, "27 de octubre 2022")); //EL SOCIO NO SE ENCUENTRA, POR LO TANTO NO HAY PRESTAMO.
        laBiblioteca.agregarPrestamo(new Prestamo(2, 101, 10021, "27 de octubre 2022")); //EL LIBRO NO SE ENCUENTRA, POR LO TANTO NO HAY PRESTAMO.
        laBiblioteca.agregarPrestamo(new Prestamo(3, 103, 1003, "27 de octubre 2022"));

        System.out.println(laBiblioteca.toString());
        
        //BUSCANDO LOS SOCIOS QUE TENGAN MAS DE 3 LIBROS PRESTADOS.
        System.out.println(laBiblioteca.sociosMaximaCantidad());
    }
}

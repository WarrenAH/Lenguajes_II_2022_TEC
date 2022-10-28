import classes.Linea;
import classes.Grupo;
import classes.Cuadrado;
import classes.*;

public class Main {
    public static void main(String[] args) {
        Documento elDocumento = new Documento("Documento1");
        Hoja h1 = new Hoja(20,20);

        h1.agregarObjeto(new Texto(1,"Hola Mundo!!!"));
        h1.agregarObjeto(new Cuadrado(2,4));
        h1.agregarObjeto(new Linea(3,20));

        Grupo g = new Grupo(4);
        g.agregarObjeto(new Texto(5,"Texto"));

        h1.agregarObjeto(g);

        elDocumento.agregarHoja(h1);

        System.out.println(elDocumento.toString());
    }

}

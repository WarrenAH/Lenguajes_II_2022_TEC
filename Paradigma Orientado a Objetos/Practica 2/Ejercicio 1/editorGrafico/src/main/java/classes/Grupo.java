package classes;

import java.util.LinkedList;

public class Grupo extends ObjRepr {
    private LinkedList<ObjRepr> objetos;

    public Grupo(int numObjeto) {
        super(numObjeto);
        this.objetos = new LinkedList<ObjRepr>();
    }

    public void agregarObjeto(ObjRepr o){
        this.objetos.add(o);
    }

    public String printObjetos(){
        final StringBuffer sb = new StringBuffer("ListaObjetos{\n");
        for (ObjRepr o : this.objetos){
            sb.append(o.toString()+"\n");
        }
        sb.append('}');

        return sb.toString();
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("Grupo{");
        sb.append(printObjetos());
        sb.append('}');
        return sb.toString();
    }
}

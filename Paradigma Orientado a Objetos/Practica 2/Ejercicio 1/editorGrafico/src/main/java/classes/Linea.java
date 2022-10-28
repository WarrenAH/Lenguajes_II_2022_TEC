package classes;

public class Linea extends ObjGeo{
    private int largoLinea;

    public Linea(int numObjeto, int largoLinea) {
        super(numObjeto);
        this.largoLinea = largoLinea;
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("Linea{");
        sb.append("largoLinea=").append(largoLinea);
        sb.append('}');
        return sb.toString();
    }
}

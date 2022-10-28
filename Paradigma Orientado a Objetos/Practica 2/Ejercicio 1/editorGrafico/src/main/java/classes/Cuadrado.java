package classes;

public class Cuadrado extends ObjGeo{
    private int largoLado;

    public Cuadrado(int numObjeto, int largoLado) {
        super(numObjeto);
        this.largoLado = largoLado;
    }

    public int getArea(){
        return this.largoLado * this.largoLado;
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("Cuadrado{");
        sb.append("largoLado=").append(largoLado);
        sb.append('}');
        return sb.toString();
    }
}

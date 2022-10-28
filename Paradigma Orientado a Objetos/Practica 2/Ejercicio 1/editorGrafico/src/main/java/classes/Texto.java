package classes;
public class Texto extends ObjRepr{
    private String texto;
    
    public Texto(int numObjeto, String texto){
        super(numObjeto);
        this.texto = texto;
    }
    
    @Override
    public String toString(){
        final StringBuffer sb = new StringBuffer("Texto{");
        sb.append("texto='").append(texto).append('\'');
        sb.append('}');
        return sb.toString();
    }
}

package com.mycompany.calculadora;

import excepcion.CaracterInvalidoException;
import excepcion.FormatoInvalidoException;
import excepcion.OperacionNoValidaException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Main { 
    public static double calcular(String cadena) throws FormatoInvalidoException, CaracterInvalidoException, OperacionNoValidaException{
            String textoNumero=cadena.replace(" ","");
            String[] data =textoNumero.split("[*+-/]");
            List<String>listaNumeros =Arrays.asList(data);
            
            //System.out.println(listaNumeros.toString());
            
            String textoOperador=cadena.replace(" ","");
            String[] data1 =textoOperador.split("[\\d]+");
            List<String>listaOperadores =Arrays.asList(data1);
            
            //System.out.println(listaOperadores.toString());
            
            if(listaOperadores.size()!= listaNumeros.size()){
                throw new FormatoInvalidoException("La cantidad de operadores ingresados "
                        + "(*, +, -, /) son superior o inferiores a los necesitados para hacer operaciones, por favor "
                        + "revise y intentelo de nuevo.");
            }
            
            List<Double> listaResultado = new ArrayList<>();
            
            for (int j = 0; j<listaOperadores.size(); j++) {
                if(j==0){
                    continue;
                }
                
                double numero1=0;
                double numero2=0;
                
                if(j==1){
                    if (listaNumeros.get(j-1).matches("[0-9]+")!=true){
                    throw new CaracterInvalidoException("\"Uno de los numeros ingresados: "+listaNumeros.get(j-1)+", no es un numero, por favor edite su"
                            + " instruccion y vuelva a intentarlo.");     
                }
                         
                    if (listaNumeros.get(j).matches("[0-9]+")!=true){
                    throw new CaracterInvalidoException("\"Uno de los numeros ingresados: "+listaNumeros.get(j)+", no es un numero, por favor edite su"
                            + " instruccion y vuelva a intentarlo.");     
                    }
                    numero1 = Integer.parseInt(listaNumeros.get(j-1));
                    numero2 = Integer.parseInt(listaNumeros.get(j));  
                }
                
                if(j>1){        
                    if (listaNumeros.get(j).matches("[0-9]+")!=true){
                    throw new CaracterInvalidoException("\"Uno de los numeros ingresados: "+listaNumeros.get(j)+", no es un numero, por favor edite su"
                            + " instruccion y vuelva a intentarlo.");     
                    }
                    numero1 = listaResultado.get(0);
                    numero2 = Integer.parseInt(listaNumeros.get(j));
                    listaResultado.remove(0);
                }       
                
                if((listaOperadores.get(j).equals("*")==false)&&(listaOperadores.get(j).equals("+")==false)&&
                        (listaOperadores.get(j).equals("/")==false)&&(listaOperadores.get(j).equals("-")==false)){
                throw new CaracterInvalidoException("Uno de los operadores ingresados: "+listaOperadores.get(j)+", no es "
                        + "valido, solo se permite: *, +, -, /.");
                }

                if(listaOperadores.get(j).equals("*")){
                    if(Double.valueOf(numero1 * numero2).isNaN()){
                        throw new OperacionNoValidaException("La operacion : "+numero1+" * "+numero2+", no es valida");
                    }
                    if((j==1)&&(listaOperadores.size()==2)){  
                        return numero1 * numero2;
                    }
                    
                    if((j==1)&&(listaOperadores.size()>=2)){
                        listaResultado.add(numero1 * numero2);
                    }
                    
                    if(j>1){
                        listaResultado.add(numero1 * numero2);
                    }
                }

                if(listaOperadores.get(j).equals("+")){
                    if(Double.valueOf(numero1 + numero2).isNaN()){
                        throw new OperacionNoValidaException("La operacion : "+numero1+" + "+numero2+", no es valida");
                    }
                    if((j==1)&&(listaOperadores.size()==2)){ 
                        return numero1 + numero2;
                    }
                    if((j==1)&&(listaOperadores.size()>=2)){
                        listaResultado.add(numero1 + numero2);
                    }
                    if(j>1){
                        listaResultado.add(numero1 + numero2);
                    }
                }

                if(listaOperadores.get(j).equals("/")){
                    if(Double.valueOf(numero1 / numero2).isNaN()){
                        throw new OperacionNoValidaException("La operacion : "+numero1+" / "+numero2+", no es valida");
                    }
                    if((j==1)&&(listaOperadores.size()==2)){  
                        return numero1 / numero2;
                    }
                    if((j==1)&&(listaOperadores.size()>=2)){
                        listaResultado.add(numero1 / numero2);
                    }
                    if(j>1){
                        listaResultado.add(numero1 / numero2);
                    }
                }

                if(listaOperadores.get(j).equals("-")){
                    if(Double.valueOf(numero1 - numero2).isNaN()){
                        throw new OperacionNoValidaException("La operacion : "+numero1+" - "+numero2+", no es valida");
                    }
                    if((j==1)&&(listaOperadores.size()==2)){  
                        return numero1 - numero2;
                    }
                    if((j==1)&&(listaOperadores.size()>=2)){
                        listaResultado.add(numero1 - numero2);
                    }
                    if(j>1){
                        listaResultado.add(numero1 - numero2);
                    }
                }
            }
            return listaResultado.get(0);
    }
    
    public static void main(String[] args) throws FormatoInvalidoException, CaracterInvalidoException, OperacionNoValidaException {
        System.out.println("El resultado: "+ calcular("9 / 8 + 1 - 2 * 90"));
    }
}

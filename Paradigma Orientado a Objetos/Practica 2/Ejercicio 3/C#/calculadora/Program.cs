using calculadora.excepcion;
using System;
using System.Text.RegularExpressions;

namespace Demo
{
    public class Inicio
    {
            static public Double calcular(String cadena)
        {
            String textoNumero = cadena.Replace(" ", "");
            Regex regexNumero = new Regex("[*+-/]");
            String[] data = regexNumero.Split(textoNumero);
            List<String> listaNumeros = data.ToList();

            //Console.WriteLine(String.Join(",", listaNumeros));

            String textoOperador = cadena.Replace(" ", "");
            Regex regexOperador = new Regex("[\\d]+");
            String[] data1 = regexOperador.Split(textoOperador);
            List<String> listaOperadores = data1.ToList();
            listaOperadores.RemoveAt(listaOperadores.Count - 1);

            //Console.WriteLine(String.Join(",", listaOperadores));

            if (listaOperadores.Count != listaNumeros.Count)
            {
                throw new FormatoInvalidoException("La cantidad de operadores ingresados "
                        + "(*, +, -, /) son superior o inferiores a los necesitados para hacer operaciones, por favor "
                        + "revise y intentelo de nuevo.");
            }

            List<Double> listaResultado = new List<Double>();

            Regex regexComparar = new Regex("[0-9]+");

            for (int j = 0; j < listaOperadores.Count; j++)
            {
                if (j == 0)
                {
                    continue;
                }
                double numero1 = 0;
                double numero2 = 0;

                if (j == 1)
                {
                    if (regexComparar.IsMatch(listaNumeros.ElementAt(j-1)) != true)
                    {
                        throw new CaracterInvalidoException("\"Uno de los numeros ingresados: " + listaNumeros.ElementAt(j - 1) + ", no es un numero, por favor edite su"
                                + " instruccion y vuelva a intentarlo.");
                    }

                    if (regexComparar.IsMatch(listaNumeros.ElementAt(j)) != true)
                    {
                        throw new CaracterInvalidoException("\"Uno de los numeros ingresados: " + listaNumeros.ElementAt(j) + ", no es un numero, por favor edite su"
                                + " instruccion y vuelva a intentarlo.");
                    }
                    numero1 = int.Parse(listaNumeros.ElementAt(j - 1));
                    numero2 = int.Parse(listaNumeros.ElementAt(j));
                }

                if (j > 1)
                {
                    if (regexComparar.IsMatch(listaNumeros.ElementAt(j)) != true)
                    {
                        throw new CaracterInvalidoException("\"Uno de los numeros ingresados: " + listaNumeros.ElementAt(j) + ", no es un numero, por favor edite su"
                                + " instruccion y vuelva a intentarlo.");
                    }
                    numero1 = listaResultado.ElementAt(0);
                    numero2 = int.Parse(listaNumeros.ElementAt(j));
                    listaResultado.RemoveAt(0);
                }

                if ((listaOperadores.ElementAt(j).Contains("*") == false) && (listaOperadores.ElementAt(j).Contains("+") == false) &&
                        (listaOperadores.ElementAt(j).Contains("/") == false) && (listaOperadores.ElementAt(j).Contains("-") == false))
                {
                    throw new CaracterInvalidoException("Uno de los operadores ingresados: " + listaNumeros.ElementAt(j) + ", no es "
                            + "valido, solo se permite: *, +, -, /.");
                }

                if (listaOperadores.ElementAt(j).Contains("*"))
                {
                    if (Double.IsNaN(numero1 * numero2))
                    {
                        throw new OperacionNoValidaException("La operacion : " + numero1 + " * " + numero2 + ", no es valida");
                    }
                    if ((j == 1) && (listaOperadores.Count == 2))
                    {
                        return numero1 * numero2;
                    }

                    if ((j == 1) && (listaOperadores.Count >= 2))
                    {
                        listaResultado.Add(numero1 * numero2);
                    }

                    if (j > 1)
                    {
                        listaResultado.Add(numero1 * numero2);
                    }
                }


                if (listaOperadores.ElementAt(j).Contains("+"))
                {
                    if (Double.IsNaN(numero1 + numero2))
                    {
                        throw new OperacionNoValidaException("La operacion : " + numero1 + " * " + numero2 + ", no es valida");
                    }
                    if ((j == 1) && (listaOperadores.Count == 2))
                    {
                        return numero1 + numero2;
                    }

                    if ((j == 1) && (listaOperadores.Count >= 2))
                    {
                        listaResultado.Add(numero1 + numero2);
                    }

                    if (j > 1)
                    {
                        listaResultado.Add(numero1 + numero2);
                    }
                }


                if (listaOperadores.ElementAt(j).Contains("/"))
                {
                    if (Double.IsNaN(numero1 / numero2))
                    {
                        throw new OperacionNoValidaException("La operacion : " + numero1 + " * " + numero2 + ", no es valida");
                    }
                    if ((j == 1) && (listaOperadores.Count == 2))
                    {
                        return numero1 / numero2;
                    }

                    if ((j == 1) && (listaOperadores.Count >= 2))
                    {
                        listaResultado.Add(numero1 / numero2);
                    }

                    if (j > 1)
                    {
                        listaResultado.Add(numero1 / numero2);
                    }
                }


                if (listaOperadores.ElementAt(j).Contains("-"))
                {
                    if (Double.IsNaN(numero1 - numero2))
                    {
                        throw new OperacionNoValidaException("La operacion : " + numero1 + " * " + numero2 + ", no es valida");
                    }
                    if ((j == 1) && (listaOperadores.Count == 2))
                    {
                        return numero1 - numero2;
                    }

                    if ((j == 1) && (listaOperadores.Count >= 2))
                    {
                        listaResultado.Add(numero1 - numero2);
                    }

                    if (j > 1)
                    {
                        listaResultado.Add(numero1 - numero2);
                    }
                }
            }
            return listaResultado.ElementAt(0);
        }
        static public void Main()
        {
            Console.WriteLine("El resultado: " + calcular("9 / 8 + 1 - 2 * 90"));
        }
    }
}
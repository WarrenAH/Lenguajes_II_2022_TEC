module Main where
--miembro 
miembro elem lista =
    or (map comparacion lista)
    --any (==True) (map comparacion lista)
    where
        comparacion x = (x == elem)

--remove_if
remove_if fun lista =
    concat (map (\x -> if (fun x) then []
                       else [x] )
                lista)

vecinos nodo grafo = do
    let resultado =  concat (map (\x -> if (head (head x)) == nodo then [x]
                            else [] ) 
                            grafo)
    if (resultado == []) then []
    else head (tail (head resultado))

extender ruta grafo =
    remove_if (==[])
    (map (\x -> if (miembro x ruta) then []
                       else x:ruta )
        (vecinos (head ruta) grafo))

prof ini fin grafo =
    prof_aux fin [[ini]] grafo

prof_aux fin rutas grafo =
    if rutas==[] then []
    else if (fin == (head (head rutas))) then
        --reverse (head rutas)
        reverse (head rutas):prof_aux fin ((extender (head rutas) grafo)++(tail rutas)) grafo
    else
        prof_aux fin ((extender (head rutas) grafo)++(tail rutas)) grafo

rutaCorta :: [[a]] -> [a]
rutaCorta []=[]
rutaCorta [y] = y
rutaCorta (x:y:lista)
 |length x > length y = rutaCorta (y:lista)            
 |otherwise = rutaCorta (x:lista)

main :: IO ()
main = do
    let grafo=[ [["inicio"],["2"]],
                [["2"],["inicio","3","8"]],      
                [["9"],["8","3"]],  
                [["8"],["2","9"]],  
                [["3"],["9","4","2"]],
                [["4"],["3","10"]],  
                [["10"],["16","4"]],          
                [["16"],["10","22"]],
                [["22"],["16","21"]],
                [["21"],["22","15"]],
                [["15"],["21","14"]],
                [["14"],["13","15","20"]],
                [["13"],["14","7"]],
                [["7"],["1","13"]],
                [["1"],["7"]],
                [["20"],["14","26"]],
                [["26"],["20","27"]],
                [["27"],["26","28"]],
                [["28"],["27","29","34"]],
                [["29"],["28","23"]],
                [["23"],["29","17"]],
                [["17"],["23","11"]],
                [["11"],["17","5"]],
                [["5"],["6","11"]],
                [["6"],["5"]],
                [["34"],["28","33","35"]],
                [["35"],["34","36"]],
                [["36"],["35","30"]],
                [["30"],["36","24"]],
                [["24"],["30","18"]],
                [["18"],["12","24"]],
                [["12"],["18"]],
                [["33"],["32","34"]],
                [["33"],["32","34"]],
                [["32"],["31","33","fin"]],
                [["31"],["32","25"]],
                [["25"],["31","19"]],
                [["19"],["25"]],
                [["fin"],["32"]]                                  
                 ]
    
    let generarRutas=prof "inicio" "fin" grafo
    let rutaMasCorta=rutaCorta generarRutas

    print(rutaMasCorta)

    --IMPRESION:
    --["inicio","2","3","4","10","16","22","21","15","14","20","26","27","28","34","33","32","fin"]

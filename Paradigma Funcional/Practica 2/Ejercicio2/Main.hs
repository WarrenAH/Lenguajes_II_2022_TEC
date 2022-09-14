module Main where

import Data.Ord (comparing)
import Data.List (sortBy)

combinar :: [a] -> [a] -> [[a]]
combinar p1 p2 = [[a,b] | (a, b) <- zip p1 p2]

convertir :: [String] -> [(String, Integer)]
convertir [] = []
convertir (k:v:t) = (k, read v) : convertir t

miembro elem lista =
    or (map comparacion lista)
    where
        comparacion x = (x == elem)

remove_if fun lista =
    concat (map (\x -> if (fun x) then []
                       else [x] )
                lista)

utilizarVertices nodo grafo = do
    let resultado =  concat (map (\x -> if (head (head x)) == nodo then [x]
                            else [] ) 
                            grafo)
    if (resultado == []) then []
    else head (tail (head resultado))

extenderConPesos ruta grafo =
    remove_if (==[])
    (map (\x -> if (miembro x ruta) then []
                       else x:ruta )
        (ordenarPesos (head ruta) grafo))

prof ini fin grafo =
    prof_aux fin [[ini]] grafo

ordenar :: [[a]] -> [a]
ordenar []=[]
ordenar [y] = y
ordenar (x:y:list)
 |length x > length y = ordenar (y:list)            
 |otherwise = ordenar (x:list)
 
utilizarPesos nodo grafo = do
    let resultado =  concat (map (\x -> if (head(head x)) == nodo then [tail x]
                            else [] ) 
                            grafo)
    if (resultado == []) then []
    else head (tail (head resultado))

prof_aux fin rutas grafo =
    if rutas==[] then []
    else if (fin == (head (head rutas))) then
        reverse (head rutas):prof_aux fin ((extenderConPesos (head rutas) grafo)++(tail rutas)) grafo
    else
        prof_aux fin ((extenderConPesos (head rutas) grafo)++(tail rutas)) grafo
    

ordenarPesos nodo grafo = do
    let {combinarVecinosPesos= combinar (utilizarVertices nodo grafo) (utilizarPesos nodo grafo); concatenar= concat combinarVecinosPesos; 
    convertirStringInt= convertir concatenar; ordenarConcat= sortBy (comparing (\(x,y) -> abs y)) convertirStringInt; 
    eliminarPesos= convertirStringInt >>= \(x,y) -> [x]} in eliminarPesos

main :: IO ()
main = do
    let grafo=[ [["i"],["a","b"], ["3","6"]],     
                [["a"],["i","c","d"], ["3","6","5"]], 
                [["b"],["i","c","d"], ["3","6","5"]],  
                [["c"],["a","b","x"], ["3","6","5"]],  
                [["d"],["a","b","f"], ["3","6","5"]], 
                [["f"],["d"], ["3"]],          
                [["x"],["c"], ["3"]]           
                 ]   
    
    let lista=prof "i" "f" grafo

    let lista_nueva=ordenar lista

    print(lista_nueva)

    --IMPRESION:
    --["i","a","d","f"]

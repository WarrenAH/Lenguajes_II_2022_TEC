module Main where

import Text.Printf

data Producto = Producto { nombre :: String
                    , descripcion :: String
                    , montoVenta :: Int
                     } deriving Show
                     
reducirLista listaProductos =
    sum listaProductos 

calcularMontoImpuesto impuesto listaProductos =
    map articulos listaProductos 
    where
        articulos prod= (impuesto* fromIntegral (montoVenta prod))

revisarMontosImpuesto rangoPagoImpuestos listaProductos =
    filter articulos listaProductos 
    where
        articulos prod = ((montoVenta prod) >= rangoPagoImpuestos)

calcularMontoVenta listaProductos =
    map articulos listaProductos 
    where
        articulos prod= ((montoVenta prod))

main :: IO ()
main = do
    let porcentajeImpuesto= 0.13
    let rangoPagoImpuestos=20000
    
    let listaProductos=[
            (Producto "tarjeta madre" "Asus" 54200),
            (Producto "mouse" "alámbrico" 15000),
            (Producto "teclado" "gamer con luces" 30000),
            (Producto "memoria ssd" "2 gb" 65000),
            (Producto "cable video" "display port 1m" 18000)]
    
    let buscarImpuestos= calcularMontoImpuesto porcentajeImpuesto (revisarMontosImpuesto rangoPagoImpuestos listaProductos)
    let reducir1= reducirLista buscarImpuestos

    let buscarPreciosProducto= calcularMontoVenta listaProductos
    let reducir2 = reducirLista buscarPreciosProducto

    let montoTotal= reducir1 + fromIntegral reducir2
    
    print("")
    putStrLn $ show "Sus impuestos de esta compra son de:" ++ " " ++ show reducir1
    putStrLn $ show "Su total sin impuestos es de:" ++ " " ++ show reducir2
    putStrLn $ show "Su total con impuestos es de:" ++ " " ++ show montoTotal

    --IMPRESION:
    --""
    --"Sus impuestos de esta compra son de:" 19396.0
    --"Su total sin impuestos es de:" 182200
    --"Su total con impuestos es de:" 201596.0
module Main where
--------------------------------------------------------
--EJERCICIO 1
filter1 :: (a -> Bool) -> [a] -> [a]
filter1 _pred [] = []
filter1 pred (x:xs)
  | pred x = x : filter1 pred xs
  | otherwise = filter1 pred xs

subcadena :: String -> String -> Bool
subcadena (_:_) [] = False
subcadena xs ys
    | prefijo xs ys = True
    | subcadena xs (tail ys) = True
    | otherwise = False

prefijo :: String -> String -> Bool
prefijo [] _ = True
prefijo (_:_) [] = False
prefijo (x:xs) (y:ys) = (x == y) && prefijo xs ys

sub_cadenas :: String -> [String] -> [String]
sub_cadenas str lista = filter1 (subcadena str) lista
--------------------------------------------------------
--EJERCICIO 2
sub_conjunto :: Eq a => [a] -> [a] -> Bool
sub_conjunto [] [] = True
sub_conjunto _ [] = False
sub_conjunto [] _ = True
sub_conjunto (x:xs) (y:ys) | x == y = sub_conjunto xs ys | otherwise = sub_conjunto (x:xs) ys
--------------------------------------------------------
--EJERCICIO 3
aplanar :: [[a]] -> [a]
aplanar [] = []
aplanar (x:xs) = x ++ aplanar xs      
--------------------------------------------------------
--EJERCICIO 4
map1:: (a -> b -> b) -> b -> [a] -> b
map1 a b = go
          where
            go []     = b
            go (x:xs) = x `a` go xs

aplanarConMap :: [[a]] -> [a]
aplanarConMap = map1 (++) []

main :: IO ()
main = do
    print("--------------------------------------------------------")
    print("EJERCICIO 1")
    print(sub_cadenas "la" ["la casa", "el perro", "pintando la cerca"])
    print("--------------------------------------------------------")
    print("EJERCICIO 2")
    print(sub_conjunto [] [1,2,3,4,5])
    print(sub_conjunto [1,2,3] [1,2,3,4,5])
    print(sub_conjunto [1,2,6] [1,2,3,4,5])
    print("--------------------------------------------------------")
    print("EJERCICIO 3")
    print(aplanar [[1,2],[3],[4,5],[6,7]])
    print("--------------------------------------------------------")
    print("EJERCICIO 4")
    print(aplanarConMap [[1,2],[3],[4,5],[6,7]])
    print("--------------------------------------------------------")

--IMPRESION:
--"--------------------------------------------------------"
--"EJERCICIO 1"
--["la casa","pintando la cerca"]
--"--------------------------------------------------------"
--"EJERCICIO 2"
--True
--True
--False
--"--------------------------------------------------------"
--"EJERCICIO 3"
--[1,2,3,4,5,6,7]
--"--------------------------------------------------------"
--"EJERCICIO 4"
--[1,2,3,4,5,6,7]
--"--------------------------------------------------------"
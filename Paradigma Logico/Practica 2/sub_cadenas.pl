subcadena(X,Y) :-
    string_to_list(X, XL),
    string_to_list(Y, YL),
    append(_,T,YL) ,
    append(XL,_,T) ,
    X \= [], !.

identicos(X, Y) :-
    subcadena(X,Y).

sub_cadenas(Subcadena, ListaCadena, Resultado) :-
    include(identicos(Subcadena), ListaCadena, Resultado).

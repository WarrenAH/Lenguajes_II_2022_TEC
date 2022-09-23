eliminar(_,[],[]).
eliminar(E,[E|T],X):- eliminar(E, T, X), !.
eliminar(E, [H|T], [H|Ts]):- eliminar(E, T, Ts).

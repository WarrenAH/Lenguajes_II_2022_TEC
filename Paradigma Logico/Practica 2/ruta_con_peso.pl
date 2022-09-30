%        a ---- c ---- x
%      /   \  /
%     /     \/
%   i       X
%     \    / \
%      \  /   \
%        b ---- d ---- f

conectado(i,a,3).
conectado(i,b,4).
conectado(a,i,6).
conectado(a,c,9).
conectado(a,d,2).
conectado(b,i,6).
conectado(b,c,7).
conectado(b,d,3).
conectado(c,a,5).
conectado(c,x,1).
conectado(c,b,9).
conectado(d,a,6).
conectado(d,f,9).
conectado(d,b,5).
conectado(x,c,3).
conectado(f,d,2).

% ruta_con_peso mostrara todas las rutas posibles de un camino con su
% coste de pesos, para ejecutar seguir el siguiente ejemplo:
%
% ruta_con_peso(x, f, Costo, Ruta).
%
% RESULTADO:
% Costo = 30,
% Ruta = [x, c, a, i, b, d] ;
% Costo = 19,
% Ruta = [x, c, a, d] ;
% Costo = 32,
% Ruta = [x, c, b, i, a, d] ;
% Costo = 24,
% Ruta = [x, c, b, d] ;
% false.

ruta_con_peso(Ini, Fin, Peso, Ruta) :- ruta_con_peso(Ini, Fin, Peso, [], Ruta).
ruta_con_peso(Ini, Fin, Peso, Visto, [Ini]) :-
    \+ memberchk(Ini, Visto),
    conectado(Ini, Fin, Peso).
ruta_con_peso(Ini, Dato, Peso, Visto, [Ini|Tail]) :-
    \+ memberchk(Ini, Visto),
    conectado(Ini, Fin, Numero),
    ruta_con_peso(Fin, Dato, Numero1, [Ini|Visto], Tail),
    \+ memberchk(Ini, Tail),
    Peso is Numero + Numero1.

% mejor_ruta_con_peso mostrara la mejor ruta segun el peso ingresado en
% los caminos, para ejecutar se debe seguir el siguiente ejemplo:
%
% mejor_ruta_con_peso(x, f, Costo, Ruta).
%
% RESULTADO:
% Costo = 19,
% Ruta = [x, c, a, d].

mejor_ruta_con_peso(Ini, Fin, Peso, Ruta) :-
    ruta_con_peso(Ini, Fin, Peso, Ruta),
    \+ (ruta_con_peso(Ini, Fin, MenorPeso, OtraRuta),
        OtraRuta \= Ruta,
        MenorPeso =< Peso).

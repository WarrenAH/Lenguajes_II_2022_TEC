aplanar([], []) :- !.
aplanar([Head|Tail], Lista) :-
    !,
    aplanar(Head, NH),
    aplanar(Tail, NT),
    append(NH, NT, Lista).
aplanar(L, [L]).

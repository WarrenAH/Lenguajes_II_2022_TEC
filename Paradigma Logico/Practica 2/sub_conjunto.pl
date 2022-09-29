sub_conjunto([],_).
sub_conjunto([H|T],Y):-
        member(H,Y),
        select(H,Y,Z),
        sub_conjunto(T,Z).

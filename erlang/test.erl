-module(test).

-export([second/1, oh_god/1, fizzbuzz/1, base/1, handred/1]).

second([_,X|_]) -> X.

oh_god(N) ->
    if N =:= 2 -> might_succeed;
    true -> always_does
end.

fizzbuzz(0) -> io:format("fin~n");
fizzbuzz(Num) when Num > 0 -> 
    if Num rem 15 =:= 0 -> io:format("~p : fizzbuzz~n", [Num])
    ; Num rem 5 =:= 0 -> io:format("~p : fizz~n", [Num])
    ; Num rem 3 =:= 0 -> io:format("~p : buzz~n", [Num])
    ; true -> io:format("~p~n", [Num])
    end,
    fizzbuzz(Num - 1).

base(A) ->
    B = A + 1,
    F = fun() -> A * B end,
    F().

base2(A) ->
    B = A + 1,
    F = fun() -> A * B end,
    F().

show(M, N) when N < 0 -> 
    io:format("~p~n", M),
show(M, _) ->
    io:format("~p~n", <<"END">>).
    

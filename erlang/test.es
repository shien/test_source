-module(argtest).
-export([main/1]).

main([A]) ->
    if
        is_atom(A) ->
            io:format("atom\n");
        is_list(A) ->
            io:format("list\n");
        true ->
            io:format("unknown\n")
        end,
        init:stop().

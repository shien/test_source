-module(elixir_study).
-export([main/1]).

main(_) ->
    Range = lists:seq(1, 1000),
    answer_elixir_probrem(273, Range).

answer_elixir_probrem(Actual, Range) -> 
    get_number(Actual, Range).

get_number(Actual, Range) ->
    Upper = lists:max(Range),
    Lower = lists:min(Range),
    Mid = (Upper - Lower) / 2 + Lower,
    check_mid_number(Actual, Mid, Upper,Lower).


check_mid_number(Actual, Mid, Upper, Lower) when Mid > Actual  ->
    Range = lists:seq(Lower, Mid -1),
    get_number(Actual, Range);
check_mid_number(Actual, Mid, Upper, Lower) when Mid < Actual  ->
    Range = lists:seq(Mid + 1, Upper),
    get_number(Actual, Range);
check_mid_number(Actual, _, _, _) ->
    io:format("actual : ~p~n", [Actual]).

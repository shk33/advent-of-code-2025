defmodule Day3Part2 do
  def solve(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.map(&solve_line/1)
    |> Enum.sum()
  end

  defp solve_line(line) do
    if String.length(line) < 12 do
      0
    else
      n = String.length(line)
      k = 12
      to_remove = n - k
      
      {final_stack, _} = 
        line
        |> String.graphemes()
        |> Enum.reduce({[], to_remove}, fn digit, {stack, removals} ->
            {new_stack, new_removals} = pop_smaller(digit, stack, removals)
            {[digit | new_stack], new_removals}
        end)

      final_stack
      |> Enum.reverse()
      |> Enum.take(k)
      |> Enum.join()
      |> String.to_integer()
    end
  end

  defp pop_smaller(digit, [h|t], removals) when removals > 0 and digit > h do
    pop_smaller(digit, t, removals - 1)
  end
  defp pop_smaller(_digit, stack, removals) do
    {stack, removals}
  end
end

input = File.read!("day3/part2/input.txt")
result = Day3Part2.solve(input)
IO.puts("The new total output joltage is: #{result}")
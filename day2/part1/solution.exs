defmodule Day2Part1 do
  def solve(input_str) do
    ranges = parse_ranges(input_str)
    max_id = find_max_id(ranges)
    
    1..max_id
    |> Stream.map(&to_string/1)
    |> Stream.map(&(&1 <> &1))
    |> Stream.map(&String.to_integer/1)
    |> Stream.take_while(&(&1 <= max_id))
    |> Enum.reduce(MapSet.new(), fn invalid_id, acc ->
      if in_any_range?(invalid_id, ranges) do
        MapSet.put(acc, invalid_id)
      else
        acc
      end
    end)
    |> Enum.sum()
  end

  defp parse_ranges(input_str) do
    input_str
    |> String.trim()
    |> String.split(",")
    |> Enum.map(fn r_str ->
      [start_str, end_str] = String.split(r_str, "-")
      {String.to_integer(start_str), String.to_integer(end_str)}
    end)
  end

  defp find_max_id(ranges) do
    ranges
    |> Enum.map(&elem(&1, 1))
    |> Enum.max()
  end

  defp in_any_range?(id, ranges) do
    Enum.any?(ranges, fn {start, finish} ->
      id >= start and id <= finish
    end)
  end
end

input = File.read!("day2/part1/input.txt")
result = Day2Part1.solve(input)
IO.puts("The sum of all invalid IDs is: #{result}")
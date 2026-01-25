defmodule Day2Part2 do
  def solve(input_str) do
    ranges = parse_ranges(input_str)
    max_id = find_max_id(ranges)
    
    found_ids = 
      Stream.unfold(1, fn base_num ->
        base_s = to_string(base_num)
        first_rep_s = base_s <> base_s
        
        case String.to_integer(first_rep_s) do
          first_rep_n when first_rep_n > max_id ->
            nil
          first_rep_n ->
            repeated_ids = 
              Stream.unfold(first_rep_s, fn current_s ->
                next_s = current_s <> base_s
                try do
                  next_n = String.to_integer(next_s)
                  if next_n > max_id, do: nil, else: {next_n, next_s}
                rescue
                  _ -> nil
                end
              end)
              |> Enum.to_list()
            
            { [first_rep_n | repeated_ids], base_num + 1 }
        end
      end)
      |> Stream.flat_map(&(&1))
      |> Enum.reduce(MapSet.new(), fn invalid_id, acc ->
        if in_any_range?(invalid_id, ranges) do
          MapSet.put(acc, invalid_id)
        else
          acc
        end
      end)

    Enum.sum(found_ids)
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

input = File.read!("day2/part2/input.txt")
result = Day2Part2.solve(input)
IO.puts("The sum of all invalid IDs is: #{result}")

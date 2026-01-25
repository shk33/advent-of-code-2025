defmodule Day1Part2 do
  def solve(input_path) do
    File.stream!(input_path)
    |> Enum.reduce({50, 0}, fn line, {current_position, zero_count} ->
      line = String.trim(line)
      if String.length(line) > 0 do
        {direction, distance_str} = String.split_at(line, 1)
        distance = String.to_integer(distance_str)

        new_zero_count =
          case direction do
            "R" ->
              zero_count + Integer.floor_div(current_position + distance, 100) - Integer.floor_div(current_position, 100)
            "L" ->
              zero_count + Integer.floor_div(current_position - 1, 100) - Integer.floor_div(current_position - distance - 1, 100)
          end

        new_position =
            case direction do
                "R" -> current_position + distance
                "L" -> current_position - distance
            end

        {new_position, new_zero_count}
      else
        {current_position, zero_count}
      end
    end)
    |> elem(1)
  end
end

password = Day1Part2.solve("day1/part2/input1.txt")
IO.puts("The password is: #{password}")
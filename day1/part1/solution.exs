defmodule Day1Part1 do
  def solve(input_path) do
    File.stream!(input_path)
    |> Enum.reduce({50, 0}, fn line, {current_position, zero_count} ->
      line = String.trim(line)
      if String.length(line) > 0 do
        {direction, distance_str} = String.split_at(line, 1)
        distance = String.to_integer(distance_str)

        new_position =
          case direction do
            "R" -> current_position + distance
            "L" -> current_position - distance
          end

        new_position = rem(new_position, 100)

        new_position = if new_position < 0, do: new_position + 100, else: new_position

        new_zero_count =
          if new_position == 0 do
            zero_count + 1
          else
            zero_count
          end

        {new_position, new_zero_count}
      else
        {current_position, zero_count}
      end
    end)
    |> elem(1)
  end
end

password = Day1Part1.solve("day1/part1/input1.txt")
IO.puts("The password is: #{password}")
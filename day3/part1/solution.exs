defmodule Day3Part1 do
  def solve(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.map(&solve_line/1)
    |> Enum.sum()
  end

  defp solve_line(line) do
    chars = String.graphemes(line)
    
    max_joltage = 
      for i <- 0..(length(chars) - 2),
          j <- (i + 1)..(length(chars) - 1) do
        joltage_str = Enum.at(chars, i) <> Enum.at(chars, j)
        String.to_integer(joltage_str)
      end
      |> Enum.max()
      
    max_joltage || 0
  end
end

input = File.read!("day3/part1/input.txt")
result = Day3Part1.solve(input)
IO.puts("The total output joltage is: #{result}")
def solve
  input_path = File.join(File.dirname(__FILE__), 'input.txt')
  lines = File.readlines(input_path)

  total_joltage = 0
  lines.each do |line|
    line = line.strip
    next if line.empty?

    max_line_joltage = 0
    (0...line.length).each do |i|
      (i + 1...line.length).each do |j|
        joltage_str = line[i] + line[j]
        joltage = joltage_str.to_i
        if joltage > max_line_joltage
          max_line_joltage = joltage
        end
      end
    end
    total_joltage += max_line_joltage
  end

  puts "The total output joltage is: #{total_joltage}"
end

solve

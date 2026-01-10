def find_max_subsequence(line, k)
  return "0" * k if line.length < k || k == 0

  result_chars = []
  current_start_index = 0

  (0...k).each do |i|
    remaining_to_find = k - i
    search_end_index = line.length - remaining_to_find

    best_digit = -1
    best_digit_index = -1
    
    # Find the best digit in the current search window
    (current_start_index..search_end_index).each do |j|
      digit = line[j].to_i
      if digit > best_digit
        best_digit = digit
        best_digit_index = j
      end
    end
    
    result_chars << best_digit.to_s
    current_start_index = best_digit_index + 1
  end

  result_chars.join
end

def solve
  input_path = File.join(File.dirname(__FILE__), 'input.txt')
  lines = File.readlines(input_path)

  total_joltage = 0
  k = 12 # Number of digits to select

  lines.each do |line|
    line = line.strip
    next if line.empty?

    max_line_joltage_str = find_max_subsequence(line, k)
    total_joltage += max_line_joltage_str.to_i
  end

  puts "The total output joltage is: #{total_joltage}"
end

solve

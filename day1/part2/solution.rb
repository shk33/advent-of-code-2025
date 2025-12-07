
def solve(rotations)
  current_position = 50
  total_zero_count = 0

  rotations.each do |line|
    next if line.empty?

    direction = line[0]
    distance = line[1..-1].to_i

    zeros_this_turn = 0
    if direction == 'R'
      # 'div' in Ruby is floor division
      zeros_this_turn = (current_position + distance).div(100) - current_position.div(100)
    elsif direction == 'L'
      zeros_this_turn = (current_position - 1).div(100) - (current_position - distance - 1).div(100)
    end
    
    total_zero_count += zeros_this_turn

    if direction == 'R'
      current_position += distance
    elsif direction == 'L'
      current_position -= distance
    end
  end
  total_zero_count
end

def main
  input_file_path = File.join(__dir__, '../part1/input1.txt')
  rotations = File.readlines(input_file_path, chomp: true) # chomp: true removes newlines

  password = solve(rotations)
  puts "The password is: #{password}"
end

main if __FILE__ == $PROGRAM_NAME

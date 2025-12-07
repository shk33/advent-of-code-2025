
def solve(rotations)
  current_position = 50
  zero_count = 0

  rotations.each do |line|
    line = line.strip
    next if line.empty?

    direction = line[0]
    distance = line[1..-1].to_i

    if direction == 'R'
      current_position += distance
    elsif direction == 'L'
      current_position -= distance
    end
    
    current_position %= 100

    if current_position == 0
      zero_count += 1
    end
  end
  zero_count
end

def main
  input_file_path = File.join(__dir__, 'input1.txt')
  rotations = File.readlines(input_file_path).map(&:strip) # Read lines and strip whitespace

  password = solve(rotations)
  puts "The password is: #{password}"
end

# Call the main function when the script is executed
main if __FILE__ == $PROGRAM_NAME

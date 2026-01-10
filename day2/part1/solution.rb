
require 'set'

def parse_ranges(input_str)
  ranges = []
  max_id = 0
  input_str.strip.split(',').each do |r_str|
    next if r_str.empty?
    start_str, end_str = r_str.split('-')
    start_val = start_str.to_i
    end_val = end_str.to_i
    ranges.push([start_val, end_val])
    max_id = end_val if end_val > max_id
  end
  [ranges, max_id]
end

def generate_invalid_ids(max_id)
  invalid_ids = []
  base = 1
  s_base_prev = '0'
  loop do
    s_base = base.to_s
    s_invalid = s_base + s_base
    n_invalid = s_invalid.to_i

    if n_invalid > max_id
      break if s_base.length > s_base_prev.length
    else
      invalid_ids.push(n_invalid)
    end
    s_base_prev = s_base
    base += 1
  end
  invalid_ids
end

def solve(input_str)
  ranges, max_id = parse_ranges(input_str)
  potential_ids = generate_invalid_ids(max_id)
  
  found_invalid_ids = Set.new

  potential_ids.each do |invalid_id|
    ranges.each do |start_val, end_val|
      if invalid_id >= start_val && invalid_id <= end_val
        found_invalid_ids.add(invalid_id)
        break
      end
    end
  end
  
  found_invalid_ids.sum
end

def main
  input_file_path = File.join(__dir__, 'input.txt')
  input_str = File.read(input_file_path)
  
  result = solve(input_str)
  puts "The sum of all invalid IDs is: #{result}"
end

main

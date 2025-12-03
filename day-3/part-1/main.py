
example_input = """
987654321111111
811111111111119
234234234234278
818181911112111
"""

def main():
  # Read input from file
  with open('input.txt', 'r') as file:
    data = file.read().strip()
  
  lines = parse_input(data)

  total_joltage = 0

  for line in lines:
    total_joltage += solve(line)
    print(f"line: {solve(line)}")

  # Print the result
  print(f"Result: {total_joltage}")

def solve(line):
  largest_number = 0
  largest_number_idx = -1
  for char in line:
    num = int(char)
    if num > largest_number:
      if line.index(char) == len(line) - 1:
        break
      largest_number_idx = line.index(char)
      largest_number = num
  
  largest_number_positioned_after_the_other_largest_number = 0
  for idx in range(largest_number_idx + 1, len(line)):
    num = int(line[idx])
    if num > largest_number_positioned_after_the_other_largest_number:
      largest_number_positioned_after_the_other_largest_number = num

  return int(f"{largest_number}{largest_number_positioned_after_the_other_largest_number}")

def parse_input(data):
  lines = data.strip().split('\n')
  return lines

if __name__ == "__main__":
  main()
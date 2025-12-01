//@ts-ignore
import fs from "fs";

// L = down
// R = up

type instruction = `${"L" | "R"}${number}`;

type instructions = instruction[];

const START = 50;

let end_at_zero = 0;

const example_instructions: instructions = [
  "L68",
  "L30",
  "R48",
  "L5",
  "R60",
  "L55",
  "L1",
  "L99",
  "R14",
  "L82",
];

function count_zero_ends(instructions: instructions) {
  let position: number = START;

  for (const instruction of instructions) {
    const direction = instruction[0] === "L" ? -1 : 1;
    const distance = Number(instruction.slice(1));

    position = wrap_around(position + direction * distance);

    if (position === 0) end_at_zero += 1;
  }

  console.log("End at zero count:", end_at_zero);
}

function wrap_around(position: number): number {
  let wrapped_position = position;
  if (position < 0) {
    wrapped_position = 100 + position;
  } else if (position > 99) {
    wrapped_position = position - 100;
  }

  if (wrapped_position > 99 || wrapped_position < 0) {
    return wrap_around(wrapped_position);
  }

  return wrapped_position;
}

function extract_from_file(path: string): instructions {
  const data = fs.readFileSync(path, "utf-8");
  return data.trim().split("\n") as instructions;
}

const data = extract_from_file("input.txt");

count_zero_ends(data);

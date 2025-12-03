//@ts-ignore
import fs from "fs";

// L = down
// R = up

type instruction = `${"L" | "R"}${number}`;

type instructions = instruction[];

const START = 50;

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

function count_zero_passes(instructions: instructions) {
  let position: number = START;
  let zero_count = 0;

  for (const instruction of instructions) {
    const direction = instruction[0] === "L" ? -1 : 1;
    const distance = Number(instruction.slice(1));

    for (let i = 0; i < distance; i++) {
      position = (position + direction + 100) % 100;
      if (position === 0) {
        zero_count++;
      }
    }
  }

  console.log("Total zero passes:", zero_count);
  return zero_count;
}

function extract_from_file(path: string): instructions {
  const data = fs.readFileSync(path, "utf-8");
  return data.trim().split("\n") as instructions;
}

const data = extract_from_file("input.txt");

count_zero_passes(example_instructions);

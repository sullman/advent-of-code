const readline = require('node:readline');

function execute(program, registers, onlyOnce) {
  let instruction = 0;
  let output = [];

  const combo = operand => {
    switch (operand) {
    case 0:
    case 1:
    case 2:
    case 3:
      return operand;
    case 4:
      return registers.A;
    case 5:
      return registers.B;
    case 6:
      return registers.C;
    case 7:
    default:
      throw new Error('You lie!');
    }
  };

  while (true) {
    if (instruction >= program.length - 1) break;

    const opcode = program[instruction++];
    const operand = program[instruction++];

    switch (opcode) {
    case 0:
      registers.A = registers.A >> combo(operand);
      break;
    case 1:
      registers.B = registers.B ^ operand;
      break;
    case 2:
      registers.B = 0b111 & combo(operand);
      break;
    case 3:
      if (registers.A) {
        instruction = operand;
      }
      break;
    case 4:
      registers.B = registers.B ^ registers.C;
      break;
    case 5:
      output.push(0b111 & combo(operand));
      if (onlyOnce) return output;
      break;
    case 6:
      registers.B = registers.A >> combo(operand);
      break;
    case 7:
      registers.C = registers.A >> combo(operand);
      break;
    default:
      throw new Error('Huh?');
    }
  }

  return output;
}

function solveBackwards(program, a, index) {
  if (index < 0) return a;

  for (let b = 0; b < 8; b++) {
    // Can't use (a << 3) | b because JS bitwise operators only work on 32 bits!
    const [out] = execute(program, { A: (a * 8) + b, B: 0, C: 0}, true);
    if (out === program[index]) {
      const solution = solveBackwards(program, (a * 8) + b, index - 1);
      if (solution) return solution;
    }
  }

  return;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let registers = {};
  let program = [];

  for await (const line of rl) {
    if (!line) continue;
    const match = /Register ([ABC]): (\d+)/.exec(line);
    if (match) {
      registers[match[1]] = parseInt(match[2], 10);
    } else {
      program = line.slice('Program: '.length).split(',').map(Number);
    }
  }

  console.log('Part 1:', execute(program, registers).join(','));

  console.log('Part 2:', solveBackwards(program, 0, program.length - 1));
}

run().then(() => {
  process.exit();
});

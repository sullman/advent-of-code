const readline = require('node:readline');

const ADD = 0b001;
const MULTIPLY = 0b010;
const CONCAT = 0b100;

function isCorrect(target, value, operands, mask) {
  if (value > target) return 0;
  if (operands.length === 0) return target === value ? mask : 0;

  const [next, ...remaining] = operands;

  return isCorrect(target, value + next, remaining, mask | ADD)
    || isCorrect(target, value * next, remaining, mask | MULTIPLY)
    || isCorrect(target, Number(`${value}${next}`), remaining, mask | CONCAT);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let part1 = 0;
  let part2 = 0;

  for await (const line of rl) {
    if (!line) continue;
    const [targetStr, rest] = line.split(': ');
    const target = Number(targetStr);
    const operands = rest.split(' ').map(Number);
    const operators = isCorrect(target, operands.shift(), operands, 0);
    if (operators) {
      if ((operators & CONCAT) === 0) {
        part1 += target;
      }

      part2 += target;
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

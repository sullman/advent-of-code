const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  const lines = [];

  for await (const line of rl) {
    lines.push(line.trim().split(/\s+/));
  }

  const operators = lines.pop();

  let part1 = 0;

  for (let i = 0; i < operators.length; i++) {
    const result = lines.reduce((acc, values) => operators[i] === '*' ? acc * Number(values[i]) : acc + Number(values[i]), operators[i] === '*' ? 1 : 0);
    part1 += result;
    // console.log('Column', i, result);
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

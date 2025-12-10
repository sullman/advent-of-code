const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  const lines = [];
  const columns = [];

  for await (const line of rl) {
    lines.push(line.trim().split(/\s+/));
    if (columns.length === 0) {
      columns.length = line.length;
      columns.fill(0);
    }
    for (let i = 0; i < line.length; i++) {
      const num = Number.parseInt(line[i], 10);
      if (!isNaN(num)) {
        columns[i] = columns[i] * 10 + num;
      }
    }
  }

  const operators = lines.pop();

  let part1 = 0;
  let part2 = 0;

  for (let i = 0; i < operators.length; i++) {
    const op = operators[i];
    const result = lines.reduce((acc, values) => op === '*' ? acc * Number(values[i]) : acc + Number(values[i]), op === '*' ? 1 : 0);
    part1 += result;
    // console.log('Column', i, result);
  }

  while (operators.length) {
    const op = operators.shift();
    let value = op === '*' ? 1 : 0;
    let operand;
    while (operand = columns.shift()) {
      value = op === '*' ? value * operand : value + operand;
    }
    part2 += value;
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

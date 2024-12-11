const readline = require('node:readline');

const PATTERN = /mul\((\d+),(\d+)\)|(do(?:n't)?)\(\)/g;

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;
  let part2 = 0;
  let enabled = true;

  for await (const line of rl) {
    for (const match of line.matchAll(PATTERN)) {
      if (match[3]) {
        enabled = match[3] === 'do';
      } else {
        product = Number.parseInt(match[1], 10) * Number.parseInt(match[2], 10);
        part1 += product;
        if (enabled) part2 += product;
      }
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

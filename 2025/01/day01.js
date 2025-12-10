const readline = require('node:readline');

const PATTERN = /([LR])(\d+)/;
const SIZE = 100;

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let position = 50;
  let part1 = 0;
  let part2 = 0;

  for await (const line of rl) {
    const match = PATTERN.exec(line);
    if (!match) break;

    if (position === 0 && match[1] === 'L') part2--;
    position += (match[1] === 'L' ? -1 : 1) * parseInt(match[2], 10);
    if (position === 0) part2++;

    while (position >= SIZE) {
      position -= SIZE;
      part2++;
    }
    while (position < 0) {
      position += SIZE;
      part2++;
      if (position === 0) part2++;
    }
    if (position === 0) part1++;
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

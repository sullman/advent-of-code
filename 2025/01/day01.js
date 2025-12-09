const readline = require('node:readline');

const PATTERN = /([LR])(\d+)/;
const SIZE = 100;

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let position = 50;
  let zeroes = 0;

  for await (const line of rl) {
    const match = PATTERN.exec(line);
    if (!match) break;

    const delta = (match[1] === 'L' ? -1 : 1) * parseInt(match[2], 10);
    position = (position + delta) % SIZE;
    // if (position < 0) position += SIZE;
    if (position === 0) zeroes++;
  }

  console.log('Part 1:', zeroes);
}

run().then(() => {
  process.exit();
});

const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;
  const ranges = [];

  for await (const line of rl) {
    if (line.includes('-')) {
      ranges.push(line.split('-').map(Number));
    } else if (line) {
      const ingredient = Number.parseInt(line, 10);
      for (const [low, high] of ranges) {
        if (ingredient >= low && ingredient <= high) {
          part1++;
          break;
        }
      }
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

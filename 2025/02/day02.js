const readline = require('node:readline');

const PATTERN = /(\d+)-(\d+)/g;

function findInvalidInRange(low, high, numRepeats) {
  const invalid = new Set();
  let begin;
  let end;

  if (low.toString().length % numRepeats) {
    begin = 10 ** (Math.floor(low.toString().length / numRepeats));
  } else {
    begin = Number(low.toString().slice(0, Math.floor(low.toString().length / numRepeats)));
  }

  if (high.toString().length % numRepeats) {
    end = 10 ** (Math.floor(high.toString().length / numRepeats)) - 1;
  } else {
    end = Number(high.toString().slice(0, Math.floor(high.toString().length / numRepeats)));
  }

  for (let i = begin; i <= end; i++) {
    const candidate = Number(i.toString().repeat(numRepeats));
    if (candidate < low || candidate > high) continue;
    invalid.add(candidate);
  }

  return invalid;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = new Set();
  let part2 = new Set();

  for await (const line of rl) {
    for (const match of line.matchAll(PATTERN)) {
      const [_, low, high] = match.map(Number);
      part1 = part1.union(findInvalidInRange(low, high, 2));

      for (let i = 3; i <= high.toString().length; i++) {
        part2 = part2.union(findInvalidInRange(low, high, i));
      }
    }
  }

  part2 = part2.union(part1);

  console.log('Part 1:', [...part1].reduce((a, b) => a + b));
  console.log('Part 2:', [...part2].reduce((a, b) => a + b));
}

run().then(() => {
  process.exit();
});

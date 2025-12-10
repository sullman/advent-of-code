const readline = require('node:readline');

const PATTERN = /(\d+)-(\d+)/g;

function sumInvalidInRange(low, high) {
  let sum = 0;
  let begin;
  let end;

  if (low.toString().length % 2) {
    begin = 10 ** (low.toString().length >> 1);
  } else {
    begin = Math.floor(low / (10 ** (low.toString().length >> 1)));
  }

  if (high.toString().length % 2) {
    end = 10 ** (high.toString().length >> 1) - 1;
  } else {
    end = Math.floor(high / (10 ** (high.toString().length >> 1)))
  }

  for (let i = begin; i <= end; i++) {
    const candidate = Number(i.toString().repeat(2));
    if (candidate < low || candidate > high) continue;
    sum += candidate;
  }

  return sum;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;

  for await (const line of rl) {
    for (const match of line.matchAll(PATTERN)) {
      const [_, low, high] = match;
      part1 += sumInvalidInRange(parseInt(low, 10), parseInt(high, 10));
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

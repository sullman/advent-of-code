const readline = require('node:readline');

const PATTERN = /(\d+)\s+(\d+)/;

async function readInput() {
  const rl = readline.createInterface({ input: process.stdin });

  const left = [];
  const right = [];

  for await (const line of rl) {
    const match = PATTERN.exec(line);
    if (!match) break;

    left.push(parseInt(match[1], 10));
    right.push(parseInt(match[2], 10));
  }

  left.sort((a, b) => a - b);
  right.sort((a, b) => a - b);

  return [left, right];
}

function part1(left, right) {
  let sum = 0;

  for (let i = 0; i < left.length; i++) {
    sum += Math.abs(left[i] - right[i]);
  }

  console.log('Part 1:', sum);
}

function part2(left, right) {
  // This is a silly way to do this, just having a bit of fun...
  let sum = 0;
  let total = 0;

  for (let i = 0, j = 0; i < left.length && j < right.length; ) {
    if (left[i] === right[j]) {
      sum += left[i];
      j++;
    } else if (left[i] < right[j]) {
      const k = i;
      while (left[k] === left[i]) {
        i++;
        total += sum;
      }
      sum = 0;
    } else {
      j++;
    }
  }

  console.log('Part 2:', total);
}

async function run() {
  const [left, right] = await readInput();

  part1(left, right);
  part2(left, right);
}

run().then(() => {
  process.exit();
});

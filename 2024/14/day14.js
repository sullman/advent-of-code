const readline = require('node:readline');

const RE = /p=(\d+),(\d+) v=(-?\d+),(-?\d+)/;
const width = 101;
const height = 103;
const midWidth = width >> 1;
const midHeight = height >> 1;

function mod(dividend, divisor) {
  let remainder = dividend % divisor;
  if (remainder < 0) remainder += divisor;
  return remainder;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const quadrants = [0, 0, 0, 0];

  for await (const line of rl) {
    if (!line) continue;
    const [x, y, dx, dy] = RE.exec(line).slice(1, 5).map(Number);
    let finalX = mod(x + 100 * dx, width);
    let finalY = mod(y + 100 * dy, height);

    if (finalX === midWidth || finalY === midHeight) continue;
    let quadrant = 0;

    if (finalX > midWidth) {
      quadrant += 1;
    }
    if (finalY > midHeight) {
      quadrant += 2;
    }

    quadrants[quadrant]++;
  }

  const part1 = quadrants.reduce((prod, num) => prod * num, 1);

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

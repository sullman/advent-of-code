const readline = require('node:readline');

const NEIGHBORS = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
  [0, -1],
  [0, 1],
  [1, -1],
  [1, 0],
  [1, 1],
];

function removeRoll(rolls, neighbors, row, col) {
  if (!rolls.has(`${row},${col}`)) return 0;

  let numRemoved = 1;

  rolls.delete(`${row},${col}`);

  for (const [dr, dc] of NEIGHBORS) {
    const key = `${row + dr},${col + dc}`;
    const previousCount = neighbors.get(key);
    neighbors.set(key, previousCount - 1);
    if (previousCount === 4) {
      numRemoved += removeRoll(rolls, neighbors, row + dr, col + dc);
    }
  }

  return numRemoved;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;
  let row = 0;
  const neighbors = new Map();
  const rolls = new Set();

  for await (const line of rl) {
    for (let col = 0; col < line.length; col++) {
      if (line[col] === '@') {
        rolls.add(`${row},${col}`);
        for (const [dr, dc] of NEIGHBORS) {
          const key = `${row + dr},${col + dc}`;
          neighbors.set(key, 1 + (neighbors.get(key) ?? 0));
        }
      }
    }

    row++;
  }

  let part2 = 0;

  for (const key of rolls) {
    if ((neighbors.get(key) ?? 0) < 4) {
      part1++;
      const [r, c] = key.split(',').map(Number);
      part2 += removeRoll(rolls, neighbors, r, c);
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

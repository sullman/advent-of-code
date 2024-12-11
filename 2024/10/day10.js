const readline = require('node:readline');

const DIRECTIONS = [
  [-1, 0],
  [0, 1],
  [1, 0],
  [0, -1],
];

function findTrails(grid, row, col, peaks) {
  const current = grid[row][col];

  if (current === 9) {
    peaks.add(`${row},${col}`);
    return 1;
  }

  let sum = 0;

  for (const dir of DIRECTIONS) {
    const newRow = row + dir[0];
    const newCol = col + dir[1];
    if (grid[newRow]?.[newCol] === current + 1) {
      sum += findTrails(grid, newRow, newCol, peaks);
    }
  }

  return sum;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];

  for await (const line of rl) {
    if (!line) continue;
    grid.push(line.split('').map(Number));
  }

  let part1 = 0;
  let part2 = 0;

  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[row].length; col++) {
      if (grid[row][col] === 0) {
        const peaks = new Set();
        part2 += findTrails(grid, row, col, peaks);
        part1 += peaks.size;
      }
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

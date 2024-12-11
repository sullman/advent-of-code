const readline = require('node:readline');

const DIRECTIONS = [
  [-1, 0],
  [0, 1],
  [1, 0],
  [0, -1],
];

function isLoop(grid, row, col) {
  const visited = new Set();
  let dir = 0;

  while (true) {
    const key = `${row},${col},${dir}`;
    if (visited.has(key)) {
      return true;
    }

    visited.add(key);
    const newRow = row + DIRECTIONS[dir][0];
    const newCol = col + DIRECTIONS[dir][1];
    const next = grid[newRow]?.[newCol];

    if (next === '#') {
      dir = (dir + 1) % DIRECTIONS.length;
    } else if (next) {
      row = newRow;
      col = newCol;
    } else {
      break;
    }
  }

  return false;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];
  let startRow, startCol;

  for await (const line of rl) {
    if (line) {
      if (line.includes('^')) {
        startRow = grid.length;
        startCol = line.indexOf('^');
      }

      grid.push(line.split(''))
    }
  }

  const visited = new Set();
  let dir = 0;
  let row = startRow;
  let col = startCol;

  while (true) {
    visited.add(`${row},${col}`);
    const newRow = row + DIRECTIONS[dir][0];
    const newCol = col + DIRECTIONS[dir][1];
    const next = grid[newRow]?.[newCol];

    if (next === '#') {
      dir = (dir + 1) % DIRECTIONS.length;
    } else if (next) {
      row = newRow;
      col = newCol;
    } else {
      break;
    }
  }

  console.log('Part 1:', visited.size);

  visited.delete(`${startRow},${startCol}`);
  let numLoops = 0;

  for (const key of visited) {
    const candidate = key.split(',').map(Number);
    grid[candidate[0]][candidate[1]] = '#';

    if (isLoop(grid, startRow, startCol)) {
      numLoops++;
    }

    grid[candidate[0]][candidate[1]] = '.';
  }

  console.log('Part 2:', numLoops);
}

run().then(() => {
  process.exit();
});

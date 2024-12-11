const readline = require('node:readline');

const DIRECTIONS = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
  [0, -1],
  [0, 1],
  [1, -1],
  [1, 0],
  [1, 1],
];

async function readInput() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];

  for await (const line of rl) {
    grid.push(Array.from(line).map(ch => ch.codePointAt(0)));
  }

  return grid;
}

function isMatch(grid, word, row, col, dir) {
  for (const ch of word) {
    if (row < 0 || col < 0 || ch !== grid[row]?.[col]) return false;
    row += dir[0];
    col += dir[1];
  }

  return true;
}

const M = 'M'.codePointAt(0);
const A = 'A'.codePointAt(0);
const S = 'S'.codePointAt(0);

function isXMas(grid, row, col) {
  if (row < 1 || col < 1 || row > grid.length - 2 || grid[row][col] !== A) return false;

  const nw = grid[row - 1][col - 1];
  const ne = grid[row - 1][col + 1];
  const sw = grid[row + 1][col - 1];
  const se = grid[row + 1][col + 1];

  if (nw === M) {
    if (se !== S) return false;
  } else if (nw === S) {
    if (se !== M) return false;
  } else {
    return false;
  }

  if (ne === M) {
    return sw === S;
  } else if (ne === S) {
    return sw === M;
  }

  return false;
}

async function run() {
  const grid = await readInput();
  const word = Array.from('XMAS').map(ch => ch.codePointAt(0));
  let total1 = 0;
  let total2 = 0;

  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[row].length; col++) {
      for (const dir of DIRECTIONS) {
        if (isMatch(grid, word, row, col, dir)) {
          total1++;
        }
      }

      if (isXMas(grid, row, col)) {
        total2++;
      }
    }
  }

  console.log('Part 1:', total1);
  console.log('Part 2:', total2);
}

run().then(() => {
  process.exit();
});

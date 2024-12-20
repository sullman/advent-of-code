const readline = require('node:readline');

const DIRECTIONS = [
  [1, 0],
  [0, -1],
  [-1, 0],
  [0, 1],
];

const cache = new Map();
function honestDistance(grid, row, col) {
  const key = `${row},${col}`;
  if (cache.has(key)) return cache.get(key);

  let distance = 0;

  cache.set(key, distance);
  return distance;
}

function computeDistances(grid, endRow, endCol) {
  const distances = new Map();
  distances.set(`${endRow},${endCol}`, 0);
  const queue = [[endRow, endCol, 0]];

  while (queue.length) {
    const [row, col, distance] = queue.shift();
    for (const dir of DIRECTIONS) {
      const r = row + dir[0];
      const c = col + dir[1];
      const val = grid[r]?.[c];
      if (!distances.has(`${r},${c}`) && val && val !== '#') {
        distances.set(`${r},${c}`, distance + 1);
        queue.push([r, c, distance + 1]);
      }
    }
  }

  return distances;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];
  let startRow;
  let startCol;
  let endRow;
  let endCol;

  for await (const line of rl) {
    if (!line) continue;
    if (line.includes('S')) {
      startRow = grid.length;
      startCol = line.indexOf('S');
    }
    if (line.includes('E')) {
      endRow = grid.length;
      endCol = line.indexOf('E');
    }
    grid.push(line);
  }

  const distances = computeDistances(grid, endRow, endCol);

  const queue = [{
    row: startRow,
    col: startCol,
    distance: 0,
    cheatsRemaining: 1,
    visited: new Set()
  }];
  const solutions = [];
  const honest = distances.get(`${startRow},${startCol}`);

  while (queue.length) {
    const state = queue.shift();

    if (!state.cheatsRemaining) {
      solutions.push(state.distance + distances.get(`${state.row},${state.col}`));
      continue;
    }
    if (state.distance > honest - 100) break;
    if (grid[state.row][state.col] === 'E') {
      solutions.push(state.distance);
      break;
    }

    state.visited.add(`${state.row},${state.col}`);

    for (const dir of DIRECTIONS) {
      let row = state.row + dir[0];
      let col = state.col + dir[1];
      if (state.visited.has(`${row},${col}`)) continue;
      if (grid[row]?.[col] === '#' && state.cheatsRemaining) {
        row += dir[0];
        col += dir[1];
        if (grid[row]?.[col] && grid[row][col] !== '#') {
          queue.push({
            row,
            col,
            distance: state.distance + 2,
            cheatsRemaining: state.cheatsRemaining - 1,
            visited: new Set(state.visited)
          });
        }
      } else if (grid[row]?.[col]) {
        queue.push({
          row,
          col,
          distance: state.distance + 1,
          cheatsRemaining: state.cheatsRemaining,
          visited: new Set(state.visited)
        });
      }
    }
  }

  let part1 = 0;

  for (const speed of solutions) {
    if (honest - speed >= 100) part1++;
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

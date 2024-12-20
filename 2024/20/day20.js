const readline = require('node:readline');

const DIRECTIONS = [
  [1, 0],
  [0, -1],
  [-1, 0],
  [0, 1],
];

function computeDistances(grid, endRow, endCol) {
  const distances = new Array(grid.length);
  for (let i = 0; i < distances.length; i++) {
    distances[i] = new Array(grid[i].length);
  }
  distances[endRow][endCol] = 0;
  const queue = [[endRow, endCol, 0]];

  while (queue.length) {
    const [row, col, distance] = queue.shift();
    for (const dir of DIRECTIONS) {
      const r = row + dir[0];
      const c = col + dir[1];
      const val = grid[r]?.[c];
      if (distances[r][c] === undefined && val && val !== '#') {
        distances[r][c] = distance + 1;
        queue.push([r, c, distance + 1]);
      }
    }
  }

  return distances;
}

function findShortcuts(grid, distances, startRow, startCol, numCheats) {
  const honest = distances[startRow][startCol];
  const threshold = 100;
  let row = startRow;
  let col = startCol;
  let numShortcuts = 0;

  const isShortcut = (r, c, dr, dc) => {
    const newRow = r + dr;
    const newCol = c + dc;
    const dist = Math.abs(dr) + Math.abs(dc) + distances[newRow]?.[newCol];
    return dist <= distances[r][c] - threshold;
  }

  while (distances[row][col] > threshold) {
    for (let dr = 0; dr <= numCheats; dr++) {
      for (let dc = 0; dc <= numCheats - dr; dc++) {
        numShortcuts += isShortcut(row, col, dr, dc);
        if (dc) {
          numShortcuts += isShortcut(row, col, dr, -dc);
        }
        if (dr) {
          numShortcuts += isShortcut(row, col, -dr, dc);
          if (dc) {
            numShortcuts += isShortcut(row, col, -dr, -dc);
          }
        }
      }
    }

    for (const dir of DIRECTIONS) {
      if (distances[row][col] === distances[row + dir[0]]?.[col + dir[1]] + 1) {
        row = row + dir[0];
        col = col + dir[1];
        break;
      }
    }
  }

  return numShortcuts;
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

  console.log('Part 1:', findShortcuts(grid, distances, startRow, startCol, 2));
  console.log('Part 2:', findShortcuts(grid, distances, startRow, startCol, 20));
}

run().then(() => {
  process.exit();
});

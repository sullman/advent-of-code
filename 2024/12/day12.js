const readline = require('node:readline');

const DIRECTIONS = [
  [-1, 0],
  [0, 1],
  [1, 0],
  [0, -1],
];

function findPerimeter(grid, row, col, visited) {
  const key = `${row},${col}`;
  const crop = grid[row][col];
  let perimeter = 0;

  if (visited.has(key)) return 0;

  visited.add(key);

  for (const dir of DIRECTIONS) {
    if (crop === grid[row + dir[0]]?.[col + dir[1]]) {
      perimeter += findPerimeter(grid, row + dir[0], col + dir[1], visited);
    } else {
      perimeter++;
    }
  }

  return perimeter;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];

  for await (const line of rl) {
    if (!line) continue;
    grid.push(line);
  }

  let part1 = 0;
  const visited = new Set();

  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[row].length; col++) {
      if (visited.has(`${row},${col}`)) continue;

      const before = visited.size;
      const perimeter = findPerimeter(grid, row, col, visited);
      const area = visited.size - before;

      part1 += perimeter * area;
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

const readline = require('node:readline');

const DIRECTIONS = [
  [-1, 0],
  [0, 1],
  [1, 0],
  [0, -1],
];

function findRegion(grid, row, col, region, visited) {
  const key = `${row},${col}`;
  const crop = grid[row][col];

  if (visited.has(key)) return;

  visited.add(key);
  region.area++;

  for (let i = 0; i < DIRECTIONS.length; i++) {
    const dir = DIRECTIONS[i];
    if (crop === grid[row + dir[0]]?.[col + dir[1]]) {
      findRegion(grid, row + dir[0], col + dir[1], region, visited);
    } else {
      region.perimeter++;
      region.edges++;

      visited.add(`${crop},${row},${col},${i}`);
      for (const dir2 of DIRECTIONS) {
        if (visited.has(`${crop},${row + dir2[0]},${col + dir2[1]},${i}`)) {
          region.edges--;
        }
      }
    }
  }
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];

  for await (const line of rl) {
    if (!line) continue;
    grid.push(line);
  }

  let part1 = 0;
  let part2 = 0;
  const visited = new Set();

  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[row].length; col++) {
      if (visited.has(`${row},${col}`)) continue;

      const region = {
        area: 0,
        edges: 0,
        perimeter: 0,
      };
      findRegion(grid, row, col, region, visited);

      part1 += region.perimeter * region.area;
      part2 += region.edges * region.area;
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

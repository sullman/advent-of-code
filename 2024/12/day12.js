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

function findEdges(grid, row, col, visited) {
  const crop = grid[row][col];
  let edges = 0;

  if (visited.has(`${row},${col}`)) return 0;

  visited.add(`${row},${col}`);

  // Top edge
  if (crop === grid[row - 1]?.[col]) {
    edges += findEdges(grid, row - 1, col, visited);
  } else if (!visited.has(`${row},${col},top`)) {
    edges++;
    for (let c = col; grid[row][c] === crop && grid[row - 1]?.[c] !== crop; c++) {
      visited.add(`${row},${c},top`);
    }
    for (let c = col - 1; grid[row][c] === crop && grid[row - 1]?.[c] !== crop; c--) {
      visited.add(`${row},${c},top`);
    }
  }

  // Bottom edge
  if (crop === grid[row + 1]?.[col]) {
    edges += findEdges(grid, row + 1, col, visited);
  } else if (!visited.has(`${row},${col},bottom`)) {
    edges++;
    for (let c = col; grid[row][c] === crop && grid[row + 1]?.[c] !== crop; c++) {
      visited.add(`${row},${c},bottom`);
    }
    for (let c = col - 1; grid[row][c] === crop && grid[row + 1]?.[c] !== crop; c--) {
      visited.add(`${row},${c},bottom`);
    }
  }

  // Right edge
  if (crop === grid[row][col + 1]) {
    edges += findEdges(grid, row, col + 1, visited);
  } else if (!visited.has(`${row},${col},right`)) {
    edges++;
    for (let r = row; grid[r]?.[col] === crop && grid[r][col + 1] !== crop; r++) {
      visited.add(`${r},${col},right`);
    }
    for (let r = row - 1; grid[r]?.[col] === crop && grid[r][col + 1] !== crop; r--) {
      visited.add(`${r},${col},right`);
    }
  }

  // Left edge
  if (crop === grid[row][col - 1]) {
    edges += findEdges(grid, row, col - 1, visited);
  } else if (!visited.has(`${row},${col},left`)) {
    edges++;
    for (let r = row; grid[r]?.[col] === crop && grid[r][col - 1] !== crop; r++) {
      visited.add(`${r},${col},left`);
    }
    for (let r = row - 1; grid[r]?.[col] === crop && grid[r][col - 1] !== crop; r--) {
      visited.add(`${r},${col},left`);
    }
  }

  return edges;
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
  const visited2 = new Set();

  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[row].length; col++) {
      if (visited.has(`${row},${col}`)) continue;

      const before = visited.size;
      const perimeter = findPerimeter(grid, row, col, visited);
      const area = visited.size - before;
      const edges = findEdges(grid, row, col, visited2);

      part1 += perimeter * area;
      part2 += edges * area;
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

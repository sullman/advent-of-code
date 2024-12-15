const readline = require('node:readline');

const DIRECTIONS = {
  '^': [0, -1],
  '>': [1, 0],
  'v': [0, 1],
  '<': [-1, 0],
};

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];
  let robotX = 0;
  let robotY = 0;
  let moving = false;

  for await (const line of rl) {
    if (!line) {
      moving = true;
    } else if (!moving) {
      if (line.includes('@')) {
        robotY = grid.length;
        robotX = line.indexOf('@');
        grid.push(line.replace('@', '.').split(''));
      } else {
        grid.push(line.split(''));
      }
    } else {
      for (const ch of line) {
        const [dx, dy] = DIRECTIONS[ch];
        let x = robotX + dx;
        let y = robotY + dy;
        while (grid[y][x] === 'O') {
          x += dx;
          y += dy;
        }
        if (grid[y][x] === '#') continue;
        while (x !== robotX || y !== robotY) {
          grid[y][x] = grid[y - dy][x - dx];
          grid[y - dy][x - dx] = '.';
          x -= dx;
          y -= dy;
        }
        robotX += dx;
        robotY += dy;
      }
    }
  }

  let part1 = 0;
  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === 'O') {
        part1 += 100 * y + x;
      }
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

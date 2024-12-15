const readline = require('node:readline');

const DIRECTIONS = {
  '^': [0, -1],
  '>': [1, 0],
  'v': [0, 1],
  '<': [-1, 0],
};

function move(grid, robot, ch) {
  const [dx, dy] = DIRECTIONS[ch];
  let x = robot[0] + dx;
  let y = robot[1] + dy;
  while (grid[y][x] === 'O') {
    x += dx;
    y += dy;
  }
  if (grid[y][x] === '#') return false;
  while (x !== robot[0] || y !== robot[1]) {
    grid[y][x] = grid[y - dy][x - dx];
    grid[y - dy][x - dx] = '.';
    x -= dx;
    y -= dy;
  }
  robot[0] += dx;
  robot[1] += dy;

  return true;
}

function move2(grid, robot, ch) {
  const [dx, dy] = DIRECTIONS[ch];
  const spotsToCheck = [[robot[0] + dx, robot[1] + dy]];
  const spotsToMove = [];
  const checked = new Set();

  while (spotsToCheck.length) {
    const [x, y] = spotsToCheck.shift();
    if (checked.has(`${x},${y}`)) continue;
    checked.add(`${x},${y}`);
    if (grid[y][x] === '#') return false;
    spotsToMove.unshift([x - dx, y - dy]);
    if (grid[y][x] === '.') continue;
    spotsToCheck.push([x + dx, y + dy]);
    if (dy) {
      spotsToCheck.push([x + (grid[y][x] === '[' ? 1 : -1), y + dy]);
    }
  }

  for (const [x, y] of spotsToMove) {
    grid[y + dy][x + dx] = grid[y][x];
    grid[y][x] = '.';
  }

  robot[0] += dx;
  robot[1] += dy;

  return true;
}

function score(grid) {
  let total = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === 'O' || grid[y][x] === '[') {
        total += 100 * y + x;
      }
    }
  }

  return total;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];
  const grid2 = [];
  const robot = [0, 0];
  const robot2 = [0, 0];
  let moving = false;

  for await (const line of rl) {
    if (!line) {
      moving = true;
    } else if (!moving) {
      if (line.includes('@')) {
        robot[1] = grid.length;
        robot[0] = line.indexOf('@');
        grid.push(line.replace('@', '.').split(''));

        robot2[1] = grid2.length;
        robot2[0] = 2 * line.indexOf('@');
        grid2.push(line.replace('@', '.').split('').flatMap(ch => ch === 'O' ? ['[', ']'] : [ch, ch]));
      } else {
        grid.push(line.split(''));
        grid2.push(line.split('').flatMap(ch => ch === 'O' ? ['[', ']'] : [ch, ch]));
      }
    } else {
      for (const ch of line) {
        move(grid, robot, ch);
        move2(grid2, robot2, ch);
      }
    }
  }

  console.log('Part 1:', score(grid));
  console.log('Part 2:', score(grid2));
}

run().then(() => {
  process.exit();
});

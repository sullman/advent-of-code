const readline = require('node:readline');

const RE = /p=(\d+),(\d+) v=(-?\d+),(-?\d+)/;
const width = 101;
const height = 103;
const midWidth = width >> 1;
const midHeight = height >> 1;

const DIRECTIONS = [
  [1, 0],
  [0, -1],
  [-1, 0],
  [0, 1],
];

function isChristmasTree(robots) {
  const occupied = new Set();
  for (const robot of robots) {
    occupied.add(`${robot.x},${robot.y}`);
  }

  let withNeighbors = 0;
  for (const robot of robots) {
    let hasNeighbor = false;
    for (const [dx, dy] of DIRECTIONS) {
      if (occupied.has(`${robot.x + dx},${robot.y + dy}`)) {
        withNeighbors++;
        break;
      }
    }
  }

  if (withNeighbors < robots.length / 2) return false;

  // Seems like this is at least a _candidate_, let's look at it
  const grid = new Array(height);
  for (let row = 0; row < height; row++) {
    grid[row] = (new Array(width)).fill(' ');
  }

  for (const robot of robots) {
    grid[robot.y][robot.x] = '^';
  }

  for (const row of grid) {
    console.log(row.join(''));
  }

  return true;
}

function mod(dividend, divisor) {
  let remainder = dividend % divisor;
  if (remainder < 0) remainder += divisor;
  return remainder;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const quadrants = [0, 0, 0, 0];
  const robots = [];

  for await (const line of rl) {
    if (!line) continue;
    const [x, y, dx, dy] = RE.exec(line).slice(1, 5).map(Number);
    robots.push({ x, y, dx, dy });
    let finalX = mod(x + 100 * dx, width);
    let finalY = mod(y + 100 * dy, height);

    if (finalX === midWidth || finalY === midHeight) continue;
    let quadrant = 0;

    if (finalX > midWidth) {
      quadrant += 1;
    }
    if (finalY > midHeight) {
      quadrant += 2;
    }

    quadrants[quadrant]++;
  }

  const part1 = quadrants.reduce((prod, num) => prod * num, 1);

  console.log('Part 1:', part1);

  let seconds = 0;

  while (!isChristmasTree(robots)) {
    seconds++;
    for (const robot of robots) {
      robot.x = mod(robot.x + robot.dx, width);
      robot.y = mod(robot.y + robot.dy, height);
    }
  }

  console.log('Part 2:', seconds);
}

run().then(() => {
  process.exit();
});

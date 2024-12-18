const readline = require('node:readline');

const DIRECTIONS = [
  [1, 0],
  [0, -1],
  [-1, 0],
  [0, 1],
];

function insertSorted(queue, seen, state) {
  const key = `${state.x},${state.y}`;
  const prev = seen.get(key);
  if (prev <= state.distance) return;
  seen.set(key, state.distance);

  for (let i = 0; i < queue.length; i++) {
    if (state.bestPossible < queue[i].bestPossible) {
      queue.splice(i, 0, state);
      return;
    }
  }

  queue.push(state);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let maxX = 0;
  let maxY = 0;
  const bytes = [];

  for await (const line of rl) {
    if (!line) continue;
    const [x, y] = line.split(',').map(Number);
    bytes.push([x, y]);
    if (x > maxX) maxX = x;
    if (y > maxY) maxY = y;
  }

  const limit = maxX < 7 ? 12 : 1024;

  const grid = new Array(maxX + 1);
  for (let x = 0; x < grid.length; x++) {
    grid[x] = (new Array(maxY + 1)).fill(true);
  }

  for (let i = 0; i < limit; i++) {
    grid[bytes[i][0]][bytes[i][1]] = false;
  }

  const seen = new Map();
  const queue = [];
  insertSorted(queue, seen, {
    x: 0,
    y: 0,
    distance: 0,
    bestPossible: maxX + maxY
  });

  while (queue.length) {
    const state = queue.shift();
    if (state.x === maxX && state.y === maxY) {
      console.log('Part 1:', state.distance);
      break;
    }

    for (const [dx, dy] of DIRECTIONS) {
      if (grid[state.x + dx]?.[state.y + dy]) {
        const newX = state.x + dx;
        const newY = state.y + dy;
        insertSorted(queue, seen, {
          x: newX,
          y: newY,
          distance: state.distance + 1,
          bestPossible: state.distance + 1 + (maxX - newX) + (maxY - newY)
        });
      }
    }
  }

}

run().then(() => {
  process.exit();
});

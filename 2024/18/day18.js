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

function minPath(gridSize, bytes, numBytes) {
  const seen = new Map();
  const queue = [];
  insertSorted(queue, seen, {
    x: 0,
    y: 0,
    distance: 0,
    bestPossible: gridSize + gridSize
  });

  while (queue.length) {
    const state = queue.shift();
    if (state.x === gridSize && state.y === gridSize) {
      return state.distance;
    }

    for (const [dx, dy] of DIRECTIONS) {
      const newX = state.x + dx;
      const newY = state.y + dy;
      if (newX >= 0 && newX <= gridSize && newY >= 0 && newY <= gridSize && !(bytes.get(`${newX},${newY}`) <= numBytes)) {
        insertSorted(queue, seen, {
          x: newX,
          y: newY,
          distance: state.distance + 1,
          bestPossible: state.bestPossible + 1 - dx - dy
        });
      }
    }
  }
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let max = 0;
  const bytes = new Map();
  const labels = [];

  for await (const line of rl) {
    if (!line) continue;
    const [x, y] = line.split(',').map(Number);
    bytes.set(`${x},${y}`, bytes.size);
    labels.push(`${x},${y}`);
    if (x > max) max = x;
    if (y > max) max = y;
  }

  console.log('Part 1:', minPath(max, bytes, max < 7 ? 12 : 1024));

  let good = 0;
  let bad = bytes.size - 1;

  while (good < bad - 1) {
    const check = good + ((bad - good) >> 2);
    if (minPath(max, bytes, check)) {
      good = check;
    } else {
      bad = check;
    }
  }

  console.log('Part 2:', labels[bad]);
}

run().then(() => {
  process.exit();
});

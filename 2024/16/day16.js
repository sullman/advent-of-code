const readline = require('node:readline');

const TURN_SCORE = 1000;
const DIRECTIONS = [
  [0, 1],
  [-1, 0],
  [0, -1],
  [1, 0],
];

function addAll(set, values) {
  for (const val of values) {
    set.add(val);
  }
}

function insertSorted(queue, seen, state) {
  const key = `${state.row},${state.col},${state.dir}`;
  const prev = seen.get(key);
  if (prev?.score === state.score) {
    addAll(prev.path, state.path);
    return;
  }
  if (prev && prev.score < state.score) return;
  seen.set(key, state);

  for (let i = 0; i < queue.length; i++) {
    if (state.score < queue[i].score) {
      queue.splice(i, 0, state);
      return;
    }
  }

  queue.push(state);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const grid = [];
  let row = 0;
  let col = 0;
  let endRow = 0;
  let endCol = 0;

  for await (const line of rl) {
    if (!line) continue;
    if (line.includes('S')) {
      row = grid.length;
      col = line.indexOf('S');
    }
    if (line.includes('E')) {
      endRow = grid.length;
      endCol = line.indexOf('E');
    }
    grid.push(line);
  }

  const queue = [];
  const seen = new Map();
  let best;
  const allUsedTiles = new Set();

  queue.push({
    row,
    col,
    dir: 0,
    score: 0,
    path: new Set([`${row},${col}`]),
  });

  while (queue.length) {
    const state = queue.shift();
    if (best && state.score > best) break;

    if (grid[state.row][state.col] === 'E') {
      if (!best) {
        console.log('Part 1:', state.score);
        best = state.score;
      }

      addAll(allUsedTiles, state.path);

      continue;
    }

    const dir = DIRECTIONS[state.dir];
    if (grid[state.row + dir[0]][state.col + dir[1]] !== '#') {
      insertSorted(queue, seen, {
        row: state.row + dir[0],
        col: state.col + dir[1],
        dir: state.dir,
        score: state.score + 1,
        path: new Set([...state.path, `${state.row + dir[0]},${state.col + dir[1]}`])
      });
    }

    insertSorted(queue, seen, {
      ...state,
      dir: (state.dir + 1) % DIRECTIONS.length,
      score: state.score + TURN_SCORE
    });
    insertSorted(queue, seen, {
      ...state,
      dir: (state.dir + 3) % DIRECTIONS.length,
      score: state.score + TURN_SCORE
    });
  }

  console.log('Part 2:', allUsedTiles.size);
}

run().then(() => {
  process.exit();
});

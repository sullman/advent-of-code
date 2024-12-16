const readline = require('node:readline');

const TURN_SCORE = 1000;
const DIRECTIONS = [
  [0, 1],
  [-1, 0],
  [0, -1],
  [1, 0],
];

function insertSorted(queue, seen, state) {
  const key = `${state.row},${state.col},${state.dir}`;
  const prev = seen.get(key);
  if (prev && prev <= state.score) return;
  seen.set(key, state.score);

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

  queue.push({
    row,
    col,
    dir: 0,
    score: 0,
    // bestPossible: Math.abs(endRow - row) + Math.abs(endCol - col) + TURN_SCORE
  });

  while (queue.length) {
    const state = queue.shift();
    const dir = DIRECTIONS[state.dir];
    if (grid[state.row + dir[0]][state.col + dir[1]] === '.') {
      insertSorted(queue, seen, {
        row: state.row + dir[0],
        col: state.col + dir[1],
        dir: state.dir,
        score: state.score + 1
      });
    } else if (grid[state.row + dir[0]][state.col + dir[1]] === 'E') {
      console.log('Part 1:', state.score + 1);
      break;
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
}

run().then(() => {
  process.exit();
});

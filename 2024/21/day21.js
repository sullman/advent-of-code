const readline = require('node:readline');

const KEYPAD = {
  '7': [0, 0],
  '8': [0, 1],
  '9': [0, 2],
  '4': [1, 0],
  '5': [1, 1],
  '6': [1, 2],
  '1': [2, 0],
  '2': [2, 1],
  '3': [2, 2],
  '0': [3, 1],
  'A': [3, 2],
};

const DIRPAD = {
  '^': [0, 1],
  'A': [0, 2],
  '<': [1, 0],
  'v': [1, 1],
  '>': [1, 2],
};

const controlCache = new Map();
function generateControls(pad, over, next) {
  const key = `${over}${next}`;
  if (controlCache.has(key)) return controlCache.get(key);

  const controls = [];
  const [nextRow, nextCol] = pad[next];
  let [row, col] = pad[over];
  const canGoLeft = row !== pad['A'][0] || nextCol !== 0;
  const canGoDown = col !== 0 || nextRow !== pad['A'][0];
  for (; canGoLeft && col > nextCol; col--) controls.push('<');
  for (; canGoDown && row < nextRow; row++) controls.push('v');
  for (; canGoDown && row > nextRow; row--) controls.push('^');
  for (; col < nextCol; col++) controls.push('>');
  for (; !canGoDown && row > nextRow; row--) controls.push('^');
  for (; !canGoLeft && col > nextCol; col--) controls.push('<');
  for (; !canGoDown && row < nextRow; row++) controls.push('v');
  controls.push('A');

  controlCache.set(key, controls);
  return controls;
}

const lengthCache = new Map();
function countRequiredMoves(over, next, depth) {
  const key = `${over}${next}:${depth}`;
  if (lengthCache.has(key)) return lengthCache.get(key);

  let len = 0;
  const controls = ['A', ...generateControls(DIRPAD, over, next)];

  if (depth === 1) {
    len = controls.length - 1;
  } else {
    for (let i = 1; i < controls.length; i++) {
      len += countRequiredMoves(controls[i - 1], controls[i], depth - 1);
    }
  }

  lengthCache.set(key, len);
  return len;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let part1 = 0;
  let part2 = 0;

  for await (const line of rl) {
    if (!line) continue;
    let len1 = 0;
    let len2 = 0;
    const sequence = ['A', ...line.split('')];

    for (let i = 1; i < sequence.length; i++) {
      const controls = ['A', ...generateControls(KEYPAD, sequence[i - 1], sequence[i])];
      for (let j = 1; j < controls.length; j++) {
        len1 += countRequiredMoves(controls[j - 1], controls[j], 2);
        len2 += countRequiredMoves(controls[j - 1], controls[j], 25);
      }
    }

    part1 += parseInt(line, 10) * len1;
    part2 += parseInt(line, 10) * len2;
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

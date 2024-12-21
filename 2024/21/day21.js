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

function* generateMovements(pad, sequence) {
  let [row, col] = pad['A'];

  for (const ch of sequence) {
    const [nextRow, nextCol] = pad[ch];
    const canGoLeft = row !== pad['A'][0] || nextCol !== 0;
    const canGoDown = col !== 0 || nextRow !== pad['A'][0];
    for (; canGoLeft && col > nextCol; col--) yield '<';
    for (; canGoDown && row < nextRow; row++) yield 'v';
    for (; row > nextRow; row--) yield '^';
    for (; col < nextCol; col++) yield '>';
    for (; !canGoLeft && col > nextCol; col--) yield '<';
    for (; !canGoDown && row < nextRow; row++) yield 'v';
    yield 'A';
  }
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let part1 = 0;

  for await (const line of rl) {
    if (!line) continue;
    let iter = generateMovements(KEYPAD, line.split(''));
    for (let i = 0; i < 2; i++) {
      iter = generateMovements(DIRPAD, iter);
    }
    let len = 0;
    for (const ch of iter) len++;
    part1 += parseInt(line, 10) * len;
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

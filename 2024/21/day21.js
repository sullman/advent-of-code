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

// Go down before you go right
// Go left before you go up
// Go left before you go down
// Up and right don't matter

function generateMovements(pad, sequence) {
  const movements = [];
  let [row, col] = pad['A'];

  for (const ch of sequence) {
    const [nextRow, nextCol] = pad[ch];
    const canGoLeft = row !== pad['A'][0] || nextCol !== 0;
    const canGoDown = col !== 0 || nextRow !== pad['A'][0];
    for (; canGoLeft && col > nextCol; col--) movements.push('<');
    for (; canGoDown && row < nextRow; row++) movements.push('v');
    for (; row > nextRow; row--) movements.push('^');
    for (; col < nextCol; col++) movements.push('>');
    for (; !canGoLeft && col > nextCol; col--) movements.push('<');
    for (; !canGoDown && row < nextRow; row++) movements.push('v');
    movements.push('A');
  }

  return movements;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let part1 = 0;

  /*
  const test = [
    '^^<<A', '<<^^A', '<^^<A', '<^<^A', '^<<^A', '^<^<A',
    'vv<<A', '<<vvA', '<vv<A', '<v<vA', 'v<<vA', 'v<v<A',
    '^^>>A', '>>^^A', '>^^>A', '>^>^A', '^>>^A', '^>^>A',
    'vv>>A', '>>vvA', '>vv>A', '>v>vA', 'v>>vA', 'v>v>A',
  ];
  for (const t of test) {
    const movements = generateMovements(DIRPAD, t.split(''));
    console.log(t, movements.length, movements.join(''));
    const movements2 = generateMovements(DIRPAD, movements);
    console.log(t, movements.length, movements2.length);
  }
    */

  for await (const line of rl) {
    if (!line) continue;
    const robot1 = generateMovements(KEYPAD, line.split(''));
    // console.log(robot1.join(''));
    const robot2 = generateMovements(DIRPAD, robot1);
    // console.log(robot2.join(''));
    const me = generateMovements(DIRPAD, robot2);
    // console.log(me.join(''));
    console.log(parseInt(line, 10), me.length);
    part1 += parseInt(line, 10) * me.length;
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

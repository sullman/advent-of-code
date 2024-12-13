const readline = require('node:readline');

const BUTTON_RE = /Button [AB]: X\+(\d+), Y\+(\d+)/;
const PRIZE_RE = /Prize: X=(\d+), Y=(\d+)/;

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const buttons = [];
  let match;
  let part1 = 0;

  for await (const line of rl) {
    if (match = BUTTON_RE.exec(line)) {
      buttons.push(match.slice(1, 3).map(Number));
    } else if (match = PRIZE_RE.exec(line)) {
      const [prizeX, prizeY] = match.slice(1, 3).map(Number);
      const b = (prizeX*buttons[0][1] - prizeY*buttons[0][0]) / (buttons[1][0]*buttons[0][1] - buttons[1][1]*buttons[0][0]);
      if (Math.floor(b) === b) {
        const a = (prizeX - b*buttons[1][0]) / buttons[0][0];
        part1 += b + 3*a;
      } else {
        // console.log('Not solvable');
      }
    } else {
      buttons.length = 0;
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

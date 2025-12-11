const readline = require('node:readline');

const PATTERN = /\[([.#]+)\] (.*) \{([0-9,]+)\}/;

function findMinimumButtonPresses(target, buttons) {
  const mask = buttons.reduce((acc, button) => acc | button, 0);

  const queue = [{ state: 0, presses: 0 }];
  const seen = new Set();

  while (queue.length) {
    const { state, presses } = queue.shift();
    // if (state === target) return presses;
    for (const button of buttons) {
      const newState = (state ^ button) & mask;
      if (newState === target) return presses + 1;
      if (seen.has(newState)) continue;
      seen.add(newState);
      queue.push({ state: newState, presses: presses + 1 });
    }
  }

  throw new Error('No solution found');
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;

  for await (const line of rl) {
    const match = PATTERN.exec(line);
    if (!match) continue;
    const [_, indicators, buttonsStr, joltage] = match;

    // const mask = ((2 ** indicators.length) - 1);
    const target = Number.parseInt(indicators.split('').map(ch => ch === '#' ? '1' : '0').reverse().join(''), 2);
    const buttons = buttonsStr.split(' ').map(str => str.slice(1, -1).split(',').map(Number).reduce((acc, bit) => acc | (1 << bit), 0));

    part1 += findMinimumButtonPresses(target, buttons);
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

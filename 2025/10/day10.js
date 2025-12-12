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

function solveMatrix(matrix) {
  const m = matrix.length;
  const n = matrix[0].length;
  let h = 0;
  let k = 0;
  const maxPresses = Math.max(...matrix.map(row => row[row.length - 1]));

  while (h < m && k < n) {
    let pivot = h;
    for (let i = h + 1; i < m; i++) {
      if (Math.abs(matrix[i][k]) > Math.abs(matrix[pivot][k])) {
        pivot = i;
      }
    }

    if (matrix[pivot][k] === 0) {
      k++;
      continue;
    }

    const temp = matrix[h];
    matrix[h] = matrix[pivot];
    matrix[pivot] = temp;

    for (let i = h + 1; i < m; i++) {
      const f = matrix[i][k] / matrix[h][k];
      matrix[i][k] = 0;
      for (let j = k + 1; j < n; j++) {
        matrix[i][j] -= matrix[h][j] * f;
        if (Math.abs(matrix[i][j]) < 0.00000001) matrix[i][j] = 0;
      }
    }

    h++;
    k++;
  }

  for (let i = matrix.length - 1; i >= 0; i--) {
    if (!matrix[i].some(Boolean)) {
      matrix.pop();
    } else {
      break;
    }
  }

  if (matrix[0].length - 1 === matrix.length) {
    let sum = 0;
    for (let i = matrix.length - 1; i >= 0; i--) {
      const value = matrix[i][matrix[i].length - 1] / matrix[i][i];
      if (Math.round(value) < 0 || Math.abs(value - Math.round(value)) > 0.0000000000001) return -1;
      sum += value;
      for (let ii = i - 1; ii >= 0; ii--) {
        matrix[ii][matrix[ii].length - 1] -= matrix[ii][i] * value;
        matrix[ii][i] = 0;
      }
    }
    return sum;
  } else {
    let best = -1;
    let indexToFix = 0;
    for (; indexToFix < matrix.length; indexToFix++) {
      if (matrix[indexToFix][indexToFix] === 0) break;
    }

    for (let p = 0; p <= maxPresses; p++) {
      let clone = matrix.map(row => row.slice());
      const fixed = (new Array(matrix[0].length)).fill(0);
      fixed[fixed.length - 1] = p;
      fixed[indexToFix] = 1;
      clone.push(fixed);

      const result = solveMatrix(clone);
      if (result !== -1 && (best === -1 || result < best)) best = result;
    }
    return best;
  }

  return -1;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;
  let part2 = 0;

  for await (const line of rl) {
    const match = PATTERN.exec(line);
    if (!match) continue;
    const [_, indicators, buttonsStr, joltageStr] = match;

    // const mask = ((2 ** indicators.length) - 1);
    const target = Number.parseInt(indicators.split('').map(ch => ch === '#' ? '1' : '0').reverse().join(''), 2);
    const buttons = buttonsStr.split(' ').map(str => str.slice(1, -1).split(',').map(Number));
    const buttons1 = buttons.map(button => button.reduce((acc, bit) => acc | (1 << bit), 0));
    const joltage = joltageStr.split(',').map(Number);

    part1 += findMinimumButtonPresses(target, buttons1);

    const matrix = new Array(joltage.length);
    for (let i = 0; i < matrix.length; i++) {
      matrix[i] = (new Array(buttons.length + 1)).fill(0);
      matrix[i][buttons.length] = joltage[i];
    }

    for (let j = 0; j < buttons.length; j++) {
      for (const index of buttons[j]) {
        matrix[index][j] = 1;
      }
    }

    const result = solveMatrix(matrix);
    if (result === -1) throw new Error('No solution found');
    part2 += result;
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

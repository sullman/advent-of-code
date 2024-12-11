const readline = require('node:readline');

function isSafe(levels) {
  let safe = true;
  const min = levels[1] > levels[0] ? 1 : -3;
  const max = min + 2;

  for (let i = 1; safe && i < levels.length; i++) {
    const delta = levels[i] - levels[i - 1];
    safe = delta >= min && delta <= max;
  }

  return safe;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let safeCount1 = 0;
  let safeCount2 = 0;

  for await (const line of rl) {
    const levels = line.split(' ').map(Number).filter(Boolean);
    if (!levels.length) break;

    if (isSafe(levels)) {
      safeCount1++;
      safeCount2++;
    } else {
      for (let i = 0; i < levels.length; i++) {
        const alternate = [...levels];
        alternate.splice(i, 1);
        if (isSafe(alternate)) {
          safeCount2++;
          break;
        }
      }
    }
  }

  console.log('Part 1:', safeCount1);
  console.log('Part 2:', safeCount2);
}

run().then(() => {
  process.exit();
});

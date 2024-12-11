const readline = require('node:readline');

function blink(stone) {
  if (stone === '0') {
    return ['1'];
  } else if (stone.length % 2 === 0) {
    const str = stone.toString();
    const midpoint = str.length >> 1;
    return [str.slice(0, midpoint), parseInt(str.slice(midpoint), 10).toString()];
  } else {
    return [(Number(stone) * 2024).toString()];
  }
}

function numStones(counts) {
  let sum = 0;
  for (const count of counts.values()) {
    sum += count;
  }

  return sum;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let counts = new Map();

  for await (const line of rl) {
    if (!line) continue;
    for (const stone of line.split(' ')) {
      counts.set(stone, 1 + (counts.get(stone) ?? 0));
    }
  }

  for (let i = 0; i < 75; i++) {
    console.log(i);
    if (i === 25) {
      console.log('Part 1:', numStones(counts));
    }

    const newCounts = new Map();

    for (const [stone, count] of counts) {
      for (const newStone of blink(stone)) {
        newCounts.set(newStone, count + (newCounts.get(newStone) ?? 0));
      }
    }

    counts = newCounts;
  }

  console.log('Part 2:', numStones(counts));
}

run().then(() => {
  process.exit();
});

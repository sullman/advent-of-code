const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  const reds = [];
  let best = 0;

  for await (const line of rl) {
    const [x, y] = line.split(',').map(Number);

    for (const [a, b] of reds) {
      const area = (Math.abs(x - a) + 1) * (Math.abs(y - b) + 1);
      if (area > best) best = area;
    }

    reds.push([x, y]);
  }

  console.log('Part 1:', best);
}

run().then(() => {
  process.exit();
});

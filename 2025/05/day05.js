const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;
  const ranges = [];

  for await (const line of rl) {
    if (line.includes('-')) {
      ranges.push(line.split('-').map(Number));
    } else if (line) {
      const ingredient = Number.parseInt(line, 10);
      for (const [low, high] of ranges) {
        if (ingredient >= low && ingredient <= high) {
          part1++;
          break;
        }
      }
    }
  }

  const merged = [];
  for (const range of ranges) {
    const toMerge = [range];
    for (let i = merged.length - 1; i >= 0; i--) {
      const other = merged[i];
      if (other[1] < range[0] || other[0] > range[1]) continue;
      toMerge.push(other);
      merged.splice(i, 1);
    }
    merged.push([
      Math.min(...toMerge.map(r => r[0])),
      Math.max(...toMerge.map(r => r[1]))
    ]);
  }

  let part2 = 0;
  for (const [low, high] of merged) {
    part2 += (high - low) + 1;
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

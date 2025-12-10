const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let beams = [];
  let part1 = 0;

  for await (const line of rl) {
    if (beams.length === 0) {
      beams.length = line.length;
      beams.fill(0);
      beams[line.indexOf('S')] = 1;
      continue;
    }

    const newBeams = [...beams];
    for (let i = 0; i < line.length; i++) {
      if (beams[i] && line[i] === '^') {
        newBeams[i - 1] += beams[i];
        newBeams[i + 1] += beams[i];
        newBeams[i] = 0;
        part1++;
      }
    }
    beams = newBeams;
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', beams.reduce((acc, val) => acc + val, 0))
}

run().then(() => {
  process.exit();
});

const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let beams = [];
  let part1 = 0;

  for await (const line of rl) {
    if (beams.length === 0) {
      beams.length = line.length;
      beams.fill(false);
      beams[line.indexOf('S')] = true;
      continue;
    }

    const newBeams = [...beams];
    for (let i = 0; i < line.length; i++) {
      if (beams[i] && line[i] === '^') {
        newBeams[i - 1] = true;
        newBeams[i + 1] = true;
        newBeams[i] = false;
        part1++;
      }
    }
    beams = newBeams;
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

const readline = require('node:readline');

function optimizeRegex(designs) {
  designs.sort((a, b) => a.length - b.length);
  const necessary = [];
  let len = designs[0].length;
  let re = /^$/;

  for (const design of designs) {
    if (design.length > len) {
      len = design.length;
      re = new RegExp(`^(${necessary.join('|')})+$`);
    }

    if (!re.test(design)) necessary.push(design);
  }

  return new RegExp(`^(${necessary.join('|')})+$`);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let re;
  let part1 = 0;

  for await (const line of rl) {
    if (!line) continue;
    if (!re) {
      re = optimizeRegex(line.split(', '));
    } else {
      part1 += re.test(line);
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

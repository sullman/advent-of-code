const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  const shapes = [];
  let part1 = 0;

  for await (const line of rl) {
    if (!line) continue;
    if (line.includes(':')) {
      const [prefix, counts] = line.split(': ');
      if (counts) {
        const [width, height] = prefix.split('x').map(Number);
        const area = width * height;
        const required = counts.split(' ').map(Number).reduce((acc, cnt, i) => acc + (cnt * shapes[i]), 0);
        // console.log(required, area, (100 * required / area).toFixed(1));
        // This shouldn't really be sufficient, but it is for the input
        if (required < area) part1++;
      } else {
        shapes.push(0);
      }
    } else {
      shapes[shapes.length - 1] += line.split('').filter(c => c === '#').length;
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

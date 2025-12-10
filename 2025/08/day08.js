const readline = require('node:readline');

// const NUM_CONNECTIONS = 10;
const NUM_CONNECTIONS = 1000;

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  const boxes = [];
  const shortest = [{ distance: Number.MAX_SAFE_INTEGER }];
  const circuits = [];

  for await (const line of rl) {
    const [x, y, z] = line.split(',').map(Number);
    circuits.push(new Set([`${x},${y},${z}`]));

    for (const [a, b, c] of boxes) {
      const distance = Math.sqrt((x - a) ** 2 + (y - b) ** 2 + (z - c) ** 2);
      const insertIndex = shortest.findIndex(s => s.distance > distance);
      if (insertIndex !== -1) {
        shortest.splice(insertIndex, 0, { distance, from: `${x},${y},${z}`, to: `${a},${b},${c}` });
        // This is a bogus clamping, no guarantee that it will finish with those connections
        shortest.length = NUM_CONNECTIONS * 7;
      }
    }

    boxes.push([x, y, z]);
  }

  for (let i = 0; i < shortest.length; i++) {
    if (i === NUM_CONNECTIONS) {
      circuits.sort((a, b) => b.size - a.size);
      console.log('Part 1:', circuits[0].size * circuits[1].size * circuits[2].size);
    }

    const { from, to } = shortest[i];
    const fromIndex = circuits.findIndex(c => c.has(from));
    const toIndex = circuits.findIndex(c => c.has(to));
    if (fromIndex === toIndex) continue;
    if (circuits.length === 2) {
      console.log('Part 2:', Number.parseInt(from, 10) * Number.parseInt(to, 10));
      break;
    }

    circuits.push(circuits[fromIndex].union(circuits[toIndex]));
    circuits.splice(Math.max(fromIndex, toIndex), 1);
    circuits.splice(Math.min(fromIndex, toIndex), 1);
  }
}

run().then(() => {
  process.exit();
});

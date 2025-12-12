const readline = require('node:readline');

function countPaths(device, devices, seen, required, memo) {
  const key = `${device},${[...seen.intersection(required)].join(',')}`;
  if (key in memo) return memo[key];
  if (device === 'out') return required.isSubsetOf(seen) ? 1 : 0;
  let count = 0;
  for (const output of devices[device]) {
    if (seen.has(output)) continue;
    seen.add(output);
    count += countPaths(output, devices, seen, required, memo);
    seen.delete(output);
  }
  memo[key] = count;
  return count;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  const devices = {};

  for await (const line of rl) {
    if (!line) continue;
    const [device, str] = line.split(': ');
    const outputs = str.split(' ');
    devices[device] = outputs;
  }

  console.log('Part 1:', countPaths('you', devices, new Set(), new Set(), {}));
  console.log('Part 2:', countPaths('svr', devices, new Set(), new Set(['dac', 'fft']), {}));
}

run().then(() => {
  process.exit();
});

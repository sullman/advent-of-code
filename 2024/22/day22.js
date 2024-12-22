const readline = require('node:readline');

function next(secret) {
  secret = ((secret << 6n) ^ secret) & 0xffffffn;
  secret = ((secret >> 5n) ^ secret) & 0xffffffn;
  secret = ((secret << 11n) ^ secret) & 0xffffffn;

  return secret;
}

function value(secret) {
  return Number(secret % 10n);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let part1 = 0n;

  const earnings = new Map();

  for await (const line of rl) {
    if (!line) continue;
    const seen = new Set();
    let secret = BigInt(line);
    let prev = value(secret);
    const changes = [];
    for (let i = 0; i < 2000; i++) {
      secret = next(secret);
      const val = value(secret);
      changes.push(val - prev);
      prev = val;
      if (changes.length === 5) changes.shift();
      if (changes.length === 4) {
        const key = changes.join('');
        if (!seen.has(key)) {
          seen.add(key);
          earnings.set(key, (earnings.get(key) ?? 0) + val);
        }
      }
    }
    part1 += secret;
  }

  console.log('Part 1:', part1.toString());

  let part2 = 0;
  for (const val of earnings.values()) {
    if (val > part2) part2 = val;
  }

  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

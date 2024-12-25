const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  let isKey = false;
  let current;
  let height = 0;
  const keys = [];
  const locks = [];

  for await (const line of rl) {
    if (!current) {
      isKey = line[0] === '#';
      current = (new Array(line.length)).fill(0);
      height = 0;
    } else if (line) {
      height++;
      for (let i = 0; i < current.length; i++) {
        current[i] += line[i] === '#';
      }
    } else {
      (isKey ? keys : locks).push(current);
      current = undefined;
    }
  }

  if (current) {
    (isKey ? keys : locks).push(current);
  }

  let count = 0;
  for (const key of keys) {
    for (const lock of locks) {
      let fit = true;
      for (let i = 0; fit && i < key.length; i++) {
        if (key[i] + lock[i] > height) fit = false;
      }
      if (fit) count++;
    }
  }

  console.log('Part 1:', count);
}

run().then(() => {
  process.exit();
});

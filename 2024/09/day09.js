const readline = require('node:readline');

function part1(blocks) {
  let checksum = 0;
  let pos = 0;
  const queue = blocks.map(b => ({
    ...b,
    remaining: b.size
  }));

  if (!queue.at(-1).file) {
    queue.pop();
  }

  while (queue.length) {
    const block = queue[0];
    if (block.remaining === 0) {
      queue.shift();
      continue;
    }

    if (block.file) {
      checksum += (pos * block.id);
      block.remaining--;
    } else {
      const tail = queue.at(-1);
      checksum += (pos * tail.id);
      block.remaining--;
      tail.remaining--;
      if (tail.remaining === 0) {
        queue.pop();
        queue.pop();
      }
    }

    pos++;
  }

  console.log('Part 1:', checksum);
}

function part2(blocks) {
  for (let i = blocks.length - 1; i >= 0; i--) {
    const block = blocks[i];
    if (!block.file) continue;
    for (let j = 0; j < i; j++) {
      if (!blocks[j].file && block.size <= blocks[j].size) {
        blocks.splice(i, 1);
        blocks[j].size -= block.size;
        blocks.splice(j, 0, block);
        blocks[i].size += block.size;
        break;
      }
    }
  }

  let checksum = 0;
  for (let pos = 0, i = 0; i < blocks.length; i++) {
    for (let j = 0; j < blocks[i].size; j++, pos++) {
      if (blocks[i].file) {
        checksum += blocks[i].id * pos;
      }
    }
  }

  console.log('Part 2:', checksum);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const blocks = [];

  for await (const line of rl) {
    if (!line) continue;
    for (let i = 0; i < line.length; i++) {
      blocks.push({
        file: i % 2 === 0,
        id: i >> 1,
        size: parseInt(line[i], 10),
      });
    }
  }

  part1(blocks);
  part2(blocks);
}

run().then(() => {
  process.exit();
});

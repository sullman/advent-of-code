const readline = require('node:readline');

function addSegment(tree, segment) {
  let node = tree;
  for (const ch of segment) {
    if (!node.children.has(ch)) {
      node.children.set(ch, { children: new Map() });
    }

    node = node.children.get(ch);
  }

  node.valid = true;
}

function countSolutions(tree, cache, str) {
  if (!str) return 1;

  if (cache.has(str)) return cache.get(str);

  let solutions = 0;
  let node = tree;

  for (let i = 0; i < str.length; i++) {
    node = node.children.get(str[i]);
    if (!node) break;
    if (node.valid) {
      solutions += countSolutions(tree, cache, str.slice(i + 1));
    }
  }

  cache.set(str, solutions);

  return solutions;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const tree = { children: new Map() };
  const cache = new Map();
  let part1 = 0;
  let part2 = 0;

  for await (const line of rl) {
    if (!line) continue;
    if (tree.children.size === 0) {
      for (const segment of line.split(', ')) {
        addSegment(tree, segment);
      }
    } else {
      const num = countSolutions(tree, cache, line);
      if (num) {
        part1++;
        part2 += num;
      }
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

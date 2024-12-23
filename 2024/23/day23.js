const readline = require('node:readline');

function addConnection(nodes, connection) {
  const [a, b] = connection.split('-');

  if (!nodes.has(a)) {
    nodes.set(a, {
      name: a,
      neighbors: new Set()
    });
  }

  if (!nodes.has(b)) {
    nodes.set(b, {
      name: b,
      neighbors: new Set()
    });
  }

  nodes.get(a).neighbors.add(b);
  nodes.get(b).neighbors.add(a);
}

function findBiggestConnectedSet(nodes, group, candidates) {
  if (!candidates.length) return group;

  const [candidate, ...rest] = candidates;
  const node = nodes.get(candidate);
  let isConnected = true;

  for (const name of group) {
    if (!node.neighbors.has(name)) {
      isConnected = false;
      break;
    }
  }

  if (isConnected) {
    const inclusive = findBiggestConnectedSet(nodes, [...group, candidate], rest);
    const exclusive = findBiggestConnectedSet(nodes, group, rest);
    return inclusive.length > exclusive.length ? inclusive : exclusive;
  }

  return findBiggestConnectedSet(nodes, group, rest);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const nodes = new Map();
  let part1 = 0;

  for await (const line of rl) {
    if (!line) continue;
    addConnection(nodes, line);
  }

  for (const node of nodes.values()) {
    const { name } = node;
    const neighbors = [...node.neighbors];
    if (name[0] !== 't') continue;
    for (let i = 0; i < neighbors.length - 1; i++) {
      const a = neighbors[i];
      if (a[0] === 't' && a < name) continue;
      for (let j = i + 1; j < neighbors.length; j++) {
        const b = neighbors[j];
        if (b[0] === 't' && b < name) continue;
        if (nodes.get(a).neighbors.has(b)) part1++;
      }
    }
  }

  console.log('Part 1:', part1);

  // The biggest cycle will be made of nodes with lots of neighbors, check those first
  const ordered = [...nodes.values()].sort((a, b) => b.neighbors.size - a.neighbors.size);
  let best = [];

  for (const node of ordered) {
    if (node.neighbors.size <= best.length) break;
    const set = findBiggestConnectedSet(nodes, [node.name], [...node.neighbors]);
    if (set.length > best.length) best = set;
  }
  console.log('Part 2:', best.sort((a, b) => a.localeCompare(b)).join(','));
}

run().then(() => {
  process.exit();
});

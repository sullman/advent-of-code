const readline = require('node:readline');

function addConnection(nodes, connection) {
  const [a, b] = connection.split('-');

  if (!nodes.has(a)) {
    nodes.set(a, {
      name: a,
      neighbors: []
    });
  }

  if (!nodes.has(b)) {
    nodes.set(b, {
      name: b,
      neighbors: []
    });
  }

  nodes.get(a).neighbors.push(b);
  nodes.get(b).neighbors.push(a);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const nodes = new Map();
  let part1 = 0;

  for await (const line of rl) {
    if (!line) continue;
    addConnection(nodes, line);
  }

  for (const { name, neighbors } of nodes.values()) {
    if (name[0] !== 't') continue;
    for (let i = 0; i < neighbors.length - 1; i++) {
      const a = neighbors[i];
      if (a[0] === 't' && a < name) continue;
      for (let j = i + 1; j < neighbors.length; j++) {
        const b = neighbors[j];
        if (b[0] === 't' && b < name) continue;
        if (nodes.get(a).neighbors.includes(b)) part1++;
      }
    }
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

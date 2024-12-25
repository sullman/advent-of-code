const readline = require('node:readline');

function addConnection(nodes, a, b) {
  if (!nodes[a]) {
    nodes[a] = new Set();
  }

  if (!nodes[b]) {
    nodes[b] = new Set();
  }

  nodes[a].add(b);
  // nodes[b].add(a);
}

function tryKarger(byLabel, targetCuts) {
  let nodes = [];
  let edges = [];

  const namedNodes = {};
  for (const name of Object.keys(byLabel)) {
    nodes.push({ size: 1 });
    namedNodes[name] = nodes.at(-1);
  }

  for (const [name, connections] of Object.entries(byLabel)) {
    for (const other of connections) {
      edges.push([namedNodes[name], namedNodes[other]]);
    }
  }

  while (nodes.length > 2) {
    const [a, b] = edges[Math.floor(Math.random() * edges.length)];
    a.contracted = true;
    b.contracted = true;
    const c = { size: a.size + b.size };

    const newEdges = [];

    for (const edge of edges) {
      if (edge[0].contracted && edge[1].contracted) {
        // self, ignore
      } else if (edge[0].contracted) {
        newEdges.push([c, edge[1]]);
      } else if (edge[1].contracted) {
        newEdges.push([edge[0], c]);
      } else {
        newEdges.push(edge);
      }
    }

    edges = newEdges;
    nodes = nodes.filter(n => !n.contracted);
    nodes.push(c);
  }

  if (edges.length === targetCuts) {
    console.log(`Success! ${nodes[0].size} x ${nodes[1].size} = ${nodes[0].size * nodes[1].size}`);
    return true;
  }

  return false;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const nodes = {};

  for await (const line of rl) {
    if (!line) continue;
    const label = line.slice(0, 3);
    for (const other of line.slice(5).split(' ')) {
      addConnection(nodes, label, other);
    }
  }

  for (let i = 0; ; i++) {
    if (tryKarger(nodes, 3)) {
      console.log(`Took ${i + 1} tries`);
      break;
    }
  }
}

run().then(() => {
  process.exit();
});

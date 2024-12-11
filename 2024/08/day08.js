const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const lines = [];

  for await (const line of rl) {
    if (!line) continue;
    lines.push(line);
  }

  const numRows = lines.length;
  const numCols = lines[0].length;

  const antennae = new Map();
  const antinodes1 = new Set();
  const antinodes2 = new Set();

  const addIfInBounds = (r, c, set) => {
    if (r >= 0 && r < numRows && c >= 0 && c < numCols) {
      set.add(`${r},${c}`);
      return true;
    }

    return false;
  };

  const addAntinodes = ([row1, col1], [row2, col2]) => {
    const rowDelta = row2 - row1;
    const colDelta = col2 - col1;

    addIfInBounds(row1 - rowDelta, col1 - colDelta, antinodes1);
    addIfInBounds(row2 + rowDelta, col2 + colDelta, antinodes1);

    for (let r = row1, c = col1; addIfInBounds(r, c, antinodes2); r -= rowDelta, c -= colDelta) {}
    for (let r = row2, c = col2; addIfInBounds(r, c, antinodes2); r += rowDelta, c += colDelta) {}
  };

  for (let row = 0; row < lines.length; row++) {
    const line = lines[row];
    for (let col = 0; col < numCols; col++) {
      const antenna = line[col];
      if (antenna === '.') continue;

      if (antennae.has(antenna)) {
        const existing = antennae.get(antenna);

        for (const [r, c] of existing) {
          addAntinodes([row, col], [r, c]);
        }

        existing.push([row, col]);
      } else {
        antennae.set(antenna, [[row, col]]);
      }
    }
  }

  console.log('Part 1:', antinodes1.size);
  console.log('Part 2:', antinodes2.size);
}

run().then(() => {
  process.exit();
});

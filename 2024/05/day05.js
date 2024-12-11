const readline = require('node:readline');

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const after = {};
  const before = {};
  const updates = [];
  let inRules = true;

  for await (const line of rl) {
    if (!line) {
      if (inRules) {
        inRules = false;
        // TODO: Process transitivity?
      } else {
        break;
      }
    } else if (inRules) {
      const [a, b] = line.split('|');

      if (!after[b]) {
        after[b] = new Set();
      }

      if (!before[a]) {
        before[a] = new Set();
      }

      after[b].add(a);
      before[a].add(b);
    } else if (line) {
      updates.push(line);
    }
  }

  let part1 = 0;
  let part2 = 0;

  for (const line of updates) {
    const pages = line.split(',');
    let ordered = true;

    for (let i = 0; ordered && i < pages.length; i++) {
      const disallowed = after[pages[i]];
      for (let j = i + 1; j < pages.length; j++) {
        if (disallowed?.has(pages[j])) {
          ordered = false;
          break;
        }
      }
    }

    if (ordered) {
      part1 += parseInt(pages[pages.length >> 1], 10);
    } else {
      const fixed = [pages.shift()];

      while (pages.length) {
        const page = pages.shift();
        let inserted = false;

        for (let i = 0; i < fixed.length; i++) {
          if (before[page]?.has(fixed[i]) && (i === 0 || after[page]?.has(fixed[i - 1]))) {
            fixed.splice(i, 0, page);
            inserted = true;
            break;
          }
        }

        if (!inserted && after[page]?.has(fixed.at(-1))) {
          fixed.push(page);
        } else if (!inserted) {
          pages.push(page);
        }
      }

      part2 += parseInt(fixed[fixed.length >> 1], 10);
    }
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

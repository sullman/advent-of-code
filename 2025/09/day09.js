const readline = require('node:readline');

const INSIDE = 1;
const OUTSIDE = 2;

function isInside(x, y, grid, points) {
  const key = `${x},${y}`;
  if (grid[key]) return grid[key] === INSIDE;

  let inside = false;
  for (let i = 0; i < points.length; i++) {
    const [ax, ay] = points.at(i);
    const [bx, by] = points.at(i - 1);
    if ((ay > y) !== (by > y)) {
      const slope = (x - ax) * (by - ay) - (bx - ax) * (y - ay);
      if (slope === 0) {
        grid[key] = INSIDE;
        return true;
      }
      if ((slope < 0) !== (by < ay)) inside = !inside;
    }
  }

  grid[key] = inside ? INSIDE : OUTSIDE;
  return inside;
}

function isRectangleInside([ax, ay], [bx, by], grid, points) {
  // Check the borders first, only check the interior if the borders are promising
  const minX = Math.min(ax, bx);
  const maxX = Math.max(ax, bx);
  const minY = Math.min(ay, by);
  const maxY = Math.max(ay, by);

  let x = minX;
  let y = minY;

  for (; x <= maxX; x += 50) {
    if (!isInside(x, y, grid, points)) return false;
  }

  for (x = maxX; y <= maxY; y += 50) {
    if (!isInside(x, y, grid, points)) return false;
  }

  for (y = maxY; x >= minX; x -= 50) {
    if (!isInside(x, y, grid, points)) return false;
  }

  for (x = minX; y >= minY; y -= 50) {
    if (!isInside(x, y, grid, points)) return false;
  }

  // TODO: Check the interior

  return true;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  const corners = [];
  const grid = {};
  let best = 0;

  for await (const line of rl) {
    const [x, y] = line.split(',').map(Number);
    grid[`${x},${y}`] = INSIDE;

    if (corners.length) {
      const [prevX, prevY] = corners.at(-1);
      if (prevY === y) {
        const dx = prevX > x ? 1 : -1;
        for (let ax = x + dx; ax != prevX; ax += dx) {
          grid[`${ax},${y}`] = INSIDE;
        }
      } else {
        const dy = prevY > y ? 1 : -1;
        for (let ay = y + dy; ay != prevY; ay += dy) {
          grid[`${x},${ay}`] = INSIDE;
        }
      }
    }

    for (const [a, b] of corners) {
      const area = (Math.abs(x - a) + 1) * (Math.abs(y - b) + 1);
      if (area > best) best = area;
    }

    corners.push([x, y]);
  }

  console.log('Part 1:', best);

  best = 0;

  for (let i = 0; i < corners.length - 1; i++) {
    for (let j = i + 1; j < corners.length; j++) {
      const area = (Math.abs(corners[i][0] - corners[j][0]) + 1) * (Math.abs(corners[i][1] - corners[j][1]) + 1);
      if (area <= best) continue;
      if (isRectangleInside(corners[i], corners[j], grid, corners)) {
        console.log('Found', area);
        best = area;
      }
    }
  }

  console.log('Part 2:', best);
}

run().then(() => {
  process.exit();
});

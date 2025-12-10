const readline = require('node:readline');

function getJoltage(str) {
  const nums = str.split('').map(Number);
  let joltage = 0;
  let maxIndex = 0;
  let maxValue = -1;

  for (let i = 0; i < nums.length - 1; i++) {
    if (nums[i] > maxValue) {
      maxValue = nums[i];
      maxIndex = i;
    }
  }

  joltage = maxValue * 10;
  maxValue = -1;

  for (let i = maxIndex + 1; i < nums.length; i++) {
    if (nums[i] > maxValue) {
      maxValue = nums[i];
    }
  }

  joltage += maxValue;

  console.log(str, joltage);

  return joltage;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;

  for await (const line of rl) {
    part1 += getJoltage(line);
  }

  console.log('Part 1:', part1);
}

run().then(() => {
  process.exit();
});

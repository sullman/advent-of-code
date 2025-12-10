const readline = require('node:readline');

function getJoltage(str, numBatteries) {
  const nums = str.split('').map(Number);
  let joltage = 0;
  let maxIndex = -1;

  while (numBatteries > 0) {
    let maxValue = -1;
    joltage *= 10;

    for (let i = maxIndex + 1; i < nums.length + 1 - numBatteries; i++) {
      if (nums[i] > maxValue) {
        maxValue = nums[i];
        maxIndex = i;
      }
    }

    joltage += maxValue;
    numBatteries--;
  }

  // console.log(str, joltage);

  return joltage;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });

  let part1 = 0;
  let part2 = 0;

  for await (const line of rl) {
    part1 += getJoltage(line, 2);
    part2 += getJoltage(line, 12);
  }

  console.log('Part 1:', part1);
  console.log('Part 2:', part2);
}

run().then(() => {
  process.exit();
});

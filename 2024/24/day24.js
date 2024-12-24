const readline = require('node:readline');

const WIRE = /^(\w+): (0|1)$/;
const GATE = /^(\w+) (AND|OR|XOR) (\w+) -> (\w+)$/;

function solveForward(wires, name) {
  const wire = wires.get(name);
  if (typeof wire.value !== 'undefined') return wire.value;
  const input1 = solveForward(wires, wire.input[0]);
  const input2 = solveForward(wires, wire.input[1]);
  let val;

  switch (wire.gate) {
  case 'AND':
    val = input1 & input2;
    break;
  case 'OR':
    val = input1 | input2;
    break;
  case 'XOR':
    val = input1 ^ input2;
    break;
  default:
    throw new Error('Oops');
  }

  wire.value = val;
  return val;
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const wires = new Map();
  let match;

  for await (const line of rl) {
    if (!line) continue;
    if (match = WIRE.exec(line)) {
      wires.set(match[1], { value: Number(match[2]) });
    } else if (match = GATE.exec(line)) {
      const [, input1, gate, input2, output] = match;
      wires.set(output, { input: [input1, input2], gate });
    }
  }

  const z = [];
  for (let i = 0; ; i++) {
    const name = `z${i.toString().padStart(2, '0')}`;
    if (wires.has(name)) {
      z.unshift(solveForward(wires, name).toString());
    } else {
      break;
    }
  }

  console.log('Part 1:', parseInt(z.join(''), 2));
}

run().then(() => {
  process.exit();
});

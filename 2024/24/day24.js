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

const GATES = {
  AND: '&',
  OR: '|',
  XOR: '^',
};

function serialize(input1, input2, gate) {
  if (input1.localeCompare(input2) < 0) {
    return `${input1}${GATES[gate] || gate}${input2}`;
  }

  return `${input2}${GATES[gate] || gate}${input1}`;
}

function numberWire(letter, num) {
  return `${letter}${num.toString().padStart(2, '0')}`;
}

function swap(wires, byInput, a, b) {
  let wire = wires.get(a);
  byInput.set(serialize(wire.input[0], wire.input[1], wire.gate), b);
  wire = wires.get(b);
  byInput.set(serialize(wire.input[0], wire.input[1], wire.gate), a);
}

async function run() {
  const rl = readline.createInterface({ input: process.stdin });
  const wires = new Map();
  const byInput = new Map();
  let match;

  for await (const line of rl) {
    if (!line) continue;
    if (match = WIRE.exec(line)) {
      wires.set(match[1], { value: Number(match[2]) });
    } else if (match = GATE.exec(line)) {
      const [, input1, gate, input2, output] = match;
      wires.set(output, { input: [input1, input2], gate });
      byInput.set(serialize(input1, input2, gate), output);
    }
  }

  const z = [];
  for (let i = 0; ; i++) {
    const name = numberWire('z', i);
    if (wires.has(name)) {
      z.unshift(solveForward(wires, name).toString());
    } else {
      break;
    }
  }

  console.log('Part 1:', parseInt(z.join(''), 2));

  // Part 2
  let prev = [];
  const swapped = [];
  for (let i = 0; ; i++) {
    const name = numberWire('z', i);
    if (wires.has(name)) {
      const wire = wires.get(name);

      if (i >= 2) {
        const a = byInput.get(serialize(numberWire('x', i), numberWire('y', i), 'XOR'));
        if (!a) break;
        const b = byInput.get(serialize(numberWire('x', i - 1), numberWire('y', i - 1), 'AND'));
        const c = byInput.get(serialize(prev[0], prev[1], 'AND'));
        const d = byInput.get(serialize(b, c, 'OR'));
        const e = byInput.get(serialize(a, d, '^'));

        if (!e) {
          if (wire.input.includes(a)) {
            const other = wire.input.find(n => n !== a);
            swap(wires, byInput, d, other);
            swapped.push(d, other);
          } else {
            const other = wire.input.find(n => n !== d);
            swap(wires, byInput, a, other);
            swapped.push(a, other);
          }
          prev = wire.input;
        } else if (e !== name) {
          swapped.push(name, e);
          swap(wires, byInput, e, name);
          prev = [a, d];
        } else {
          prev = wire.input;
        }
      } else {
        prev = wire.input;
      }
    } else {
      break;
    }
  }

  if (swapped.length !== 8) throw new Error('Uh-oh');
  console.log('Part 2:', swapped.sort().join(','));
}

run().then(() => {
  process.exit();
});

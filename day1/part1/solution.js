const fs = require('fs');
const path = require('path');

function solve(input) {
    let currentPosition = 50;
    let zeroCount = 0;

    const rotations = input.split('\n').filter(line => line.trim() !== '');

    for (const rotation of rotations) {
        const direction = rotation[0];
        const distance = parseInt(rotation.substring(1), 10);

        if (direction === 'R') {
            currentPosition += distance;
        } else if (direction === 'L') {
            currentPosition -= distance;
        }

        // Handle the circular dial, ensuring the result is always positive
        currentPosition = ((currentPosition % 100) + 100) % 100;

        if (currentPosition === 0) {
            zeroCount++;
        }
    }

    return zeroCount;
}

function main() {
    const inputFilePath = path.join(__dirname, 'input1.txt');
    const input = fs.readFileSync(inputFilePath, 'utf8');
    const password = solve(input);
    console.log(`The password is: ${password}`);
}

if (require.main === module) {
    main();
}

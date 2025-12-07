const fs = require('fs');
const path = require('path');

function solve(input) {
    let currentPosition = 50;
    let totalZeroCount = 0;

    const rotations = input.split('\n').filter(line => line.trim() !== '');

    for (const rotation of rotations) {
        const direction = rotation[0];
        const distance = parseInt(rotation.substring(1), 10);

        let zerosThisTurn = 0;
        if (direction === 'R') {
            zerosThisTurn = Math.floor((currentPosition + distance) / 100) - Math.floor(currentPosition / 100);
        } else if (direction === 'L') {
            zerosThisTurn = Math.floor((currentPosition - 1) / 100) - Math.floor((currentPosition - distance - 1) / 100);
        }
        
        totalZeroCount += zerosThisTurn;

        // Update the unwrapped position for the next turn
        if (direction === 'R') {
            currentPosition += distance;
        } else if (direction === 'L') {
            currentPosition -= distance;
        }
    }

    return totalZeroCount;
}

function main() {
    const inputFilePath = path.join(__dirname, '../part1/input1.txt');
    const input = fs.readFileSync(inputFilePath, 'utf8');
    const password = solve(input);
    console.log(`The password is: ${password}`);
}

if (require.main === module) {
    main();
}

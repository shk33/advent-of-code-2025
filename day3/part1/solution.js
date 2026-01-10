const fs = require('fs');
const path = require('path');

function solve() {
    const inputPath = path.join(__dirname, 'input.txt');
    const lines = fs.readFileSync(inputPath, 'utf-8').split('\n');

    let totalJoltage = 0;
    for (const line of lines) {
        const trimmedLine = line.trim();
        if (!trimmedLine) {
            continue;
        }

        let maxLineJoltage = 0;
        for (let i = 0; i < trimmedLine.length; i++) {
            for (let j = i + 1; j < trimmedLine.length; j++) {
                const joltageStr = trimmedLine[i] + trimmedLine[j];
                const joltage = parseInt(joltageStr, 10);
                if (joltage > maxLineJoltage) {
                    maxLineJoltage = joltage;
                }
            }
        }
        totalJoltage += maxLineJoltage;
    }

    console.log(`The total output joltage is: ${totalJoltage}`);
}

solve();

const fs = require('fs');
const path = require('path');

function findMaxSubsequence(line, k) {
    if (line.length < k || k === 0) {
        return "0".repeat(k);
    }

    let resultChars = [];
    let currentStartIndex = 0;

    for (let i = 0; i < k; i++) {
        const remainingToFind = k - i;
        const searchEndIndex = line.length - remainingToFind;

        let bestDigit = -1;
        let bestDigitIndex = -1;

        for (let j = currentStartIndex; j <= searchEndIndex; j++) {
            const digit = parseInt(line[j], 10);
            if (digit > bestDigit) {
                bestDigit = digit;
                bestDigitIndex = j;
            }
        }
        
        resultChars.push(bestDigit.toString());
        currentStartIndex = bestDigitIndex + 1;
    }

    return resultChars.join('');
}


function solve() {
    const inputPath = path.join(__dirname, 'input.txt');
    const lines = fs.readFileSync(inputPath, 'utf-8').split('\n');

    let totalJoltage = 0n; // Use BigInt for total joltage
    const k = 12; // Number of digits to select

    for (const line of lines) {
        const trimmedLine = line.trim();
        if (!trimmedLine) {
            continue;
        }

        const maxLineJoltageStr = findMaxSubsequence(trimmedLine, k);
        totalJoltage += BigInt(maxLineJoltageStr);
    }

    console.log(`The total output joltage is: ${totalJoltage}`);
}

solve();

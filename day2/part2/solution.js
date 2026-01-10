
const fs = require('fs');
const path = require('path');

function parseRanges(inputStr) {
    const ranges = [];
    let maxId = BigInt(0); // Use BigInt for maxId as well
    const rangeStrs = inputStr.trim().split(',');
    for (const rStr of rangeStrs) {
        if (!rStr) continue;
        const [startStr, endStr] = rStr.split('-');
        const start = BigInt(startStr);
        const end = BigInt(endStr);
        ranges.push({ start, end });
        if (end > maxId) {
            maxId = end;
        }
    }
    return { ranges, maxId };
}

function generateInvalidIds(maxId) {
    const invalidIds = [];
    let baseNum = 1;

    while (true) {
        const baseS = baseNum.toString();
        
        // Optimization: if the smallest possible repetition (baseS + baseS) is too big, stop.
        const firstRepetitionVal = BigInt(baseS + baseS);
        if (firstRepetitionVal > maxId) {
            break;
        }
        
        let currentRepeatedS = baseS;
        
        while (true) {
            currentRepeatedS += baseS; // Append baseS again
            const nInvalid = BigInt(currentRepeatedS);

            if (nInvalid > maxId) {
                break; // This baseS repetition is too long
            }
            
            invalidIds.push(nInvalid);
        }
        
        baseNum++;
    }
    return invalidIds;
}

function solve(inputStr) {
    const { ranges, maxId } = parseRanges(inputStr);
    const potentialIds = generateInvalidIds(maxId);
    
    const foundInvalidIds = new Set();

    for (const invalidId of potentialIds) {
        for (const range of ranges) {
            if (range.start <= invalidId && invalidId <= range.end) {
                foundInvalidIds.add(invalidId);
                break;
            }
        }
    }

    let totalSum = BigInt(0);
    foundInvalidIds.forEach(id => {
        totalSum += id;
    });

    return totalSum.toString();
}

function main() {
    const inputFilePath = path.join(__dirname, 'input.txt');
    const inputStr = fs.readFileSync(inputFilePath, 'utf8');
    
    const result = solve(inputStr);
    console.log(`The sum of all invalid IDs is: ${result}`);
}

main();

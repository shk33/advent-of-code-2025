
const fs = require('fs');
const path = require('path');

function parseRanges(inputStr) {
    const ranges = [];
    let maxId = 0;
    const rangeStrs = inputStr.trim().split(',');
    for (const rStr of rangeStrs) {
        if (!rStr) continue;
        const [startStr, endStr] = rStr.split('-');
        const start = parseInt(startStr, 10);
        const end = parseInt(endStr, 10);
        ranges.push({ start, end });
        if (end > maxId) {
            maxId = end;
        }
    }
    return { ranges, maxId };
}

function generateInvalidIds(maxId) {
    const invalidIds = [];
    let base = 1;
    let sBasePrev = '0';
    while (true) {
        const sBase = base.toString();
        const sInvalid = sBase + sBase;
        // Use BigInt for potentially large numbers
        const nInvalid = BigInt(sInvalid);

        if (nInvalid > maxId) {
            if (sBase.length > sBasePrev.length) {
                break;
            }
        } else {
            invalidIds.push(nInvalid);
        }
        sBasePrev = sBase;
        base++;
    }
    return invalidIds;
}

function solve(inputStr) {
    const { ranges, maxId } = parseRanges(inputStr);
    // Convert maxId to BigInt for comparison
    const potentialIds = generateInvalidIds(BigInt(maxId));
    
    const foundInvalidIds = new Set();

    for (const invalidId of potentialIds) {
        for (const range of ranges) {
            // Convert range start/end to BigInt for comparison
            if (BigInt(range.start) <= invalidId && invalidId <= BigInt(range.end)) {
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

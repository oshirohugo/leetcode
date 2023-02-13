
function checkZeroOnes(s: string): boolean {
    let maxZeroCount = -1;
    let zeroCount = 0;
    let oneCount = 0;
    let maxOneCount = -1;
    for (let c of s) {
        if (c === '0') {
            maxOneCount = Math.max(oneCount, maxOneCount);
            oneCount = 0;
            zeroCount++;
        } else {
            maxZeroCount = Math.max(zeroCount, maxZeroCount);
            zeroCount = 0;
            oneCount++;
        }
    }

    if (zeroCount !== 0) {
        maxZeroCount = Math.max(zeroCount, maxZeroCount);
    } else {
        maxOneCount = Math.max(oneCount, maxOneCount);
    }

    return maxOneCount > maxZeroCount;
};

function countSegments(s: string): number {
    let countingSegment = false;
    let count = 0;
    for (let c of s) {
        if (c === " ") {
            if (countingSegment) {
                countingSegment = false;
                count++;
            }
        } else {
            countingSegment = true;
        }
    }

    if (countingSegment) {
        count++;
    }

    return count;
};

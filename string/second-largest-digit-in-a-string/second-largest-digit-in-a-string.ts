function secondHighest(s: string): number {
    let first = -1;
    let second = -1;

    for (let c of s) {
        const num = parseInt(c);
        if (Number.isNaN(num)) continue

        if (num > first) {
            second = first;
            first = num;
        } else if (num > second && num < first) {
            second = num;
        }
    }

    return second;
};

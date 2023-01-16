function balancedStringSplit(s: string): number {
    let res = 0;
    let count = 0;

    for (let c of s) {
        count += (c === 'L' ? 1 : -1)
        if (count === 0) {
            res++;
        }
    }

    return res;
};

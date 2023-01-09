function hasGroupsSizeX(deck: number[]): boolean {
    const counter = deck.reduce((acc, num) => {
        if (acc[num]) {
            acc[num]++;
            return acc;
        }
        return {
            ...acc,
            [num]: 1,
        }
    }, {})

    let min = Number.MAX_SAFE_INTEGER;
    for (let num of Object.keys(counter)) {
        const curr = counter[num];
        min = Math.min(min, curr);
    }

    if (min < 2) {
      return false;
    }
    
    for (let num of Object.keys(counter)) {
        const curr = counter[num];
        if (curr % min !== 0) {
            return false;
        }
    }

    return true;
};

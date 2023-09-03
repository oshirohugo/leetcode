function detectCapitalUse(word: string): boolean {
    let nOfCaps = 0;

    for (let c of word) {
        if (c === c.toUpperCase()) {
            nOfCaps++;
        }
    }

    return (nOfCaps === 0) ||
        (word[0] === word[0].toLocaleUpperCase() && nOfCaps === 1) ||
        (nOfCaps === word.length)
};

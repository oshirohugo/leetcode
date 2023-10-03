function checkAlmostEquivalent(word1: string, word2: string): boolean {
    const frequency = new Array<number>(26).fill(0);

    for (let c of word1) {
        frequency[c.charCodeAt(0) - 'a'.charCodeAt(0)]++
    }
    for (let c of word2) {
        frequency[c.charCodeAt(0) - 'a'.charCodeAt(0)]--
    }

    for (let f of frequency) {
        if (f > 3 || f < -3) {
            return false;
        }
    }

    return true;
};

function makeEqual(words: string[]): boolean {
    const count = new Array<number>(26).fill(0);

    for (let word of words) {
        for (let char of word) {
            const index = char.charCodeAt(0) - "a".charCodeAt(0);
            count[index]++;
        }
    }

    for (let el of count) {
        if ((el !== 0) && (el % words.length !== 0)) {
            return false;
        }
    }

    return true;
};

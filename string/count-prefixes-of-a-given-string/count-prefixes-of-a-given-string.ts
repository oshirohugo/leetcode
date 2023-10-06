function countPrefixes(words: string[], s: string): number {
    let count = 0;
    for (let word of words) {
        if (s.indexOf(word) === 0) {
            count++;
        }
    }

    return count;
};

function vowelStrings(words: string[], left: number, right: number): number {
    const vowels = new Set(['a', 'e', 'i', 'o', 'u']);
    let count = 0;
    for (let i = left; i <= right; i++) {
        if (vowels.has(words[i][0]) && vowels.has(words[i][words[i].length - 1])) {
            count++;
        }
    }

    return count;
};

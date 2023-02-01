function numOfStrings(patterns: string[], word: string): number {
    return patterns.reduce((acc, pattern) =>
        word.indexOf(pattern) != -1 ? acc + 1 : acc, 0);
};

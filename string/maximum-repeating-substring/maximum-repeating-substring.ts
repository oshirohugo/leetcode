function maxRepeating(sequence: string, word: string): number {
    let k = 0;
    let acc = word;
    while (sequence.indexOf(acc) !== -1) {
        k++;
        acc += word;
    }
    return k;
};

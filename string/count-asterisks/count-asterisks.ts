function countAsterisks(s: string): number {
    let asterisks = 0;
    let counting = true;

    for (let c of s) {
        if (c === '|') {
            if (counting) {
                counting = false;
                continue;
            }
            counting = true;
        }

        if (c === '*' && counting) {
            asterisks++;
        }
    }

    return asterisks;
};

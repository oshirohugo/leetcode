function digitCount(num: string): boolean {
    const count = Array<number>(10).fill(0)

    for (let c of num) {
        count[parseInt(c)]++;
    }

    for (let i = 0; i < num.length; i++) {
        if (count[i] !== parseInt(num[i])) {
            return false;
        }
    }

    return true;
};
// autobiographical numbers, introduced by Da Vinci and their is only certain set of numbers which satisfy the condition.

function minDeletionSize(strs: string[]): number {
    const nOfCols = strs[0].length;
    let del = 0;

    for (let col = 0; col < nOfCols; col++) {
        for (let row = 0; row < strs.length - 1; row++) {
            if (strs[row][col].charCodeAt(0) > strs[row + 1][col].charCodeAt(0)) {
                del++;
                break;
            }
        }
    }

    return del;
};

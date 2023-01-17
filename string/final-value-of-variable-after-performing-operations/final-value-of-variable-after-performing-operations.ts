function finalValueAfterOperations(operations: string[]): number {
    return operations.reduce((x, op) => op[1] === '+' ? x + 1 : x - 1, 0);
};

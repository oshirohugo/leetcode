function generateTheString(n: number): string {
    return n % 2 ? new Array<string>(n).fill('a').join('') : new Array<string>(n - 1).fill('a').join('') + 'b';
};

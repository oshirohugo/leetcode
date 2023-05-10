function generateTheString(n: number): string {
    const num = n / 2;
    const left = new Array<string>(num -1).fill('a');
    const right = new Array<string>(num +1).fill('b');

    return left.concat(right).join('');
};

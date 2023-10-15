function numberOfLines(widths: number[], s: string): number[] {
    let buffer = 100;
    let lines = 0;

    for (let i = 0; i < s.length;) {
        const attempt = buffer - widths[s[i].charCodeAt(0) - 'a'.charCodeAt(0)];

        if (attempt < 0) {
            lines++;
            buffer = 100;
        } else {
            buffer = attempt;
            i++;
        }
    }

    if (buffer < 100) {
        lines++;
    }

    return [lines, 100 - buffer];
};

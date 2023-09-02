function maximumValue(strs: string[]): number {
    let max = -1;
    for (let str of strs) {
        let value = Number(str)
        if (isNaN(value)) {
            value = str.length;
        }
        if (value > max) {
            max = value;
        }
    }

    return max;
};

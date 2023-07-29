function removeTrailingZeros(num: string): string {
    let i = num.length - 1;
    while (i) {
        if (num[i] !== "0") {
            break;
        }
        i--;
    }

    return num.slice(0, i + 1)
};

function removeTrailingZerosEasy(num: string): string {
    return num.replace(/0+$/, "")
}

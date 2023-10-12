function areNumbersAscending(s: string): boolean {
    const tokens = s.split(' ');
    let last = -1;

    for (let token of tokens) {
        let number = parseInt(token)
        if (isNaN(number)) {
            continue;
        }
        if (number <= last) {
            return false;
        }
        last = number;
    }

    return true;
};

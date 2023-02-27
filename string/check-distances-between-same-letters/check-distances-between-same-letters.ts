const a_CODE = "a".charCodeAt(0);

function checkDistances(s: string, distance: number[]): boolean {
    const lookup: { [key: string]: number } = {}

    let i = 0;
    for (let c of s) {
        if (lookup[c] !== undefined) {
            if ((i - lookup[c] - 1) !== distance[c.charCodeAt(0) - a_CODE]) {
                return false;
            }
        } else {
            lookup[c] = i;
        }
        i++;
    }

    return true;
};

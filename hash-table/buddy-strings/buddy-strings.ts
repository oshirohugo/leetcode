function buddyStrings(s: string, goal: string): boolean {
    if (s.length !== goal.length) {
        return false;
    }

    let diffs: number[] = [];

    for (let i = 0; i < s.length; i++) {
        if (s[i] !== goal[i]) {
            diffs.push(i);
            if (diffs.length > 2) {
                return false;
            }
        }
    }

    if (diffs.length === 0 || diffs.length === 2) {
        return s[diffs[0]] === goal[diffs[1]] && s[diffs[1]] === goal[diffs[0]];
    }

    return false;
};


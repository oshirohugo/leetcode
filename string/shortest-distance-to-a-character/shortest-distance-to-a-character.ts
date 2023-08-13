function shortestToChar(s: string, c: string): number[] {
    const n = s.length;
    let pos = -n;
    const dist = new Array<number>(n)

    for (let i = 0; i < n; i++) {
        if (s[i] === c) pos = i;
        dist[i] = i - pos;
    }

    for (let i = pos - 1; i >= 0; i--) {
        if (s[i] === c) pos = i;
        dist[i] = Math.min(dist[i], pos - i);
    }

    return dist
};

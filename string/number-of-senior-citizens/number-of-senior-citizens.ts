function countSeniors(details: string[]): number {
    let total = 0;

    for (let detail of details) {
        if (Number.parseInt(detail.slice(11, 13), 10) > 60) {
            total++;
        }
    }

    return total;
};

function countPoints(rings: string): number {
    const lookup: { [rod: number]: { [color: string]: boolean } } = {};

    for (let i = 0; i < rings.length - 1; i += 2) {
        const color = rings[i];
        const rod = parseInt(rings[i + 1]);
        lookup[rod] = {
            ...lookup[rod],
            [color]: true
        };
    }

    let fullRod = 0;
    for (let rod = 0; rod <= 9; rod++) {
        let counter = 0;
        for (let color of 'RGB') {
            if (lookup[rod] && lookup[rod][color]) {
                counter++;
            }
        }
        if (counter === 3) {
            fullRod++;
        }
    }

    return fullRod;
};

function sortPeople(names: string[], heights: number[]): string[] {
    const lookup : {[key: number]: string} = {};

    for (let i = 0; i < names.length; i++) {
        lookup[heights[i]] = names[i];
    }

    heights.sort((a, b) => b - a);

    for (let i = 0; i < names.length; i++) {
        names[i] = lookup[heights[i]];
    }

    return names;
};

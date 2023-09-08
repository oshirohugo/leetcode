function destCity(paths: string[][]): string {
    const cities: { [key: string]: number } = {};

    for (let path of paths) {
        cities[path[0]] = (cities[path[0]] || 0) + 1;
        cities[path[1]] = (cities[path[1]] || 0) - 1;
    }

    for (let city in cities) {
        if (cities[city] === -1) {
            return city
        }
    }

    return ""
};

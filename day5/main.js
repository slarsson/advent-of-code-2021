const fs = require('fs');

let size = 0;
const input = fs.readFileSync('input.txt').toString().split(/\n/).map(v => v.split('->').map(x => x.split(',').map(y => {
    const n = Number.parseInt(y);
    if (n > size) size = n;
    return n;
})));
const grid = (new Array(size+1)).fill([]).map(v => (new Array(size+1).fill(0)));

const count = (g) => {
    let hits = 0;
    for (let i = 0; i < g.length; i++) {
        for (let j = 0; j < g[i].length; j++) {
            if (g[i][j] >= 2) hits++;
        }
    }
    return hits;
}

for (let line of input) {
    let xStart = line[0][0];
    let xEnd = line[1][0];
    let yStart = line[0][1];
    let yEnd = line[1][1];    

    if (xEnd < xStart) [xStart, xEnd] = [xEnd, xStart]
    if (yEnd < yStart) [yStart, yEnd] = [yEnd, yStart]
    if (yStart == yEnd && xStart != xEnd) {
        for (let i = xStart; i <= xEnd; i++) {
            grid[yStart][i]++;
        }
    } else if (xStart == xEnd && yStart != yEnd) {
        for (let i = yStart; i <= yEnd; i++) {
            grid[i][xStart]++;
        }
    }
}

console.log('a:', count(grid));

for (let line of input) {
    let xStart = line[0][0];
    let xEnd = line[1][0];
    let yStart = line[0][1];
    let yEnd = line[1][1];

    const xStepLength = Math.max(xEnd-xStart, xStart-xEnd);
    const yStepLength = Math.max(yEnd-yStart, yStart-yEnd)
    if (xStepLength == yStepLength) {
        let dx = (xEnd < xStart) ? -1 : 1;
        let dy = (yEnd < yStart) ? -1 : 1; 
        for (let i = 0; i <= xStepLength; i++) {
            grid[yStart+(dy*i)][xStart+(dx*i)]++;
        }
    }
}

console.log('b:', count(grid));

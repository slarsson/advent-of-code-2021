const fs = require('fs');

const fn = (val, most = true) => {
    let newScore = new Array(val[0].length).fill(0);
    for (let i = 0; i < val[0].length; i++) {
        for (let item of val) {
            newScore[i] += item[i];
        }
    }
    return newScore.map(v => {
        if (most) {
            return (v >= (val.length / 2)) ? 1 : 0;
        }
        return (v >= (val.length / 2)) ? 0 : 1;
    });
}

const toDecimal = (arr) => {
    let tot = 0;
    for (let i = 0; i < arr.length; i++) {
        tot += arr[i] * Math.pow(2, (arr.length - i - 1));
    }
    return tot;
}


let matrix = fs.readFileSync('input.txt').toString().split(/\n/).map(v => {
    let arr = [];
    for (const x of v) {
        arr.push(Number.parseInt(x));
    }
    return arr;
});

console.log('a:', toDecimal(fn(matrix, true)) * toDecimal(fn(matrix, false)));

// part b
let gammaMatrix = matrix;
let epsilonMatrix = matrix.map(v => [...v]);

for (let i = 0; i < gammaMatrix[0].length; i++) {
    let gamma = fn(gammaMatrix, true);
    gammaMatrix = gammaMatrix.filter(v => {
        return gamma[i] == v[i];
    })
    if (gammaMatrix.length == 1) break;
}

for (let i = 0; i < epsilonMatrix[0].length; i++) {
    let epsilon = fn(epsilonMatrix, false);
    epsilonMatrix = epsilonMatrix.filter(v => {
        return epsilon[i] == v[i];
    })
    if (epsilonMatrix.length == 1) break;
}

console.log('b:', toDecimal(gammaMatrix[0]) * toDecimal(epsilonMatrix[0]));

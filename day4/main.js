const fs = require('fs');

const input = fs.readFileSync('input.txt').toString().split(/\n/);
const numbers = input[0].split(',').map(v => Number.parseInt(v));
const rows = input.slice(1).filter(v => v != '').map(v => v.split(' ').filter(x => x != '').map(x => Number.parseInt(x)));

let boards = [];
for (let i = 0; i < rows.length; i+=5) {
    let board = [];
    for (let j = 0; j < 5; j++) {
        board = board.concat(rows[i + j])
    }
    boards.push(board);
}

const mark = (num, board) => {
    for (let i = 0; i < board.length; i++) {
        if (board[i] == num) {
            board[i] = -1;
        }    
    }
    let sum = 0;
    for (let i = 0; i < board.length; i++) {
        if (i%5 == 0) sum = 0;
        if (board[i] == -1) sum--;
        if (sum == -5) return true;
    }
    for (let i = 0; i < 5; i++) {
        sum = 0;
        for (let j = 0; j < 5; j++) {
            if (board[i+(j*5)] == -1) sum--; 
        }
        if (sum == -5) return true;
    }
    return false;
}

const score = (board) => board.filter(v => v > 0).reduce((accu, v) => accu + v);

let done = [];
for (const number of numbers) {
    for (let i = 0; i < boards.length; i++) {
        if (done.includes(i)) continue;
        if (mark(number, boards[i])) {
            done.push(i);
            if (done.length == 1) console.log('a:', number * score(boards[i]));
            if (done.length == boards.length) console.log('b:', number * score(boards[i]));
        }
    }
}

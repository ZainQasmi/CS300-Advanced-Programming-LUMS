'use strict'

var board = [
    [ 0 , 1 , 0 , 1 , 0 ] ,
    [ 2 , 3 , 2 , 3 , 2 ] ,
    [ 0 , 1 , 0 , 1 , 0 ] ,
    [ 2 , 3 , 2 , 3 , 2 ] ,
    [ 0 , 1 , 0 , 1 , 0 ] ,
]

const cE = React.createElement;
const socket = io()
let state = {board}
var switcher = true;
var box = 'Abox'

const setState = updates => {
    Object.assign(state, updates)
    ReactDOM.render(React.createElement(Root, state), document.getElementById('root'))
}

const handleClick = event => {
    // event.preventDefault()
    console.log(event.target.name)
    if (switcher) {
        box = 'Abox'
        switcher = !switcher
    } else {
        box = 'Bbox'
        switcher = !switcher
    }
    setState({button1: box});
}

const deployBoard = (board) => {
    // Part 3: BEFORE PROCEEDING, CHECK:
    // 1. Changing only board dimensions in code changes the display
    // board.push([ 0 , 1 , 0 , 1 , 0 ])
    // YES ABOVE LINE WORKS

    let rows;
    let grid = [];
    let rowElem = [];
    let r = 0;
    let c = 0;

    grid.push( cE('div', null, `Dots and Boxes`) )

    for (r = 0; r < board.length; r++) {
        for (c = 0; c < board[0].length; c++) {
            var button_id = r*5 + c + 1;
            if (board[r][c] == 0) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'dot',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            if (board[r][c] == 1) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'hbar',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            if (board[r][c] == 2) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'vbar',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            if (board[r][c] == 3) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'box',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
        }
    rows = cE('div', {id: 'myrow'}, rowElem);
    rowElem = [];
    grid.push(rows)
    }
    return grid;
}

const Root = state => {
   return cE('div', null, deployBoard(board))
}

setState({button1: box})


// 0 => dot for a small spacer, 
// 1 => hbar for horizontal bar, 
// 2 => vbar for vertical bar,
// 3 => box for a box, 
// hbarFilled for black horizontal bar, 
// vbarFilled for black vertical bar, 
// Abox for a green box, 
// and Bbox for a red box.
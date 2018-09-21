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
var lolbox = 'Abox'
var myID = -1
var opponentID = -1
var kiskiBari = -1

const setState = updates => {
    Object.assign(state, updates)
    ReactDOM.render(React.createElement(Root, state), document.getElementById('root'))
}

socket.on('on_connect_1', data => {console.log(`You are ${data}`); myID = data;})
socket.on('on_connect_2', data => {console.log(`Your opponent is ${data}`); opponentID = data;})
socket.on('send_init_board', data => {board = data})

const handleClick = event => {
    // event.preventDefault()
    let clickedOn = event.target.name
    let className = event.target.className
    console.log(`cliekdOn:${clickedOn} Name:${className}`);
    let tempList = []; tempList.push(clickedOn); tempList.push(className);
    socket.emit('handle_click', tempList)
    
    socket.on('updated_board', newBoard => { 
        console.log('Board Received')
        board = newBoard; 
        setState();
    })

    setState({button1: lolbox});
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

    grid.push( cE('title', null, `Dots and Boxes`) )
    grid.push( cE('div', null, `Player ${kiskiBari}'s Turn`) )

    for (r = 0; r < board.length; r++) {
        for (c = 0; c < board[0].length; c++) {
            var button_id = r*5 + c + 1;
            if (board[r][c] == 0) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'dot',
                        disabled: true,
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            else if (board[r][c] == 1) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'hbar',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            else if (board[r][c] == 2) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'vbar',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            else if (board[r][c] == 3) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'box',
                        disabled: true,
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            else if (board[r][c] == 4) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'hbarFilled',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            else if (board[r][c] == 5) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'vbarFilled',
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            else if (board[r][c] == 6) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'Abox',
                        disabled: true,
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }
            else if (board[r][c] == 7) {
                rowElem.push(
                    cE('button',
                    {
                        name: button_id,
                        className: 'Bbox',
                        disabled: true,
                        onClick: ev => handleClick(ev)
                    }, button_id)
                )
            }

        }
    rows = cE('div', {id: 'myrow'}, rowElem);
    rowElem = [];
    grid.push(rows)
    }

    grid.push( cE('div', null, `You are Player ${myID}`) )
    grid.push( cE('div', null, `Your opponent is Player ${opponentID}`) )

    return grid;
}

const Root = state => {
   return cE('div', null, deployBoard(board))
}

setState({button1: lolbox})

// 0 => dot for a small spacer, 
// 1 => hbar for horizontal bar, 
// 2 => vbar for vertical bar,
// 3 => box for a box, 
// 4 => hbarFilled for black horizontal bar, 
// 5 => vbarFilled for black vertical bar, 
// 6 => Abox for a green box, 
// 7 => and Bbox for a red box.
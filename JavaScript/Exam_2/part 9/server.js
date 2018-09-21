'use strict'

const fs = require('fs'),
    http = require('http'),
    socketio = require('socket.io')

const readFile = file => new Promise((resolve, reject) =>
    fs.readFile(file, (err, data) => err ? reject(err) : resolve(data)))

const server = http.createServer(async (request, response) => {
    try {
        response.end(await readFile(request.url.substr(1)))
    } catch (err) {
        console.log(err)
        response.end()
    }
})

const io = socketio(server)

var board = [
    [ 0 , 1 , 0 , 1 , 0 ] ,
    [ 2 , 3 , 2 , 3 , 2 ] ,
    [ 0 , 1 , 0 , 1 , 0 ] ,
    [ 2 , 3 , 2 , 3 , 2 ] ,
    [ 0 , 1 , 0 , 1 , 0 ] ,
]

let clients = [];
let gamesList = [];
let turns = [];
let winnerIS = -1;

io.sockets.on('connection', socket => {
    clients = [...clients, socket]
    socket.emit('on_connect_1', clients.length);
    if (clients.length%2 == 0) {
        socket.emit('on_connect_2', clients.length-1)
        clients[clients.length-1].emit('send_init_board', board)
        clients[clients.length-2].emit('send_init_board', board)
        // gamesList.push(initMatrix);
        turns = [...turns, false]
    } else {
        socket.emit('on_connect_2', clients.length+1)
        gamesList.push(board);
        // console.log(gamesList.length)
        turns = [...turns, true]
    }

    console.log(`Client ID ${clients.length} connected`);
    // socket.on('disconnect', () => console.log('a client disconnected'))
    socket.on('disconnect', () => clients = clients.filter(s => s !== socket))
    // socket.emit('send_init_board', board);

    socket.on('handle_click', list => {
        let clickedOn = list[0]
        let className = list[1]

        for(let i = 0; i < clients.length; i++) {
            if (clients[i] == socket) {
                if (turns[i] == true) {
                    // console.log(`Client ${i+1} clicked ${clickedOn}`)
                    var newMatrix = isValid(i+1, clickedOn, className);
                    if ((i+1)%2 == 0) {
                        turns[i] = !turns[i];
                        turns[i-1] = !turns[i-1];
                        clients[i-1].emit('updated_board', newMatrix);
                        // clients[i-1].emit('whose_turn', i);
                    } else {
                        turns[i] = !turns[i];
                        turns[i+1] = !turns[i+1];
                        clients[i+1].emit('updated_board', newMatrix);
                        // clients[i+1].emit('whose_turn', i);
                    }
                    socket.emit('updated_board', newMatrix);
                    socket.emit('whose_turn', ( (i+1)%2 === 0 ? (i) : (i+2) ) );
                    (i+1)%2 === 0 ? gamesList[(i+1)-2]=newMatrix : gamesList[(i+1)-1]=newMatrix ;
                    // (i+1)%2 === 0 ? console.log(gamesList[(i+1)-2]) : console.log(gamesList[(i+1)-1]) ;
                    if (winnerIS !== -1) {
                        socket.emit('win_condition', winnerIS);
                    }

                } else {
                    console.log(`Not Client ${i+1}'s' turn. Clicked ${clickedOn}`)
                }

            }
        }

        // let clickedOn = list[0]
        // let className = list[1]
        // console.log(`${className}`)
        // console.log(`${clickedOn}`)
        // let row = Math.floor( (clickedOn/5 - 0.1));
        // let col = ( (clickedOn-1)%5 );

        // if (className === 'hbar') {
        //     board[row][col] = 4;
        //     socket.emit('updated_board', board);
        // }
        // else if (className === 'vbar') {
        //     board[row][col] = 5;
        //     socket.emit('updated_board', board);
        // }


        // for (let tempR = 0; tempR < board.length; tempR++ ) {
        //     for (let tempC = 0; tempC < board[0].length; tempC++ ) {
        //         if (validPosition(tempR-1,tempC) && board[tempR-1][tempC] === 4 
        //             && validPosition(tempR+1,tempC) && board[tempR+1][tempC] === 4
        //             && validPosition(tempR,tempC-1) && board[tempR][tempC-1] === 5
        //             && validPosition(tempR,tempC+1) && board[tempR][tempC+1] === 5
        //             ) 
        //         {
        //             console.log('All Black');
        //             board[tempR][tempC] = 6;
        //             socket.emit('updated_board', board);
        //         }
        //     }
        // }    
        

    })
})


// var theMatrix = clientID%2 ===0 ? gamesList[(clientID/2)-1] : gamesList[(Math.floor(clientID/2))]

const isValid = (clientID, clickedOn, className) => {
    console.log(`${className}`)
    console.log(`${clickedOn}`)

    let row = Math.floor( (clickedOn/5 - 0.1));
    let col = ( (clickedOn-1)%5 );
    let turnKiski = (clientID%2 === 0 ? 2: 1)

    console.log(`GAMELIST LENGTH ${gamesList.length}`)
    console.log(`CLIENT_ID ${clientID}`)

    // clientID--;
    var tempMatrix = clientID%2 ===0 ? gamesList[(clientID/2)-1] : gamesList[(Math.floor(clientID/2))]
    // var tempMatrix = clientID%2 === 0 ? gamesList[clientID-2] : gamesList[clientID-1] ;

    if (className === 'hbar') {
        tempMatrix[row][col] = 4;
    }
    else if (className === 'vbar') {
        tempMatrix[row][col] = 5;
    }

    var i = clientID-1;

    for (let tempR = 0; tempR < tempMatrix.length; tempR++ ) {
        for (let tempC = 0; tempC < tempMatrix[0].length; tempC++ ) {
            if (validPosition(tempR-1,tempC) && tempMatrix[tempR-1][tempC] === 4 
                && validPosition(tempR+1,tempC) && tempMatrix[tempR+1][tempC] === 4
                && validPosition(tempR,tempC-1) && tempMatrix[tempR][tempC-1] === 5
                && validPosition(tempR,tempC+1) && tempMatrix[tempR][tempC+1] === 5
                ) 
            {
                console.log('All Black');
                if (tempMatrix[tempR][tempC] !== 6 && tempMatrix[tempR][tempC] !== 7) {
                    if (turnKiski === 1) {
                        tempMatrix[tempR][tempC] = 6;
                        if ((i+1)%2 == 0) {
                            turns[i] = !turns[i];
                            turns[i-1] = !turns[i-1];
                        } else {
                            turns[i] = !turns[i];
                            turns[i+1] = !turns[i+1];
                        }
                    }
                    else if (turnKiski === 2) {
                        tempMatrix[tempR][tempC] = 7;
                        if ((i+1)%2 == 0) {
                            turns[i] = !turns[i];
                            turns[i-1] = !turns[i-1];
                        } else {
                            turns[i] = !turns[i];
                            turns[i+1] = !turns[i+1];
                        }
                    }
                }
            }
        }
    }

    // CHECK WIN WHEN NO BOX IS OF ID 3 
    let chkWinBool = true;
    for (let tR = 0; tR < tempMatrix.length; tR++) {
        for (let tC = 0; tC < tempMatrix[0].length; tC++) {
            if (tempMatrix[tR][tC] === 3) {
                chkWinBool = false;
            }
        }
    }
    console.log(`tester`)
    if (chkWinBool) {
        let waste = winCondition(tempMatrix);
        console.log(`wow clientID ${winnerIS} won`);
    }

    return tempMatrix
}

const winCondition = (tempMatrixWin) => {
    // let winnerIS = -1;
    let playerOneCount = 0;
    let playerTwoCount = 0;
    
    for (let r = 0; r < tempMatrixWin.length; r++) {
        for (let c = 0; c < tempMatrixWin[0].length; c++) {
            if (tempMatrixWin[r][c] === 6) {
                playerOneCount++;
            }
            else if (tempMatrixWin[r][c] === 7) {
                playerTwoCount++;
            }
        } 
    }

    if (playerOneCount === playerTwoCount) {
        winnerIS = 0;
    } else if (playerOneCount > playerTwoCount) {
        winnerIS = 1;
    } else if (playerTwoCount > playerOneCount) {
        winnerIS = 2;
    } else {
        winnerIS = -1;
    }

    return winnerIS;
}

const validPosition = (tempRow, tempCol) => {
    return (tempRow >= 0 && tempRow <=4) && (tempCol >= 0 && tempCol <=4);
}

// 1. Click on the bar, it should turn black
// 2. Stop the server and click on another bar, it should NOT turn black
// 3. Restart the server and click on the same bar, it should now turn black
// 4. Add a console.log(‘Board received’) in the message coming from server to client 
// and ensure that the message is NOT received if an already black bar is clicked AND 
// it IS received if a white bar is clicked.

server.listen(8001)

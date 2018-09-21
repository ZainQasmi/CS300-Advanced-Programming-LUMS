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
    socket.on('disconnect', () => console.log('a client disconnected'))
    // socket.emit('send_init_board', board);

    socket.on('handle_click', list => {
        let clickedOn = list[0]
        let className = list[1]
        console.log(`${className}`)
        console.log(`${clickedOn}`)

        let row = Math.floor( (clickedOn/5 - 0.1));
        let col = ( (clickedOn-1)%5 );

        if (className === 'hbar') {
            board[row][col] = 4;
            socket.emit('updated_board', board);
        }
        else if (className === 'vbar') {
            board[row][col] = 5;
            socket.emit('updated_board', board);
        }


        for (let tempR = 0; tempR < board.length; tempR++ ) {
            for (let tempC = 0; tempC < board[0].length; tempC++ ) {
                if (validPosition(tempR-1,tempC) && board[tempR-1][tempC] === 4 
                    && validPosition(tempR+1,tempC) && board[tempR+1][tempC] === 4
                    && validPosition(tempR,tempC-1) && board[tempR][tempC-1] === 5
                    && validPosition(tempR,tempC+1) && board[tempR][tempC+1] === 5
                    ) 
                {
                    console.log('All Black');
                    board[tempR][tempC] = 6;
                    socket.emit('updated_board', board);
                }
            }
        }    
        

    })
})

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

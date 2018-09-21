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

io.sockets.on('connection', socket => {
    console.log('a client connected')
    socket.on('disconnect', () => console.log('a client disconnected'))

    // socket.emit('on_connect', board);
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

    })
})


// 1. Click on the bar, it should turn black
// 2. Stop the server and click on another bar, it should NOT turn black
// 3. Restart the server and click on the same bar, it should now turn black
// 4. Add a console.log(‘Board received’) in the message coming from server to client 
// and ensure that the message is NOT received if an already black bar is clicked AND 
// it IS received if a white bar is clicked.

server.listen(8001)

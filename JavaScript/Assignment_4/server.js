const fs = require('fs'),
	http = require('http'),
	socketio = require('socket.io')

const readFile = file => new Promise((resolve, reject) => 
	fs.readFile(file, (err,data) => err ? reject(err) : resolve(data)))

const server = http.createServer(async (request, response) => {
	// console.log(`Request received for ${request.url}`)
	try {
		response.end(await readFile (request.url.substr(1)))
	} catch (err) {
		response.end()
	}
})

// 1 is X, 2 is 0

const initMatrix = [
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 ,  1 ,  0 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 ,  0 ,  1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ]
]

let clients = [];
let gamesList = [];
let turns = [];
const io = socketio(server)
let winnerIS = -1;

io.sockets.on('to_server', data => clients.forEach(s => {
		s.emit('to_client', data);
		console.log(data)
	}))

io.sockets.on('connection', socket => {
	clients = [...clients, socket]
	socket.emit('on_connect_1', clients.length);
	if (clients.length%2 == 0) {
		socket.emit('on_connect_2', clients.length-1)
		clients[clients.length-1].emit('send_board', initMatrix)
		clients[clients.length-2].emit('send_board', initMatrix)
		// gamesList.push(initMatrix);
		turns = [...turns, false]
	} else {
		socket.emit('on_connect_2', clients.length+1)
		gamesList.push(initMatrix);
		// console.log(gamesList.length)
		turns = [...turns, true]
	}
	// console.log(turns);
	console.log(`Client ID ${clients.length} connected`);

	socket.on('disconnect', () => clients = clients.filter(s => s !== socket))
	// socket.on('to_server', data => clients.forEach(s => s.emit('to_client', data)))
	
	socket.on('handle_click', clickedOn => {
		for(i = 0; i < clients.length; i++) {
			if (clients[i] == socket) {
				if (turns[i] == true) {
					// console.log(`Client ${i+1} clicked ${clickedOn}`)
					var newMatrix = isValid(i+1, clickedOn);
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

					if (winnerIS !== -1) {
						socket.emit('win_condition', winnerIS);
					}

					// (i+1)%2 === 0 ? console.log(gamesList[(i+1)-2]) : console.log(gamesList[(i+1)-1]) ;
				} else {
					console.log(`Not Client ${i+1}'s' turn. Clicked ${clickedOn}`)
				}

			}
		}
		// console.log(turns);
	})
})

// const clientIDoddX = 0;
// const clientIDevenO = 1;

// const isValid = (clientID, row, col) => {

// }

const validPosition = (tempRow, tempCol) => {
	return (tempRow >= 0 && tempRow <=7) && (tempCol >= 0 && tempCol <=7);
}

const isValid = (clientID, clickedOn) => {
	// console.log(`${clientID} .... ${clickedOn}`)
	clickedOn--;
	let row = Math.floor( (clickedOn/8)%8 );
	let col = Math.floor( clickedOn%8 );
	// isValid(clientID, row,col)
	console.log(`HERE ${gamesList.length}`)
	console.log(`CLIENT_ID ${clientID}`)
	var theMatrix = clientID%2 === 0 ? gamesList[clientID-2] : gamesList[clientID-1] ;
	var opponentSymbol = clientID%2 === 0 ? 0 : 1 ;
	var currentSymbol = clientID%2 === 0 ? 1 : 0 ; 
	var tempRow;	
	var tempCol;
	var entireLine = [];

	// check all dirs
	for(var horiz = -1; horiz <= 1; horiz++) {
		for (var vert = -1; vert <= 1; vert++) {
			if (vert === 0 && horiz === 0) {
				continue;
			}
			tempRow = row + horiz;
			tempCol = col + vert;

			var potentialList = []; //unused

			while (validPosition(tempRow,tempCol) && theMatrix[tempRow][tempCol] !== -1 && theMatrix[tempRow][tempCol] === opponentSymbol) {
				potentialList.push([tempRow,tempCol]);
				//next in line
				tempRow = tempRow + horiz;
				tempCol = tempCol + vert;
				// console.log(`2:: Trow::${tempRow} :: Tcol::${tempCol} :: validPosition(tempRow,tempCol)`)
			}

			if(potentialList.length && validPosition(tempRow,tempCol)) { //if List is NOT empty
				if (theMatrix[tempRow][tempCol] === currentSymbol) { //check if next in LINE is CURRENT player
					entireLine.push([row, col]);
					potentialList.forEach(one => entireLine.push(one));
				}
			}
		}
	}

	// console.log(entireLine);
	// UPDATE BOARD
	entireLine.forEach(oneTuple => {
		theMatrix[oneTuple[0]][oneTuple[1]] = currentSymbol;
	})

	if (winCondition(theMatrix, clientID)) {
		console.log('wow someone won');
	}

	return theMatrix;

}

const winCondition = (theMatrix, clientID) => {
	win = true;
	countX = 0;
	countO = 0;
	for (r = 0; r < 8; r++) {
		for (c = 0; c < 8; c++) {
			if (theMatrix[r][c] == -1) {
				win = false;
			}
			else if (theMatrix[r][c] == 1) {
				countX++;
			} else if (theMatrix[r][c] == 0 ) {
				countO++;
			}
		}
	}
	if (win) {
		if (countX === countO) {
			winnerIS = 2
		} else if (countX > countO){
			winnerIS = 1;
		} else if (countX < countO){
			winnerIS = 0;
		} else {
			winnerIS = -1;
		}	
	}
	return win;
}

server.listen(8001)
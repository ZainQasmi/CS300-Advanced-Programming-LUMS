const cE = React.createElement;
const socket = io()
let state = {}
var myID = -1
var opponentID = -1
var kiskiBari = -1

var theMatrix = [
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 ,  1 ,  0 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 ,  0 ,  1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
	[ -1 , -1 , -1 , -1 , -1 , -1 , -1 , -1 ] ,
]

const setState = updates => {
	// console.log('click hoa bee sea')
	Object.assign(state, updates)
	// console.log(state)
	ReactDOM.render( cE(Root, state), document.getElementById('root') )
}

socket.on('to_client', data => setState({msgList: [...state.msgList, data]}))
socket.on('on_connect_1', data => {console.log(`You are ${data}`); myID = data;})
socket.on('on_connect_2', data => {console.log(`Your opponent is ${data}`); opponentID = data;})
socket.on('send_board', data => { theMatrix = data; /*console.log(theMatrix);*/ })
// socket.on('whose_turn', turns => { console.log(`1:It is Player ${turns}'s turn`)})

const handleClick = event => {
	// event.preventDefault()
	console.log(event.target.name)
	socket.emit('handle_click', event.target.name)
	socket.on('updated_board', newMatrix => { theMatrix = newMatrix; setState();})
	socket.on('whose_turn', turns => {  kiskiBari = turns ;console.log(`2:It is Player ${turns}'s turn`); setState();})
	socket.on('win_condition', isWin => {
		if (isWin == 1) {
			console.log('X wins')
		} else if (isWin == 0) {
			console.log('O wins')
		} else if (isWin == 2) {
			console.log('its a tie')
		}
	})

	setState({message: ''})
}

var button_height = '32px'
var button_width = '32px'

const deployBoard = theMatrix => {
	var rows;
	grid = [];
	rowElem = [];
	var letter;

	// grid.push( cE( 'input', {value: 'Play', type: 'submit'}) )

	grid.push( cE('div', null, `Player ${kiskiBari}'s Turn`) )
	grid.push( cE('title', null, 'Dimagh Kharab') )

	for (r = 0; r < 8; r++) {
		for (c = 0; c < 8; c++) {
			var button_id = r*8 + c + 1
			if (theMatrix[r][c] == -1) {
				rowElem.push(
					cE('button',
					{
						type: 'submit',
						name: button_id,
						style: {height: button_height, width: button_width},
						disabled: false,
						onClick: ev => handleClick(ev)
						// value: '-',
						// onClick: ev => setState({message: ev.target.value})
					}, '-')
				)
			} else if (theMatrix[r][c] == 0) {
				rowElem.push(
					cE('button',
					{
						type: 'submit',
						name: button_id,
						style: {height: button_height, width: button_width},
						disabled: false,
						onClick: ev => handleClick(ev)
						// value: 'X',
						// onClick: ev => setState({message: ev.target.value})
					}, 'O')
				)
			} else if (theMatrix[r][c] == 1) {
				rowElem.push(
					cE('button',
					{
						type: 'submit',
						name: button_id,
						style: {height: button_height, width: button_width},
						disabled: false,
						onClick: ev => handleClick(ev)
						// value: 'O',
						// onClick: ev => setState({message: ev.target.value})
					}, 'X')
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

const Root = ({message, msgList}) => {
	return cE('div', null, deployBoard(theMatrix))
}

setState({message: '', msgList: []})

/*
	cE(`form`, {onClick: handleClick},
		cE('title', null, 'Dimagh Kharab'),
		cE('h2',null,'Wanna Reversi?'),

		
		table2 = table.map (tableAik => 
			cE(`button`, {
				value: val, 
				type: 'text',
				onClick: ev => setState({message: ev.target.value})
			}, val)
		),
		table2.map(table => cE('div', null, table2)),
		

		
		React.createElement(`input`, {
			value: message, 
			type: 'text',
			onChange: ev => setState({message: ev.target.value})
		}),
		React.createElement('input', {
			value: 'Send', 
			type: 'submit'
		}),		
		
		msgList.map(msg => React.createElement('div', null, msg))	

	)
*/

'use strict'

const socket = io()
let state = {}
var switcher = true;
var box = 'Abox'

const setState = updates => {
    Object.assign(state, updates)
    ReactDOM.render(React.createElement(Root, state), document.getElementById('root'))
}

const handleClick = event => {
    // event.preventDefault()
    if (switcher) {
        box = 'Abox'
        switcher = !switcher
    } else {
        box = 'Bbox'
        switcher = !switcher
    }
    console.log(event.target.className)
    setState({button1: box});
}

const Root = state =>
    React.createElement('div', null, 
        React.createElement('div', null,
            React.createElement('button', {
                className: 'dot'
            }, '1'),
            React.createElement('button', {
                className: 'hbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '2'),
            React.createElement('button', {
                className: 'dot'
            }, '3'),
            React.createElement('button', {
                className: 'hbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4'),
            React.createElement('button', {
                className: 'dot',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4')
        ),
        React.createElement('div', null,
            React.createElement('button', {
                className: 'vbar'
            }, '1'),
            React.createElement('button', {
                className: 'box',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '2'),
            React.createElement('button', {
                className: 'vbar'
            }, '3'),
            React.createElement('button', {
                className: 'box',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4'),
            React.createElement('button', {
                className: 'vbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4')
        ),
        React.createElement('div', null,
            React.createElement('button', {
                className: 'dot'
            }, '1'),
            React.createElement('button', {
                className: 'hbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '2'),
            React.createElement('button', {
                className: 'dot'
            }, '3'),
            React.createElement('button', {
                className: 'hbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4'),
            React.createElement('button', {
                className: 'dot',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4')
        ),
        React.createElement('div', null,
            React.createElement('button', {
                className: 'vbar'
            }, '1'),
            React.createElement('button', {
                className: 'box',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '2'),
            React.createElement('button', {
                className: 'vbar'
            }, '3'),
            React.createElement('button', {
                className: 'box',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4'),
            React.createElement('button', {
                className: 'vbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4')
        ),
        React.createElement('div', null,
            React.createElement('button', {
                className: 'dot'
            }, '1'),
            React.createElement('button', {
                className: 'hbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '2'),
            React.createElement('button', {
                className: 'dot'
            }, '3'),
            React.createElement('button', {
                className: 'hbar',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4'),
            React.createElement('button', {
                className: 'dot',
                onClick: ev => handleClick(ev)
                // onClick: ev => console.log('bar clicked')
            }, '4')
        )
        
    )

setState({button1: box})


// hbar for horizontal bar, hbarFilled for black horizontal bar, vbar for vertical bar,
// vbarFilled for black vertical bar, dot for a small spacer, box for a box, Abox for a 
// green box, and Bbox for a red box.
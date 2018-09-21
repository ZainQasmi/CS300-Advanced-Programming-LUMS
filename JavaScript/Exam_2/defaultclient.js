'use strict'

const socket = io()
let state = {}

const setState = updates => {
    Object.assign(state, updates)
    ReactDOM.render(React.createElement(Root, state), document.getElementById('root'))
}

const Root = state =>
    React.createElement('div', null,
        React.createElement('button', {
            className: state.button1
        }, null),
        React.createElement('button', {
            className: 'vbar',
            onClick: ev => console.log('bar clicked')
        }, null))

setState({button1: 'Abox'})

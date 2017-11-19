import React, { Component } from 'react';
import Board from './Board';
export default class BoardHolder extends Component {
    constructor() {
        super()
        this.state = {
            mode: 'singleplayer',
        }
    }
    toggleMode = () => {
        const { mode } = this.state
        this.setState({
            mode: !mode
        })
    }
    onModeSwitch = (e) => {
        this.toggleMode()
    }
    render() {
        return (
            <div>
            <p> Welcome to Connect4 Live! </p>
              
                <div className="board-container">
                    <Board />
                </div>
            </div>
        )
    }
}

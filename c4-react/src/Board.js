import React, { Component } from 'react';

export default class Board extends Component {
    constructor() {
        super()
        this.socket = new WebSocket("ws://localhost:8080/ws")
        var gameComplete = false
        this.socket.onmessage = (e) => {
            const data = JSON.parse(e.data);
            console.log(e.data)
            this.setState({
                board: data.gameboard,
                status: data.status,
                numMoves: data.nummoves,
                // etc. etc.
            })
            if (data.status==="Game has finished!"){
                gameComplete = true
                window.alert("the game is over! refresh to play again")
            }

        }
        this.state = {
            board: []
        }
    }
    onClick = (col) => {
        this.socket.send(col)
    }
    renderBoard = () => {
        const { board } = this.state
        return board.map((row) => {
            return (
                <div className="row">
                    {
                        row.map((slot, i) => {
                            const { active, symbol } = slot
                            if (active) {
                                return (
                                    <div className={`${symbol === 'X' ? 'slot--red' : 'slot--black' }`}>
                                    </div>
                                )
                            } else {
                                return <div onClick={ () => { this.onClick(i) }  } className="slot--empty"></div>
                            }
                        })
                    }
                </div>
            )
        })
    }
    render() {
        return (
            <div className="board">
                { this.renderBoard() }
            </div>
        )
    }
}

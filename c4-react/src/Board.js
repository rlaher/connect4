import React, { Component } from 'react';

export default class Board extends Component {
    constructor() {
        super()
        var gameComplete = false
        //this.socket = new WebSocket("ws://ec2-18-217-35-187.us-east-2.compute.amazonaws.com:8080/ws")
        this.socket = new WebSocket("ws://localhost:8080/ws")
        this.socket.onmessage = (e) => {
            const data = JSON.parse(e.data);
            console.log(e.data)
            this.setState({
                board: data.gameboard,
                status: data.status,
                numMoves: data.nummoves,
                playersTurn: data.playersturn,
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

    togglePlayerMode = () =>{
        this.socket.send(999)
    }
    render() {
        const { playersTurn } = this.state
        const {status} = this.state
        return (
            <div>
            <div className="playersTurn">
            Players Turn:
            {playersTurn  === 0 ? 'Red' : 'Blue'}
            </div>
            <div className="status">
            Game Status: {status}
            </div>
           
            <div className="board">
                { this.renderBoard() }
            </div>
            <div className="mode-selector">
                    <label class="switch">
                        <input type="checkbox"onClick ={this.togglePlayerMode}/>
                        <span class="slider">Tick this box to play vs AI</span>
                    </label>
            </div>
            </div>
        )
    }
}

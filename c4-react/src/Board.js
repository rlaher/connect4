import React, { Component } from 'react';

export default class Board extends Component {
    constructor() {
        super()
        var gameComplete = false
        this.socket = new WebSocket("ws://ec2-18-217-35-187.us-east-2.compute.amazonaws.com:8080/ws")
        //this.socket = new WebSocket("ws://localhost:8080/ws")
        this.socket.onmessage = (e) => {
            const data = JSON.parse(e.data);
            console.log(e.data)
            this.setState({
                board: data.gameboard,
                status: data.status,
                numMoves: data.nummoves,
                playersTurn: data.playersturn,
                computerdifficulty: data.computerdifficulty,
                // etc. etc.
            })
            if (data.status==="Game has finished!"){
                gameComplete = true
                this.render()
                window.alert("the game is over! refresh to play again")
            }

        }
        this.state = {
            checked: false,
            board: []
        }
    }
    onClick = (col) => {
        this.socket.send(col)
    }

    buttonHandler = (i)=>{
        this.socket.send(990+i)
 
    }

    renderButtons = () => {
        if(this.state.checked){
            const buttons = []
            for (let i = 1; i <= 5; i++) {
                buttons.push(
                    <button type="button" className="difficulty-button" onClick={ () => this.buttonHandler(i) }>
                        Level{ i }
                    </button>
                )
            }
            return buttons
        }
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
        this.state.checked = !this.state.checked
        this.socket.send(999)
    }
    render() {
        const { playersTurn } = this.state
        const {status} = this.state
        const {computerdifficulty} = this.state
        return (
            <div>
                <p className="welcome"> Connect4 Live! </p>
                <div className="big-box">      
                     <div className="left-panel">
                         <div className="modeselektor">
                            <label class="switch">
                                <input type="checkbox"onClick ={this.togglePlayerMode}/>
                                <span class="slider">Play vs AI</span>
                            </label>
                        </div>
                        <div className="buttons">
                            {this.renderButtons()}
                        </div>
                    </div>
                    <div className="board-container">
                        
                        <div className="status">
                        Game Status: {status}
                        </div>
                        <div className="board">
                            { this.renderBoard() }
                        </div>
                        <div className="playersTurn">
                            Players Turn:
                            {playersTurn  === 0 ? 'Red' : 'Blue'}
                        </div>
                        <div className="computer-difficulty">
                            Computer difficulty: {computerdifficulty}
                        </div>
                    </div>
            
                    <div className="right-panel">
                        <div className="about-text">
                            <p>This is a project I made in my free time as a way to familiarize myself with aspects of programming/the internet that I haven’t had much experience with. The backend is coded in Go, and utilizes channels and websockets to allow multiple groups of users to play against each other simultaneously. Send the link to a friend and try it out! (Or just open it again in a new tab).</p>
                            <p>The AI was coded using the minimax algorithm. Basically, it calculates every possible game state that could occur up to a variable depth. It assigns each game state a value based on a simple heuristic I came up with, and returns the move that yields the most profitable future state. The difficulty level is how many moves ahead the computer calculates.</p>
                            <p>Check out the code here: https://github.com/rlaher/connect4-live </p>
                        </div>
                    </div>
                </div>
           
        </div>
        )
    }
}

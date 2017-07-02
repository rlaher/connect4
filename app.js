import React from 'react'

 class Connect4 extends React.Component {
    render(){
        return(
            <div>TEST</div>
            <div className="connect4-column" key={`column-${y}`}>

         {column.map((cell, x) => {

           let cellClasses = classNames({
             'connect4-cell': true,
             'connect4-cell--red': (cell === 'red'),
             'connect4-cell--blue': (cell === 'blue')
           });

           return (
             <Cell key={`cell-${x}-${y}`}
               x={x}
               y={y}
               cell={cell.toString()}
               nextPlayer={board.nextPlayer}
               addPiece={this.props.addPiece} />
           );

         })}
       </div>
        )

    }
}

ReactDOM.render(<Connect4 />, document.getElementById('container'));

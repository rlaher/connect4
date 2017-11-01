import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Board from'./Board';



class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2 id="status">Put game status here</h2>
        </div>
    
        <Board/>

    </div>


    );
  }
}

export default App;

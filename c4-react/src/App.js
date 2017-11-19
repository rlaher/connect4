import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Board from'./Board';
import BoardHolder from './Boardholder';



class App extends Component {
  render() {
    return (
      <div className="App">
        <BoardHolder/>

    </div>


    );
  }
}

export default App;

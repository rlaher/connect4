import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

var socket = new WebSocket("ws://localhost:8080/ws");

function clickHandler(e){
    socket.send(1)
    console.log("1 sent")
}

socket.onmessage = function(evt){
    var newData = JSON.parse(evt.data);
    console.log(evt.data); //remove later
    document.getElementByID("myboardstring") = newData.boardasstring
};


class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to React</h2>
        </div>
        <button id = "testbutton" onClick={clickHandler}>testing button </button>
        <div id="myboardstring"></div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;

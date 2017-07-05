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
    document.getElementById("myboardstring1").innerHTML = newData.boardasstring1
    document.getElementById("myboardstring2").innerHTML = newData.boardasstring2
    document.getElementById("myboardstring3").innerHTML = newData.boardasstring3
    document.getElementById("myboardstring4").innerHTML = newData.boardasstring4
    document.getElementById("myboardstring5").innerHTML = newData.boardasstring5
    document.getElementById("myboardstring6").innerHTML = newData.boardasstring6

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
        <p id="myboardstring1">no board yet</p>
        <p id="myboardstring2">no board yet</p>
        <p id="myboardstring3">no board yet</p>
        <p id="myboardstring4">no board yet</p>
        <p id="myboardstring5">no board yet</p>
        <p id="myboardstring6">no board yet</p>

        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;

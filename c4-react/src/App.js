import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

var socket = new WebSocket("ws://localhost:8080/ws");

function clickHandler0(e){
    socket.send(0)
    console.log("0 sent")
}
function clickHandler1(e){
    socket.send(1)
    console.log("1 sent")
}
function clickHandler2(e){
    socket.send(2)
    console.log("2 sent")
}
function clickHandler3(e){
    socket.send(3)
    console.log("3 sent")
}
function clickHandler4(e){
    socket.send(4)
    console.log("4 sent")
}
function clickHandler5(e){
    socket.send(5)
    console.log("5 sent")
}
function clickHandler6(e){
    socket.send(6)
    console.log("6 sent")
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

    document.getElementById("status").innerHTML = newData.status

};



class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2 id="status">Put game status here</h2>
        </div>
        <div id="buttons">
            <button id = "zero" onClick={clickHandler0}>col zero </button>
            <button id = "one" onClick={clickHandler1}>col one </button>
            <button id = "two" onClick={clickHandler2}>col two </button>
            <button id = "three" onClick={clickHandler3}>col three </button>
            <button id = "four" onClick={clickHandler4}>col four </button>
            <button id = "five" onClick={clickHandler5}>col five </button>
            <button id = "six" onClick={clickHandler6}>col six </button>
        </div>
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

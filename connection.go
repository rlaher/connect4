package main

import (
	"connect4/game"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

// connections stores all the hubs
var connections []*connectionPair

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connection struct {
	// Channel which triggers the connection to update the gameState
	doBroadcast chan bool
	// The connectionPair. Holds up to 2 connections.
	cp *connectionPair
	// playerNum represents the players Slot. Either 0 or 1
	playerNum int
}

type connectionPair struct {
	// the mutex to protect connections
	connectionsMx sync.RWMutex
	// Registered connections.
	connections map[*connection]struct{}
	// Inbound messages from the connections.
	receiveMove chan bool
	logMx       sync.RWMutex
	log         [][]byte
	game        *game.Game
}

func newConnectionPair() *connectionPair {
	cp := &connectionPair{
		connectionsMx: sync.RWMutex{},
		receiveMove:   make(chan bool),
		connections:   make(map[*connection]struct{}),
		game:          game.NewGame(),
	}
	return cp
}

// addConnection adds a players connection to the connectionPair
func (h *connectionPair) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()

	h.connections[conn] = struct{}{}

}

// removeConnection removes a players connection from the connectionPair
func (h *connectionPair) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.doBroadcast)
	}
	log.Println("Player disconnected")
	//SHOULD DO SOMETHING HERE
}

// reader reads the moves from the clients ws-connection
func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		//Reading next move from connection here
		_, clientMoveMessage, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		field, _ := strconv.ParseInt(string(clientMoveMessage[:]), 10, 32) //Getting FieldValue From Player Action
		c.cp.game.MakeMove(c.playerNum, int(field))
		c.cp.receiveMove <- true //telling connectionPair to broadcast the gameState
	}
}

// writer broadcasts the current gameState to the two players in a connectionPair
func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for range c.doBroadcast {
		sendGameStateToConnection(wsConn, c)
	}
}

// getConnectionPairWithEmptySlot looks trough all connectionPairs and finds one which has only 1 player
// if there is none a new connectionPair is created and the player is added to that pair
func getConnectionPairWithEmptySlot() (*connectionPair, int) {
	sizeBefore := len(connections)
	// find connections with 1 player first and pair if possible
	for _, h := range connections {
		if len(h.connections) == 1 {
			log.Printf("Players paired")
			return h, len(h.connections)
		}
	}

	//TODO: I need to remove orphaned connectionPairs from the stack

	// if no emtpy slow was found at all, we create a new connectionPair
	h := newConnectionPair()
	connections = append(connections, h)
	log.Printf("Player seated in new connectionPair no. %v", len(connections))
	return connections[sizeBefore], 0
}
func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("connection upgraded!")
	//Adding Connection to connectionPair
	cp, pn := getConnectionPairWithEmptySlot()
	c := &connection{doBroadcast: make(chan bool), cp: cp, playerNum: pn}
	c.cp.addConnection(c)

	//If the connectionPair existed before but one player was disconnected
	//we can now reinitialize the gameState after the remaining player has
	// //been paired again
	// if c.cp.game.Status == resetWaitPaired {
	// 	c.cp.gs = newGameState()
	// 	//there is already one player connected when we re-pair
	// 	c.cp.gs.numberOfPlayers = 1
	// 	log.Println("gamestate resetted")
	// }

	//inform the gameState about the new player
	c.cp.game.AddPlayer()
	//telling connectionPair to broadcast the gameState
	c.cp.receiveMove <- true

	//creating the writer and reader goroutines
	//the websocket connection is open as long as these goroutines are running
	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, conn)
	go c.reader(&wg, conn)
	wg.Wait()
	conn.Close()
}

// sendGameStateToConnection broadcasts the current gameState as JSON to all players
// within a connectionPair
func sendGameStateToConnection(wsConn *websocket.Conn, c *connection) {
	err := wsConn.WriteMessage(websocket.TextMessage, c.cp.game.JsonEncode())
	//removing connection if updating gameState fails
	if err != nil {
		c.cp.removeConnection(c)
	}
}

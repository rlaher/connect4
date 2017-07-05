package main

import (
	"connect4/game"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// connections is slice of all the pairs of players
var connections []*connectionPair

//upgrader necessary to turn our connection into a websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connection struct {
	// Channel when true triggers writer go routine to send out game state to all
	doBroadcast chan bool
	cp          *connectionPair
	playerNum   int
}

type connectionPair struct {
	// the mutex to protect connections NOTE: could be implemented more extensively
	connectionsMx sync.RWMutex
	connections   map[*connection]struct{}
	//channel when true means inbound move
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

	go func() {
		for {
			//accept move
			<-cp.receiveMove
			for c := range cp.connections {
				select {
				//tell writer to broadcast game state
				case c.doBroadcast <- true:
					//if 5 seconds pass without signal, something wonky
					//happened to connection so remove it
				case <-time.After(5 * time.Second):
					cp.removeConnection(c)
				}
			}
		}
	}()

	return cp
}

// adds player to connection pair
func (h *connectionPair) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()

	h.connections[conn] = struct{}{}

}

// removes connection from connectionpair and sets game status to broken
func (h *connectionPair) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.doBroadcast)
	}
	log.Println("Player disconnected, connection pair removed")
	//SHOULD DO MORE HERE
	h.game.Status = game.Broken
}

//reads move from websocket and makes move in Game
func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		//Reading next move from connection here
		_, clientMoveMessage, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Print("movemessage:")
		fmt.Println(clientMoveMessage)

		field, _ := strconv.ParseInt(string(clientMoveMessage[:]), 10, 32)
		fmt.Print("conv to int:")
		fmt.Println(field)

		c.cp.game.MakeMove(c.playerNum, int(field))
		c.cp.receiveMove <- true
		//tells connectionpair that move has occurred
		//connectionpair then sends doBroadcast true
		//which tells writer to broadcast the gamestate
	}
}

// sends out gamestate when doBroadcast <- true
func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for range c.doBroadcast {
		sendGameStateToConnection(wsConn, c)
	}
}

//if empty connectionpair exists, matches players together to form full connection pair
//otherwise, create new connectionpair which will wait for next player
func matchMaker() (*connectionPair, int) {
	sizeBefore := len(connections)

	for _, h := range connections {
		if len(h.connections) == 1 {
			log.Printf("Players paired")
			return h, 1
		}
	}

	// create new connectionpair
	h := newConnectionPair()
	//add it to connections[]
	connections = append(connections, h)
	log.Printf("Player seated in new connectionPair no. %v", len(connections))
	return connections[sizeBefore], 0
}

//handler is called when user visits page
//upgrades to websocket connection
//puts player in a connection pair
//calls reader and writer go routines
func handler(w http.ResponseWriter, r *http.Request) {
	//upgrade to websocket connection for real time updates
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("connection upgraded!")
	//Adding Connection to connectionPair
	cp, pn := matchMaker()
	c := &connection{doBroadcast: make(chan bool), cp: cp, playerNum: pn}
	c.cp.addConnection(c)

	//add player to the game
	c.cp.game.AddPlayer()

	//i think this line is unnecessary
	//uncomment if code breaks later on
	//c.cp.receiveMove <- true

	//waitgroup waits for reader and writer to both be closed before proceeding
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
	//removing connection if error
	if err != nil {
		fmt.Println(err.Error())
		c.cp.removeConnection(c)
	}
}

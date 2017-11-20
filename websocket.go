package main

import (
	"connect4/game"
	"connect4/minimax"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var multiplayer = true

// connections is slice of all the pairs of players
var connections []*gameRoom

//upgrader necessary to turn our connection into a websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connection struct {
	// Channel when true triggers writer go routine to send out game state to all
	doBroadcast chan bool
	//holds the gameRoom it belongs to
	gr        *gameRoom
	playerNum int
}

//holds two connections
//think of gameRoom as a lobby for a game
type gameRoom struct {
	connectionsMx sync.RWMutex // the mutex to protect connections NOTE: could be implemented more extensively
	connections   map[*connection]struct{}
	receiveMove   chan bool //channel when true means inbound move
	logMx         sync.RWMutex
	log           [][]byte
	game          *game.Game
}

func newgameRoom() *gameRoom {
	gr := &gameRoom{
		connectionsMx: sync.RWMutex{},
		receiveMove:   make(chan bool),
		connections:   make(map[*connection]struct{}),
		game:          game.NewGame(),
	}

	go func() {
		for {
			//accept move
			<-gr.receiveMove
			for c := range gr.connections {
				select {
				//tell writer to broadcast game state
				case c.doBroadcast <- true:
					//if 5 seconds pass without signal, something wonky
					//happened to connection so remove it
				case <-time.After(5 * time.Second):
					gr.removeConnection(c)
				}
			}
		}
	}()

	return gr
}

// adds player to connection pair
func (h *gameRoom) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()

	h.connections[conn] = struct{}{}

}

// removes connection from gameRoom and sets game status to broken
func (h *gameRoom) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.doBroadcast)
	}
	log.Println("Player disconnected, ")
	// err := wsConn.WriteMessage(websocket.TextMessage, conn.gr.game.JsonEncode())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//SHOULD DO MORE HERE
	h.game.Status = game.Broken
	for i, k := range connections {
		if h == k {
			log.Printf("shutting down game room")
			connections = append(connections[:i], connections[i+1:]...)
		}
	}

}

//reads move from websocket and makes move in Game
func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn, multiplayerFlag bool) {
	defer wg.Done()
	for {
		//Reading next move from connection here
		_, clientMoveMessage, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Print("message:")

		field, _ := strconv.ParseInt(string(clientMoveMessage[:]), 10, 32)
		fmt.Println(field)
		if field == 999 {
			c.gr.toggleGameMode(wsConn, c, multiplayerFlag)
			return
		}
		c.gr.game.MakeMove(c.playerNum, int(field))
		c.gr.receiveMove <- true
		//tells gameRoom that move has occurred
		//gameRoom then sends doBroadcast true
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

//if empty gameRoom exists, matches players together to form full room
//otherwise, create new gameRoom which will wait for next player
func matchMaker() (*gameRoom, int) {
	sizeBefore := len(connections)

	for _, h := range connections {
		if len(h.connections) == 1 {
			log.Printf("Players paired")
			return h, 1
		}
	}

	// create new gameRoom
	h := newgameRoom()
	//add it to connections[]
	connections = append(connections, h)
	log.Printf("Player seated in new gameRoom no. %v", len(connections))
	return connections[sizeBefore], 0
}

//handler is called when user visits page
//upgrades to websocket connection
//puts player in a connection pair
//calls reader and writer go routines
func handler(w http.ResponseWriter, r *http.Request) {
	//upgrade to websocket connection for real time updates
	fmt.Println("handler appears to have been called")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	multiplayerFlag := true
	twoplayer(conn, multiplayerFlag)
}
func twoplayer(conn *websocket.Conn, multiplayerFlag bool) {
	//Adding Connection to gameRoom
	gr, pn := matchMaker()
	c := &connection{doBroadcast: make(chan bool), gr: gr, playerNum: pn}
	c.gr.addConnection(c)

	//add player to the game
	c.gr.game.AddPlayer()

	//i think this line is unnecessary
	//uncomment if code breaks later on
	c.gr.receiveMove <- true

	//waitgroup waits for reader and writer to both be closed before proceeding
	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, conn)
	go c.reader(&wg, conn, multiplayerFlag)
	wg.Wait()
	conn.Close()
}

func oneplayer(conn *websocket.Conn, multiplayerFlag bool) {
	myGameRoom := newgameRoom()
	c := &connection{doBroadcast: make(chan bool), gr: myGameRoom, playerNum: 0}

	mygame := myGameRoom.game
	mygame.Status = "You are going to lose against the AI"
	err := conn.WriteMessage(websocket.TextMessage, mygame.JsonEncode())
	if err != nil {
		fmt.Printf("could not write ")
	}
	for !mygame.IsComplete {
		//Reading next move from connection here
		_, clientMoveMessage, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Print("movemessage:")

		field, _ := strconv.ParseInt(string(clientMoveMessage[:]), 10, 32)
		fmt.Println(field)
		if field == 999 {
			myGameRoom.toggleGameMode(conn, c, multiplayerFlag)
			return
		}

		mygame.MakeMove(0, int(field))
		aimove := minimax.Minimax(4, *mygame, 1)
		mygame.MakeMove(1, aimove)
		err = conn.WriteMessage(websocket.TextMessage, mygame.JsonEncode())

	}
}

func (h *gameRoom) toggleGameMode(conn *websocket.Conn, c *connection, multiplayerFlag bool) {
	multiplayer := !multiplayerFlag
	fmt.Println("multiplayer:")
	fmt.Println(multiplayer)
	if !multiplayer {
		h.removeConnection(c)
		oneplayer(conn, multiplayer)
	} else {
		//might need to call remove connection...
		twoplayer(conn, multiplayer)

	}

}

// sendGameStateToConnection broadcasts the current gameState as JSON to all players
// within a gameRoom
func sendGameStateToConnection(wsConn *websocket.Conn, c *connection) {
	err := wsConn.WriteMessage(websocket.TextMessage, c.gr.game.JsonEncode())
	//removing connection if error
	if err != nil {
		fmt.Println(err.Error())
		c.gr.removeConnection(c)
	}
}

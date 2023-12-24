class WebSocketClient {
  constructor(name, roomName, msgCbFunc = (message) => {}) {
    this.ws = null;
    this.serverUrl = "ws://localhost:8080/ws?name=" + name;
    this.messages = null;
    this.onReceiveMessage = msgCbFunc;
    this.users = [];

    this._connectToWebsocket();
    this.joinRoom(roomName);
  }

  _connectToWebsocket() {
    this.ws = new WebSocket(this.serverUrl);
    this.ws.addEventListener('open', (event) => { this._onWebsocketOpen(event) });
    this.ws.addEventListener('message', (event) => { this._handleNewMessage(event) });
  }

  _onWebsocketOpen() {
    console.log("connected to WS!");
  }

  _handleNewMessage(event) {
    let data = event.data;
    data = data.split(/\r?\n/);
  
    for (const element of data) {
      let msg = JSON.parse(element);
      switch (msg.action) {
        case "send-message":
          this._handleSendMessage(msg.message);
          break;
        case "user-join":
          this._handleUserJoined(msg);
          break;
        case "user-left":
          this._handleUserLeft(msg);
          break;
        case "room-joined":
          this._handleRoomJoined(msg);
          break;
        default:
          break;
      }
  
    }
  }

  _handleSendMessage(msg) {
    this.messages.push(msg);
    this.onReceiveMessage(msg);
  }

  _handleUserJoined(msg) {
    this.users.push(msg.sender);
  }

  _handleUserLeft(msg) {
    for (let i = 0; i < this.users.length; i++) {
      if (this.users[i].id == msg.sender.id) {
        this.users.splice(i, 1);
      }
    }
  }

  _handleRoomJoined(message) {
    this.room = message.target;
    this.room.name = this.room.private ? message.sender.name : this.room.name;
    this.room["messages"] = [];
  }

  /**
   * PUBLIC
   */

  sendMessage(message) {
    this.ws.send(JSON.stringify({
      action: 'send-message',
      message: message,
      target: {
        id: room.id,
        name: room.name
      }
    }));
  }

  joinRoom(roomName) {
    this.ws.send(JSON.stringify({ action: 'join-room', message: roomName }));
  }

  leaveRoom() {
    this.ws.send(JSON.stringify({ action: 'leave-room', message: this.room.id }));
  }

  joinPrivateRoom() {
    this.ws.send(JSON.stringify({ action: 'join-room-private', message: this.room.id }));
  }
}

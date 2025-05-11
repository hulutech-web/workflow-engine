class SocketClient {
    url = "";
    onReFn = null;
    onErrFn = null;
    onSucFn = null;
  
    isSocketClose = false;
    reconnectCount = 1;
    heartbeatInterval = "";
    socketTask = null;
  
    againTimer = null;
  
    system_id = null;
  
    constructor(socketUrl, onReceive, onErrorEvent, onSuccess, system_id = null) {
      this.url = socketUrl;
      this.onReFn = onReceive;
      this.onErrFn = onErrorEvent;
      this.onSucFn = onSuccess;
      this.isSocketClose = false;
      this.system_id = system_id || null;
  
      if (this.socketTask) {
        this.socketTask.close();
        if (this.heartbeatInterval) {
          clearInterval(this.heartbeatInterval);
        }
      }
  
      this.socketTask = new WebSocket(this.url);
  
      this.socketTask.addEventListener('open', (event) => {
        console.log("WebSocket打开");
        clearInterval(this.againTimer);
        onSuccess({ isShow: false });
        this.heartbeatInterval && clearInterval(this.heartbeatInterval);
        this.heartbeatInterval = setInterval(() => {
          this.sendMsg('PING');
        }, 5000);
      });
  
      this.socketTask.onerror = (err) => {
        console.error('WebSocket连接打开失败，请检查', err);
        clearInterval(this.heartbeatInterval);
        if (!this.isSocketClose) {
          this.reconnect(this.url, this.onErrFn);
        }
      };
  
      this.socketTask.addEventListener('close', (event) => {
        console.error("断开连接", event);
        clearInterval(this.heartbeatInterval);
        if (!this.isSocketClose) {
          this.reconnect(this.url, this.onErrFn);
        }
      });
  
      this.socketTask.addEventListener('message', (event) => {
        let serverData = event.data;
        if (serverData === "PING") {
          console.error("PONG");
          this.socketTask.send("PONG");
        } else {
          serverData && this.onReFn(JSON.parse(serverData));
        }
      });
  
      return this;
    }
  
    reconnect = (url, onErrorEvent) => {
      clearInterval(this.againTimer);
      clearInterval(this.heartbeatInterval);
      this.socketTask.close();
      this.socketTask = null;
  
      onErrorEvent({ isShow: true, message: '重连中，请稍后...' });
  
      if (this.reconnectCount <= 0) {
        alert("已超出重连次数");
        return;
      }
  
      this.againTimer = setInterval(() => {
        let skt = new SocketClient(this.url, this.onReFn, this.onErrFn, this.onSucFn, this.system_id);
        this.reconnectCount--;
        skt.reconnectCount = this.reconnectCount;
        clearInterval(this.againTimer);
        console.log('在重新连接中...');
      }, 5000);
    };
  
    sendMsg = (msg) => {
      try {
        this.socketTask.send(JSON.stringify(msg));
      } catch (e) {
        if (this.isSocketClose) {
          return;
        } else {
          this.reconnect(this.url, this.onErrFn);
        }
      }
    };
  
    stop = () => {
      this.isSocketClose = true;
      clearInterval(this.heartbeatInterval);
      clearInterval(this.againTimer);
      this.socketTask.close();
      this.socketTask = null;
    };
  }
  
  export default SocketClient;
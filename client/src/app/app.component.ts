import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'client';
  ws = new WebSocket("ws://localhost:8080/ws");

  ngOnInit(): void {
    this.connect()
  }

  connect() {
    this.ws.onopen = () => {
      console.log("Successfully connected")
    }

    this.ws.onmessage = (msg) => {
      console.log(msg)
    }

    this.ws.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
    };

    this.ws.onerror = (error) => {
      console.log("Socket Error: ", error);
    };
  }

  sendMessage(msg: string) {
    console.log("sending msg: ", msg);
    this.ws.send(msg);
  }
}

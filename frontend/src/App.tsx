import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { grpc } from "grpc-web-client";
import { BrowserHeaders } from "browser-headers";
import { MealPlannerService } from "./proto/service/service_pb_service"
import {LoginUserRequest, LoginUserResponse} from "./proto/service/service_pb"

function getResponse(){

  const message = new LoginUserRequest();
  message.setPassword("password");
  message.setUsername("chvck");

  grpc.invoke(MealPlannerService.LoginUser ,{
    request : message,
    host : "http://localhost:3001",
    onHeaders: (headers: BrowserHeaders) => {
      console.log("MealPlannerService.onHeaders", headers);
    },
    onMessage: (message: LoginUserResponse) => {
      console.log("Echo.onMessage", message.toObject());
      const elt = document.getElementById("greeting");
      if (elt != null ) {
        elt.innerText = message.toString();
      }
    },
    onEnd: (code: grpc.Code, msg: string, trailers: BrowserHeaders) => {
      console.log("Echo.onEnd", code, msg, trailers);
    }
  })

}
getResponse();

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  }
}

export default App;

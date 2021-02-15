import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ActivatedRoute} from "@angular/router";
import {getLoginURL} from "../app.component";

export interface Player {
  name : string;
  time : number;
}

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent implements OnInit {
  username : string;
  game : string;
  players: Player[];


  constructor(private http: HttpClient, private activatedRoute : ActivatedRoute) { }

  ngOnInit(): void {
    this.username = this.activatedRoute.snapshot.paramMap.get("username");
    this.game = this.activatedRoute.snapshot.paramMap.get("game");
    /*this.refreshPlayers();
    setInterval(() => {
      this.refreshPlayers();
    }, 1000)*/
  }

  refreshPlayers() : void {

    const url : string = getLoginURL() + "&type=update&user=" + this.username + "&game=" + this.game;

    console.log("UPDATING FROM " + url);
    this.http.get(url).subscribe((data : {}) => {
      console.log(data);
      if(data["status"] == 0) {
        this.players = data["players"].map(pdata => {
          var p : Player = {name: pdata["name"], time: pdata["time"]};
          return p;
        })
      }
    })
  }
}

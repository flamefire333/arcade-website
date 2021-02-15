import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ActivatedRoute} from "@angular/router";

//Beans from https://among-us.fandom.com/wiki/Category:Colors
//Hats from https://among-us.fandom.com/wiki/Cosmetics

@Component({
  selector: 'app-fe3hau',
  templateUrl: './fe3hau.component.html',
  styleUrls: ['./fe3hau.component.css']
})
export class Fe3hauComponent implements OnInit {
  name: string;
  game: string;
  characterName: string;
  colorID: number;
  hatID: number;
  startingNewRound = false;
  constructor(private http: HttpClient, private activatedRoute : ActivatedRoute) { }

  ngOnInit(): void {
    this.name = this.activatedRoute.snapshot.paramMap.get("username");
    this.game = this.activatedRoute.snapshot.paramMap.get("game");
    this.refreshPreview();
    setInterval(() => {
      this.refreshPreview();
    }, 5000)
  }

  refreshPreview() : void {

    const url : string = "http://jaketheduck.com/games/phpapi/gamedata.php?type=mapping&name=" + this.name + "&game=" + this.game;
    this.http.get(url).subscribe((data : {}) => {
      if(data["status"] == 0) {
        this.characterName = data["name"];
        this.colorID = data["colorID"];
        this.hatID = data["hatID"];
      }
    })
  }

  startRound() : void {
    this.startingNewRound = true;
    const url : string = "http://jaketheduck.com/games/phpapi/gamedata.php?type=setup&game=fe3hau";
    this.http.get(url).subscribe((data : {}) => {
      if(data["status"] == 0) {
        console.log("New round started successfully");
      }
      this.startingNewRound = false;
    })
  }

}

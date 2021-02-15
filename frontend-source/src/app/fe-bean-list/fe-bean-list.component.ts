import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-fe-bean-list',
  templateUrl: './fe-bean-list.component.html',
  styleUrls: ['./fe-bean-list.component.css']
})
export class FeBeanListComponent implements OnInit {
  characters : string[];
  constructor(private http : HttpClient) { }

  ngOnInit(): void {
    const url: string = "http://jaketheduck.com/games/phpapi/gamedata.php?type=characters";
    this.http.get(url).subscribe((data: {}) => {
      console.log("Data: " + data);
      console.log("Status: " + data["status"]);
      if (data["status"] == 0) {
        this.characters = data["chars"];
      }
    });
  }

}

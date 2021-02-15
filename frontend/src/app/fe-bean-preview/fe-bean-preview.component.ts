import {Component, Input, OnInit} from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-fe-bean-preview',
  templateUrl: './fe-bean-preview.component.html',
  styleUrls: ['./fe-bean-preview.component.css']
})
export class FeBeanPreviewComponent implements OnInit {
  colorID : string;
  hatID : string;
  showPrimary = true;
  @Input() character : string;

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    const url: string = "http://jaketheduck.com/games/phpapi/gamedata.php?type=chardata&character=" + this.character;
    this.http.get(url).subscribe((data: {}) => {
      if (data["status"] == 0) {
        console.log("COLORID: " + data["colorID"]);
        this.colorID = data["colorID"];
        this.hatID = data["hatID"];
      }
    });
  }

}

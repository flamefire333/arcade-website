import {Component, Input, OnChanges, OnInit, SimpleChange, SimpleChanges} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ActivatedRoute} from "@angular/router";
import {MatDivider} from "@angular/material/divider";

@Component({
  selector: 'app-bean-preview',
  templateUrl: './bean-preview.component.html',
  styleUrls: ['./bean-preview.component.css']
})
export class BeanPreviewComponent implements OnInit, OnChanges {
  @Input() colorID : string;
  @Input() hatID : string;
  @Input() scale: number;
  width : number;
  colorSrc: string = "";
  hatSrc: string = "";
  hatdx: number;
  hatdy: number;
  showPrimary: boolean = true;

  constructor(private http: HttpClient, private activatedRoute: ActivatedRoute) {
  }

  ngOnInit(): void {
    //var colorID = this.activatedRoute.snapshot.paramMap.get("color");
    //var hatID = this.activatedRoute.snapshot.paramMap.get("hat");
    this.runWithColorAndHat(this.colorID, this.hatID);
  }

  ngOnChanges(changes: SimpleChanges) : void {
    this.runWithColorAndHat(this.colorID, this.hatID);
  }

  runWithColorAndHat(colorID: string, hatID: string): void {
    if(colorID == null || hatID == null) {
      return;
    }
    console.log("Updating color and hat");
    const url: string = "http://jaketheduck.com/games/phpapi/gamedata.php?type=beandata&color=" + colorID + "&hat=" + hatID;
    this.http.get(url).subscribe((data: {}) => {
      console.log("Data: " + data);
      console.log("Status: " + data["status"]);
      if (data["status"] == 0) {
        console.log("Color src: " + data["colorSrc"]);
        this.colorSrc = data["colorSrc"];
        this.hatSrc = data["hatSrc"];
        this.hatdx = data["hatdx"];
        this.hatdy = data["hatdy"];
        this.width = data["width"];
        console.log(this.hatdx);
        console.log(this.hatdy);
      }
    })
  }

}

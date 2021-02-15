import {Component, Input, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {HttpClient} from "@angular/common/http";
import {TimeInterval} from "rxjs";

export interface LobbyItem {
  name: string;
  icon: string;
  time: number;
  isGrayed: boolean;
  roleName: string;
  roleIcon: string;
  roleDescription: string;
}

export interface LobbySection {
  title: string;
  items: LobbyItem[];
}

@Component({
  selector: 'app-lobby',
  templateUrl: './lobby.component.html',
  styleUrls: ['./lobby.component.css']
})
export class LobbyComponent implements OnInit {
  @Input() sections: LobbySection[];
  @Input() active: string;
  constructor() { }

  ngOnInit(): void {

  }
}

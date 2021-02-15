import { Component } from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'games';
}

function randomInt(): number {
  const random = Math.floor(Math.random() * 1000000000000);
  return random;
}

export function getLoginURL(): string {
  /*const randomVal = randomInt();
  if(location.protocol === 'https:') {
    return 'https://jaketheduck.com/games/phpapi/login.php?randomGen=' + randomVal;
  } else {
    return 'http://jaketheduck.com/games/phpapi/login.php?randomGen=' + randomVal;
  }*/
  return "/api/login";
}

export function getMafiaURL(): string {
  /*const randomVal = randomInt();
  if(location.protocol === 'https:') {
    return 'https://jaketheduck.com/games/phpapi/mafia.php?randomGen=' + randomVal;
  } else {
    return 'http://jaketheduck.com/games/phpapi/mafia.php?randomGen=' + randomVal;
  }*/
  return "/api/mafia"
}

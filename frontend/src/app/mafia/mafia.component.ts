import {Component, Input, OnInit} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {MatTooltip} from '@angular/material/tooltip';
import {LobbyItem, LobbySection} from '../lobby/lobby.component';
import {Player} from '../game/game.component';
import {AppComponent, getMafiaURL} from '../app.component';
import {Observable, Subject, Subscription} from 'rxjs';
import {catchError, debounceTime, switchMap} from 'rxjs/operators';
import {empty} from 'rxjs';
import {FormControl, FormGroup} from "@angular/forms";

interface Role {
  id: number;
  name: string;
  team: number;
  description: string;
  icon: string;
}

interface Character {
  name: string;
  avatar: string;
  alive: boolean;
  roleID: number;
  id: number;
}

interface VoteField {
  type: string;
  options: string[];
  barrierID: number;
}

interface Vote {
  name: string;
  vote: string;
}

interface VoteContainer {
  title: string;
  list: Vote[];
  fields: VoteField[];
}

@Component({
  selector: 'app-mafia',
  templateUrl: './mafia.component.html',
  styleUrls: ['./mafia.component.scss']
})
export class MafiaComponent implements OnInit {
  @Input() username: string;
  userID = 0;
  isLeader: boolean;
  players: Player[];
  started = false;
  phase = 0;
  isRefreshingStatus = false;
  isRefreshingMain = false;
  roles: Role[];
  selectedRoles: Map<Role, number>;
  roleCounts: { role: Role, count: number }[];
  hoveredRoles: Map<string, boolean>;
  characters: Character[];
  myCharacter: string;
  myCharacterAvatar: string;
  lobbySections: LobbySection[];
  aliveCharacterNames: string[];
  isAlive: boolean;
  /*voteChoice: string;
  votes: { name: string, vote: string }[];
  votingTitle: string;
  voteBarrierID: number;*/
  voteContainers: VoteContainer[];
  isDay = true;
  dayCount = 1;
  activeRoles: { role: Role, amount: number }[];
  willSubject: Subject<string>;
  myRole = -1;
  characterGroups: { id: number, name: string }[];
  characterGroupChoice = 1;
  lastCheckedPhase = -1;
  votingFormGroup: FormGroup;

  constructor(private http: HttpClient) {
  }

  ngOnInit(): void {
    console.log('Mafia Version 0.1');
    this.voteContainers = [];
    this.votingFormGroup = new FormGroup({});
    this.roleCounts = [];
    this.activeRoles = [];
    this.characterGroups = [];
    this.aliveCharacterNames = [];
    this.isAlive = false;
    this.selectedRoles = new Map<Role, number>();
    this.hoveredRoles = new Map<string, boolean>();
    this.lobbySections = [];
    this.roles = [];
    this.willSubject = new Subject<string>();
    this.willSubject.pipe(debounceTime(5000), switchMap(value => this.updateWill(value))).pipe(catchError(() => empty())).subscribe(() => {
      console.log('WILL UPDATED');
    });
    this.refreshStatus();
    setInterval(() => {
      this.refreshStatus();
    }, 2000);
    const roleURL: string = getMafiaURL() + '/roles';
    this.http.get(roleURL).subscribe((data: Role[]) => {
      this.roles = data;
    });
    const characterGroupURL: string = getMafiaURL() + '/character/groups';
    this.http.get(characterGroupURL).subscribe((data: { info: { id: number, name: string }[] }) => {
      console.log('CHARACTER GROUPS: ' + data);
      this.characterGroups = data.info;
    });
    this.isLeader = false;
  }

  addRole(role: Role): void {
    if (this.selectedRoles.has(role)) {
      this.selectedRoles.set(role, this.selectedRoles.get(role) + 1);
    } else {
      this.selectedRoles.set(role, 1);
    }
    this.roleCounts = this.getSelectedRoleCounts();
  }

  removeRole(role: Role): void {
    if (this.selectedRoles.has(role)) {
      this.selectedRoles.set(role, this.selectedRoles.get(role) - 1);
      if (this.selectedRoles.get(role) <= 0) {
        this.selectedRoles.delete(role);
      }
    }
    this.roleCounts = this.getSelectedRoleCounts();
  }

  getSelectedRoleCounts(): { role: Role, count: number }[] {
    const data = [];
    for (const key of this.selectedRoles.keys()) {
      data.push({role: key, count: this.selectedRoles.get(key)});
    }
    return data;
  }

  refreshStatus(): void {
    interface StatusResponse {
      status: number;
      info: {
        lobbyData: Player[];
        gameStatus: {
          started: boolean;
          phase: number;
          day: boolean;
          dayCount: number;
        };
        activeRoles: { name: string, amount: number }[];
        characters: Character[];
        myCharacter: string;
        votingData: VoteContainer[];
      };
    }

    console.log('REFRESH STATUS');
    this.refreshLobbyItems();
    if (!this.isRefreshingStatus) {
      this.isRefreshingStatus = true;
      const url: string = getMafiaURL() + '/status/' + this.username;
      console.log('REFRSH STATUS URL');
      this.http.get(url).subscribe((data: StatusResponse) => {
        console.log('Status Data: ');
        console.log(data);
        if (!!data && data.status === 0) {
          this.players = data.info.lobbyData;
          this.started = data.info.gameStatus.started;
          this.phase = data.info.gameStatus.phase;
          this.isDay = data.info.gameStatus.day;
          this.dayCount = data.info.gameStatus.dayCount;
          if (this.started) {
            if (this.lastCheckedPhase !== this.phase && !!this.roles && this.roles.length > 0) {
              this.activeRoles = data.info.activeRoles.map(rdata => {
                const relatedRole = this.roles.find(role => role.name === rdata.name);
                return {
                  role: relatedRole,
                  amount: rdata.amount
                };
              });
              this.lastCheckedPhase = this.phase;
              // this.voteChoice = null;
            }
            this.characters = data.info.characters;
            this.aliveCharacterNames = this.characters.filter(char => char.alive).map(char => char.name);
            this.myCharacter = data.info.myCharacter;
            this.myRole = this.characters.find(char => char.name === this.myCharacter)?.roleID;
            this.isAlive = this.characters.find(char => char.name === this.myCharacter)?.alive;
            this.myCharacterAvatar = this.characters.find(char => char.name === this.myCharacter)?.avatar;
            if (!!data.info.votingData) {
              data.info.votingData.forEach(vd => {
                vd.fields.forEach(field => {
                  const fieldName = this.getFormControlName(field);
                  if (!this.votingFormGroup.contains(fieldName)) {
                    this.votingFormGroup.addControl(fieldName, new FormControl(''));
                  }
                });
              });
              this.voteContainers = data.info.votingData;
            } else {
              this.voteContainers = [];
            }
            /*if (!!data.info.votingData && data.info.votingData.length > 0) {
              this.votes = data.info.votingData[0].list;
              this.votingTitle = data.info.votingData[0].title;
              this.voteBarrierID = data.info.votingData[0].id;
            } else {
              this.votes = [];
              this.votingTitle = '';
              this.voteBarrierID = 0;
            }*/
          } else {
            this.activeRoles = [];
          }
        }
        this.isRefreshingStatus = false;
      });
    }
    this.isLeader = !!this.players && this.players.length > 0 && this.username === this.players[0].name;
  }

  getCharacterIDFromCharacterName(name: string): number {
    return this.characters.find(char => char.name === name)?.id;
  }

  getLobbyRoleInfo(roleID: number): Role {
    return this.roles.find(role => role.id === roleID);
  }

  refreshLobbyItems(): void {
    if (this.started) {
      const aliveLobbyItems = this.characters.filter(char => char.alive).map(character => {
        const roleInfo = this.getLobbyRoleInfo(character.roleID);
        return {
          name: character.name,
          time: 0,
          icon: character.avatar,
          isGrayed: false,
          roleIcon: roleInfo?.icon,
          roleName: roleInfo?.name,
          roleDescription: roleInfo?.description
        };
      });
      const deadLobbyItems = this.characters.filter(char => !char.alive).map(char => {
        const roleInfo = this.getLobbyRoleInfo(char.roleID);
        return {
          name: char.name,
          time: 0,
          icon: char.avatar,
          isGrayed: true,
          roleIcon: roleInfo?.icon,
          roleName: roleInfo?.name,
          roleDescription: roleInfo?.description
        };
      });
      const aliveLobby: LobbySection = {
        title: 'Players',
        items: aliveLobbyItems
      };
      const deadLobby: LobbySection = {
        title: 'Graveyard',
        items: deadLobbyItems
      };
      this.lobbySections = [aliveLobby, deadLobby];
    } else if (!!this.players) {
      this.lobbySections = [{
        title: 'Players',
        items:
          this.players.map(player => {
            const item: LobbyItem = {
              name: player.name,
              time: player.time,
              icon: null,
              isGrayed: false,
              roleIcon: null,
              roleName: null,
              roleDescription: null
            };
            return item;
          })
      }];
    } else {
      this.lobbySections = [];
    }
  }

  getChatName(): string {
    if (!!this.myCharacter && this.started) {
      return this.myCharacter;
    } else {
      return this.username;
    }
  }

  getChatAvatar(): string {
    if (!!this.myCharacterAvatar && this.started) {
      return this.myCharacterAvatar;
    } else {
      return '';
    }
  }

  startGame(): void {
    let url: string = getMafiaURL() + '/setup?group=' + this.characterGroupChoice;
    this.roles.forEach(role => {
      let roleCount = this.selectedRoles.get(role);
      roleCount = !!roleCount ? roleCount : 0;
      url += '&' + role.name + '=' + roleCount;
    });
    this.http.get(url).subscribe((data: {}) => {
    });
  }

  canStartGame(): boolean {
    let roles = 0;
    for (const roleCount of this.roleCounts) {
      roles += roleCount.count;
    }
    return roles === this.players.length;
  }

  sendDayVote(voteName: string, containerID: number): void {
    const url: string = getMafiaURL() + '/vote/' + this.username + '/' + containerID + '/' + voteName;
    this.http.get(url).subscribe((data: {}) => {
      console.log('Vote sent');
    });
  }

  sendVoteTextEvent(field: VoteField): void {
    const containerID = field.barrierID;
    const voteName = this.votingFormGroup.get(this.getFormControlName(field)).value;
    if(voteName.length >= 4) {
      const url: string = getMafiaURL() + '/vote/' + this.username + '/' + containerID + '/' + voteName;
      this.http.get(url).subscribe((data: {}) => {
        console.log('Vote sent');
      });
    }
  }

  updateWill(will: string): Observable<any> {
    const url: string = getMafiaURL() + '/will';
    const body = {name: this.username, will};
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    };
    return this.http.post(url, body, httpOptions);
  }

  getFormControlName(field: VoteField): string {
    return 'vote-barrier-field-' + field.barrierID;
  }

  voteContainerTrackBy(index: number, item: VoteContainer): number {
    return item.fields[0].barrierID;
  }

  voteFieldTrackBy(index: number, item: VoteField): number {
    return item.barrierID;
  }
}

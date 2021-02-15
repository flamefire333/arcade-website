import {
  Component,
  ElementRef,
  Input,
  OnChanges,
  OnInit,
  QueryList,
  SimpleChanges,
  ViewChild,
  ViewChildren
} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {getMafiaURL} from "../app.component";
import {Subject} from "rxjs";

interface ChatData {
  sender: string,
  avatar: string,
  message: string,
  type: number;
  phase: number;
}

interface ChatMessageData {
  code: string
  image?: string
  isEmote: boolean
}

interface ChatMessageDataHolder {
  sender: string;
  avatar: string;
  type: number;
  phase: number;
  data: ChatMessageData[];
}

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnInit {
  messages: ChatMessageDataHolder[];
  emoteMap: Map<string, string>;
  isRefreshingChat = false;
  chatStartID = 0
  chatPhase = 0
  chatSubject: Subject<any>;
  maxPhase: number = 0
  @Input() phase: number
  @Input() isDead: boolean;
  @Input() username: string
  @Input() userID: number
  @Input() characterAvatar: string
  @Input() characterName: string
  @Input() canSend: boolean
  @ViewChild('chatEntry') chatEntry : ElementRef;
  @ViewChild('chatTableHolder') chatTableHolder : ElementRef;
  @ViewChild('chatTable') chatTable : ElementRef;

  constructor(private http: HttpClient) {
  }

  ngOnInit(): void {
    this.emoteMap = new Map<string, string>();
    this.emoteMap.set('crybichu', 'https://cdn.discordapp.com/emojis/622202821940609065.png?v=1');
    this.emoteMap.set('honk', 'https://cdn.discordapp.com/emojis/633488458022780939.png?v=1');
    this.emoteMap.set('salutetri', 'https://cdn.discordapp.com/emojis/603325799948484609.png?v=1');
    this.emoteMap.set('potatolenz', 'https://cdn.discordapp.com/emojis/609861969352196101.png?v=1');
    this.emoteMap.set('furbaby', 'https://cdn.discordapp.com/emojis/695258347124817941.png?v=1');
    this.emoteMap.set('tpose', 'https://cdn.discordapp.com/emojis/624458657010024449.png?v=1');
    this.emoteMap.set('woolbaby', 'https://cdn.discordapp.com/emojis/757515237824921760.png?v=1');
    this.emoteMap.set('begone', 'https://cdn.discordapp.com/emojis/653691327715868682.png?v=1');
    this.messages = [];
    this.chatSubject = new Subject<any>();
    this.chatSubject.subscribe(data => this.parseChatData(data));
    this.refreshChat();
    setInterval(() => {
      this.refreshChat();
    }, 2000)
  }

  convertMessage(chatData: ChatData): ChatMessageDataHolder {
    let data = [];
    let currentMessage = '';
    let couldBeEmote = false;
    chatData.message.split('').forEach(char => {
      if(char === ':') {
        if(couldBeEmote) {
          let emote = this.convertToEmote(currentMessage);
          if(!!emote) {
            data.push({code: currentMessage, image: emote, isEmote: true});
            currentMessage = '';
            couldBeEmote = false;
          } else {
            data.push({code: ':' + currentMessage, isEmote: false});
            currentMessage = ''
          }
        } else {
          data.push({code: currentMessage, isEmote: false});
          currentMessage = '';
          couldBeEmote = true;
        }
      } else {
        currentMessage = currentMessage + char;
      }
    });
    if(couldBeEmote) {
      data.push({code: ':' + currentMessage, isEmote: false});
    } else {
      data.push({code: currentMessage, isEmote: false});
    }
    return {
      sender: chatData.sender,
      avatar: chatData.avatar,
      type: chatData.type,
      phase: chatData.phase,
      data
    };
  }

  convertToEmote(emoteName: string): (string | undefined) {
    return this.emoteMap.get(emoteName);
  }

  getMessageType(): number {
    if(this.isDead) {
      return -1;
    } else {
      return 0;
    }
  }

  ChatOnKey(event: KeyboardEvent) {
    if(event.code === 'Enter' && this.canSend) {
      const url: string = getMafiaURL() + "/chat/send"
      /*const body = new URLSearchParams();
      body.set("startID", String(this.chatStartID))
      body.set("phase", String(this.phase))
      body.set("avatar", this.characterAvatar);
      body.set("user_name", this.username);
      body.set("character_name", this.characterName);
      body.set("message", this.chatEntry.nativeElement.value);*/
      var phaseToUse = this.chatPhase
      var startID = this.chatStartID

      const body = {startID: startID, phase: phaseToUse, avatar: this.characterAvatar, user_name: this.username, character_name: this.characterName, message: this.chatEntry.nativeElement.value}
      this.chatEntry.nativeElement.value = '';
      const httpOptions = {
        headers: new HttpHeaders({
          'Content-Type':  'application/json',
        })
      };
      this.http.post(url, body, httpOptions).subscribe((data: {}) => {
        console.log("Post Data: " + data["status"] + " " + data["query"]);
        this.chatSubject.next(data);
      });
    }
  }

  refreshChat() : void {
    const startPhase = this.phase;
    if(startPhase == this.phase && !this.isRefreshingChat) {
      this.isRefreshingChat = true;
      /*const url: string = getMafiaURL() + "&action=read&phase=" + startPhase + "&messageType=" + this.getMessageType() +
        "&startID=" + this.chatStartID + "&isDead=" + (this.isDead ? 1 : 0) + "&name=" + this.username;*/
      if (startPhase > this.maxPhase) {
        this.chatStartID = 0
        this.chatPhase = startPhase
        this.maxPhase = this.chatPhase
      }
      const url: string = getMafiaURL() + "/chat/read/" + this.username + "/" + startPhase + "/" + this.chatStartID
      this.http.get(url).subscribe(data => {
        this.chatSubject.next(data);
        this.isRefreshingChat = false;
      });
    }
  }

  parseChatData(data: any) {
      if (data["status"] == 0) {
        const oldStartID = this.chatStartID;
        data["chat"].forEach(cdata => {
          let mid: number = +cdata["id"];
          if(mid + 1 > this.chatStartID) {
            this.chatStartID = mid + 1;
          }
        })
        console.log(data)
        let newChats : ChatMessageDataHolder[] = data["chat"].filter(cdata => cdata["id"] >= oldStartID).map(cdata => {
          const c: ChatData = {sender: cdata["name"], avatar: cdata["avatar"], message: cdata["message"], type: cdata["messageType"], phase: cdata["phase"]};
          const m: ChatMessageDataHolder = this.convertMessage(c);
          return m;
        }).reverse();
        console.log(newChats)
        this.messages = this.messages.filter(message => message.phase == this.phase);
        if(newChats.length > 0) {
          this.messages = newChats.concat(this.messages);
        }
        console.log(this.messages)
        console.log(this.phase)
      }
  }
}

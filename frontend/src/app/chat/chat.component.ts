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
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {getMafiaURL} from '../app.component';
import {Subject} from 'rxjs';

interface ChatData {
  sender: string;
  avatar: string;
  message: string;
  type: number;
  phase: number;
}

interface ChatMessageData {
  code: string;
  image?: string;
  isEmote: boolean;
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
  emotes: string[];
  defaultAvatar: string;
  filteredEmotes: string[];
  isRefreshingChat = false;
  chatStartID = 0;
  chatPhase = 0;
  chatSubject: Subject<any>;
  maxPhase = 0;
  @Input() phase: number;
  @Input() isDead: boolean;
  @Input() username: string;
  @Input() userID: number;
  @Input() characterAvatar: string;
  @Input() characterName: string;
  @Input() canSend: boolean;
  @ViewChild('chatEntry') chatEntry: ElementRef;
  @ViewChild('chatTableHolder') chatTableHolder: ElementRef;
  @ViewChild('chatTable') chatTable: ElementRef;

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
    this.emoteMap.set('biff', 'https://cdn.discordapp.com/emojis/801632481979138058.png?v=1');
    this.emoteMap.set('BigF', 'https://cdn.discordapp.com/emojis/641062063002746891.png?v=1');
    this.emoteMap.set('bonk', 'https://cdn.discordapp.com/emojis/657725330466799626.png?v=1');
    this.emoteMap.set('breadsive', 'https://cdn.discordapp.com/emojis/700748202931912814.png?v=1');
    this.emoteMap.set('bygun', 'https://cdn.discordapp.com/emojis/624458153131769878.png?v=1');
    this.emoteMap.set('chefskiss', 'https://cdn.discordapp.com/emojis/614981976390238219.png?v=1');
    this.emoteMap.set('BiologyNotes', 'https://cdn.discordapp.com/emojis/811279742745313300.png?v=1');
    this.emotes = [];
    for (const emote of this.emoteMap.keys()) {
      this.emotes.push(emote);
    }
    this.defaultAvatar = this.emoteMap.get(this.emotes[Math.floor(Math.random() * this.emotes.length)]);
    this.tryToFindDefaultAvatar();
    this.emotes = this.emotes.sort((a, b) => a.localeCompare(b));
    this.filteredEmotes = [];
    this.messages = [];
    this.chatSubject = new Subject<any>();
    this.chatSubject.subscribe(data => this.parseChatData(data));
    this.refreshChat();
    setInterval(() => {
      this.refreshChat();
    }, 2000);
  }

  convertMessage(chatData: ChatData): ChatMessageDataHolder {
    const data = [];
    let currentMessage = '';
    let couldBeEmote = false;
    chatData.message.split('').forEach(char => {
      if (char === ':') {
        if (couldBeEmote) {
          const emote = this.convertToEmote(currentMessage);
          if (!!emote) {
            data.push({code: currentMessage, image: emote, isEmote: true});
            currentMessage = '';
            couldBeEmote = false;
          } else {
            data.push({code: ':' + currentMessage, isEmote: false});
            currentMessage = '';
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
    if (couldBeEmote) {
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

  getChatAvatar(): string {
    if (this.characterAvatar !== '') {
      return this.characterAvatar;
    } else {
      return this.defaultAvatar;
    }
  }

  convertToEmote(emoteName: string): (string | undefined) {
    return this.emoteMap.get(emoteName);
  }

  getMessageType(): number {
    if (this.isDead) {
      return -1;
    } else {
      return 0;
    }
  }

  getAutocompletedMessage(emote: string): string {
    const currentText: string = this.chatEntry.nativeElement.value;
    const words = currentText.split(' ');
    if (words.length > 0) {
      const caretPos = this.chatEntry.nativeElement.selectionStart;
      words[this.getActiveWordIndex(words, caretPos)] = ':' + emote + ':';
      return words.join(' ');
    } else {
      return '';
    }
  }

  getActiveWordIndex(words: string[], caretPos: number): number {
    let prevLength = 0;
    let wordIndex = 0;
    for (const word of words) {
      const wordLength = word.length;
      if (caretPos <= prevLength + wordLength) {
        return wordIndex;
      }
      prevLength += wordLength + 1;
      wordIndex++;
    }
    return words.length - 1;
  }

  ChatOnKey(event: KeyboardEvent): void {
    if (event.code === 'Enter' && this.canSend && this.filteredEmotes.length === 0) {
      const url: string = getMafiaURL() + '/chat/send';
      /*const body = new URLSearchParams();
      body.set("startID", String(this.chatStartID))
      body.set("phase", String(this.phase))
      body.set("avatar", this.characterAvatar);
      body.set("user_name", this.username);
      body.set("character_name", this.characterName);
      body.set("message", this.chatEntry.nativeElement.value);*/
      const phaseToUse = this.chatPhase;
      const startID = this.chatStartID;

      const body = {
        startID,
        phase: phaseToUse,
        avatar: this.getChatAvatar(),
        user_name: this.username,
        character_name: this.characterName,
        message: this.chatEntry.nativeElement.value
      };
      this.chatEntry.nativeElement.value = '';
      const httpOptions = {
        headers: new HttpHeaders({
          'Content-Type': 'application/json',
        })
      };
      this.http.post(url, body, httpOptions).subscribe((data: { status: number, query: string }) => {
        console.log('Post Data: ' + data.status + ' ' + data.query);
        this.chatSubject.next(data);
      });
    }
    const currentText: string = this.chatEntry.nativeElement.value;
    const words = currentText.split(' ');
    if (words.length > 0) {
      const caretPos = this.chatEntry.nativeElement.selectionStart;
      let currentWord = words[this.getActiveWordIndex(words, caretPos)];
      if (currentWord.startsWith(':')) {
        currentWord = currentWord.substr(1);
        this.filteredEmotes = this.emotes.filter(emote => emote.toLowerCase().startsWith(currentWord.toLowerCase()));
      } else {
        this.filteredEmotes = [];
      }
    } else {
      this.filteredEmotes = [];
    }
  }

  refreshChat(): void {
    const startPhase = this.phase;
    if (startPhase === this.phase && !this.isRefreshingChat) {
      this.isRefreshingChat = true;
      /*const url: string = getMafiaURL() + "&action=read&phase=" + startPhase + "&messageType=" + this.getMessageType() +
        "&startID=" + this.chatStartID + "&isDead=" + (this.isDead ? 1 : 0) + "&name=" + this.username;*/
      if (startPhase > this.maxPhase) {
        this.chatStartID = 0;
        this.chatPhase = startPhase;
        this.maxPhase = this.chatPhase;
      }
      const url: string = getMafiaURL() + '/chat/read/' + this.username + '/' + startPhase + '/' + this.chatStartID;
      this.http.get(url).subscribe(data => {
        this.chatSubject.next(data);
        this.isRefreshingChat = false;
      });
    }
  }

  parseChatData(data: any): void {
    if (data.status === 0) {
      const oldStartID = this.chatStartID;
      data.chat.forEach(cdata => {
        const mid: number = +cdata.id;
        if (mid + 1 > this.chatStartID) {
          this.chatStartID = mid + 1;
        }
      });
      console.log(data);
      const newChats: ChatMessageDataHolder[] = data.chat.filter(cdata => cdata.id >= oldStartID).map(cdata => {
        const c: ChatData = {
          sender: cdata.name,
          avatar: cdata.avatar,
          message: cdata.message,
          type: cdata.messageType,
          phase: cdata.phase
        };
        const m: ChatMessageDataHolder = this.convertMessage(c);
        return m;
      }).reverse();
      console.log(newChats);
      this.messages = this.messages.filter(message => message.phase === this.phase);
      if (newChats.length > 0) {
        this.messages = newChats.concat(this.messages);
      }
      console.log(this.messages);
      console.log(this.phase);
    }
  }

  tryToFindDefaultAvatar(): void {
    const name = this.characterName;
    const options: ([string[], string])[] = [
      [['jake'], 'https://cdn.discordapp.com/attachments/738860771731308609/820533784918097971/unknown.png'],
      [['j.*lian'], 'https://media.discordapp.net/attachments/738860771731308609/820534137806127104/unknown.png'],
      [['zy.*'], 'https://cdn.discordapp.com/attachments/738860771731308609/820534373156651038/344e722ea7f880fad5b0bc1d2855615c.png'],
      [['sh.*ve.*', 'soup.*', 'wooloo.*'], 'https://media.discordapp.net/attachments/738860771731308609/813204242181259264/unknown.png'],
      [['dino.*', 'cow.*'], 'https://cdn.discordapp.com/attachments/738860771731308609/820534518297001994/5304162d8e15b5f3865eda80dbab8495.png'],
      [['may.*', '.*union.*'], 'https://media.discordapp.net/attachments/532394137970933768/820512946085822494/poke_RedYellowBlueGreen.jpg'],
      [['beep.*'], 'https://media.discordapp.net/attachments/738860771731308609/820534639634153482/3dd6ab13659e58c46296636248a61ebf.png'],
      [['quag.*'], 'https://cdn.discordapp.com/attachments/738860771731308609/820512895569231902/4c9a49e2e9dd4b81a122f765a54c97b7.png'],
      [['nik.*'], 'https://media.discordapp.net/attachments/738860771731308609/820534819011821568/e2bab855d641dfcb2f70275041f3c19a.png'],
      [['lio.*'], 'https://media.discordapp.net/attachments/738860771731308609/820534983847313438/a_05915a54b6c6a08811aac838c551713b.png'],
      [['saj.*'], 'https://cdn.discordapp.com/attachments/738860771731308609/820535280926326784/4abad0b6fdcb8166e4b189bf60ab6c61.png'],
      [['.*hsc.*', 'tae.*', '.*ide', '.*ene', '.*ase', '.*ine', '.*ylene'], 'https://media.discordapp.net/attachments/738860771731308609/820535396458692668/4fb577085d18410eff4dd8d59240ca43.png'],
      [['.*mika.*'], 'https://cdn.discordapp.com/attachments/738860771731308609/813206571029889034/unknown.png'],
      [['.*rj.*'], 'https://media.discordapp.net/attachments/738860771731308609/813207754139369522/unknown.png'],
      [['cl.*', '.*atsumu.*'], 'https://media.discordapp.net/attachments/738860771731308609/820535593301835786/1f8593ae86120a1cea9daf4496b52829.png'],
      [['.*jade.*', '.*hockey.*'], 'https://media.discordapp.net/attachments/738860771731308609/820535723534581780/05bc9fe91b17b7d5624621afc4757461.png'],
      [['.*earthy.*'], 'https://media.discordapp.net/attachments/738860771731308609/813208742820446268/unknown.png'],
      [['.*kris.*'], 'https://cdn.discordapp.com/attachments/738860771731308609/820536334845476894/80356e726fb6bf0be4917c669ed423d3.png'],
      [['.*lance.*', '.*hose.*'], 'https://media.discordapp.net/attachments/738860771731308609/820576990490918952/5we-GTac_400x400.png'],
      [['.*zhou.*'], 'https://media.discordapp.net/attachments/738860771731308609/820536478739464242/6b9473bb4c378e5562f8cbb5fc9e6ed1.png']
    ];
    options.forEach(option => {
      option[0].forEach(possibleName => {
        if (name.toLowerCase().match(possibleName) !== null) {
          this.defaultAvatar = option[1];
        }
      });
    });
  }
}

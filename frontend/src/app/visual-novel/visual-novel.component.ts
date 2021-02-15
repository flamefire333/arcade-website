import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';

@Component({
  selector: 'app-visual-novel',
  templateUrl: './visual-novel.component.html',
  styleUrls: ['./visual-novel.component.css']
})
export class VisualNovelComponent implements OnInit {

  constructor() { }

  targetText: string;
  currentLength: number;
  breaks : number[];
  currLineIndex : number;
  @ViewChild('widthChecker') widthChecker: ElementRef;

  ngOnInit(): void {
    this.breaks = [0, 0, 0, 0, 0]
    this.currLineIndex = 0;
    this.targetText = "Do you believe in magic? Oh never believe its not so for if the magic you believe in can not be more than cheese than who am I to say that cheese is not but a reality."
    this.currentLength = 0;
    setInterval(() => {
      this.breaks[this.currLineIndex + 1] += 1;
      if(this.widthChecker.nativeElement.offsetWidth > 400) {
        this.currLineIndex += 1;
        this.breaks[this.currLineIndex + 1] = this.breaks[this.currLineIndex];
      }
    }, 50)
  }

}

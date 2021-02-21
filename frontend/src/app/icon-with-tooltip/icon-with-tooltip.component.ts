import {Component, ElementRef, Input, OnInit, ViewChild} from '@angular/core';

@Component({
  selector: 'app-icon-with-tooltip',
  templateUrl: './icon-with-tooltip.component.html',
  styleUrls: ['./icon-with-tooltip.component.css']
})
export class IconWithTooltipComponent implements OnInit {
  @Input() name: string;
  @Input() icon: string;
  @Input() description: string;
  @ViewChild('iconRef') iconRef: ElementRef;
  shouldShow = false;

  constructor() {
  }

  ngOnInit(): void {

  }

  isOnRight(): boolean {
    const x = this.iconRef.nativeElement.getBoundingClientRect().left;
    const width = window.innerWidth;
    return x * 2 > width;
  }

}

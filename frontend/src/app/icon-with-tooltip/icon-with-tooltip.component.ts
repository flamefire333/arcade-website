import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-icon-with-tooltip',
  templateUrl: './icon-with-tooltip.component.html',
  styleUrls: ['./icon-with-tooltip.component.css']
})
export class IconWithTooltipComponent implements OnInit {
  @Input() name: string;
  @Input() icon: string;
  @Input() description: string;
  shouldShow: boolean = false;
  constructor() { }

  ngOnInit(): void {
  }

}

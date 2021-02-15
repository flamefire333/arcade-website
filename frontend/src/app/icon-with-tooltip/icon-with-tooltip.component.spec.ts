import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconWithTooltipComponent } from './icon-with-tooltip.component';

describe('IconWithTooltipComponent', () => {
  let component: IconWithTooltipComponent;
  let fixture: ComponentFixture<IconWithTooltipComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IconWithTooltipComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IconWithTooltipComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

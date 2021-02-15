import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Fe3hauComponent } from './fe3hau.component';

describe('Fe3hauComponent', () => {
  let component: Fe3hauComponent;
  let fixture: ComponentFixture<Fe3hauComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ Fe3hauComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(Fe3hauComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

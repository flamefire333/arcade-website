import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FeBeanListComponent } from './fe-bean-list.component';

describe('FeBeanListComponent', () => {
  let component: FeBeanListComponent;
  let fixture: ComponentFixture<FeBeanListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FeBeanListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FeBeanListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

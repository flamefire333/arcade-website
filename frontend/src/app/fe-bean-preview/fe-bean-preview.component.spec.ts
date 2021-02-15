import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FeBeanPreviewComponent } from './fe-bean-preview.component';

describe('FeBeanPreviewComponent', () => {
  let component: FeBeanPreviewComponent;
  let fixture: ComponentFixture<FeBeanPreviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FeBeanPreviewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FeBeanPreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

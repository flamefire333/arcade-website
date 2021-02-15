import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BeanPreviewComponent } from './bean-preview.component';

describe('BeanPreviewComponent', () => {
  let component: BeanPreviewComponent;
  let fixture: ComponentFixture<BeanPreviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BeanPreviewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BeanPreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

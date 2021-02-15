import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiscordMappingsAddComponent } from './discord-mappings-add.component';

describe('DiscordMappingsAddComponent', () => {
  let component: DiscordMappingsAddComponent;
  let fixture: ComponentFixture<DiscordMappingsAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DiscordMappingsAddComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DiscordMappingsAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

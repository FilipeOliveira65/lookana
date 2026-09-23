import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Monitors } from './monitors';

describe('Monitors', () => {
  let component: Monitors;
  let fixture: ComponentFixture<Monitors>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Monitors],
    }).compileComponents();

    fixture = TestBed.createComponent(Monitors);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

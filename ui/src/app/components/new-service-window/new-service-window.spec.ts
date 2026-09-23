import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewServiceWindow } from './new-service-window';

describe('NewServiceWindow', () => {
  let component: NewServiceWindow;
  let fixture: ComponentFixture<NewServiceWindow>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NewServiceWindow],
    }).compileComponents();

    fixture = TestBed.createComponent(NewServiceWindow);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

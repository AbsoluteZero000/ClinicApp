import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PatientSlotEditComponent } from './patient-slot-add.component';

describe('PatientSlotEditComponent', () => {
  let component: PatientSlotEditComponent;
  let fixture: ComponentFixture<PatientSlotEditComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PatientSlotEditComponent]
    });
    fixture = TestBed.createComponent(PatientSlotEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

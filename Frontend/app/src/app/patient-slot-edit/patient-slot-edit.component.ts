import { Component, OnInit } from '@angular/core';
import { SlotsService } from '../services/slots.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { FormBuilder, FormGroup } from '@angular/forms';
import { SharedService } from '../shared.service';

export interface Slot {
  id: string;
  Date: string;
  DoctorName: string;
}


@Component({
  selector: 'app-patient-slot-edit',
  templateUrl: './patient-slot-edit.component.html',
  styleUrls: ['./patient-slot-edit.component.scss'],
})
export class PatientSlotEditComponent {
  displayedColumns: string[] = ['id', 'Date', 'DoctorName'];
  dataStore!: MatTableDataSource<Slot>;
  editForm: FormGroup;

  constructor(
    private _formBuilder: FormBuilder,
    private slotService: SlotsService,
    private sharedService :SharedService
    ) {
      this.editForm = this._formBuilder.group({
      slot: '',
    });
  }
  onFormSubmit() {

    this.slotService
      .AddPatientSlot(this.editForm.value).subscribe({
        next: (res) => {
          alert('slot added ');
          this.getSlotList();
          this.sharedService.triggerDataUpdate();
        },
        error: console.log,
      })

  }
  ngOnInit(): void {
    this.getSlotList();
  }

  private getSlotList(): void {
    this.slotService.getFreeSlots().subscribe({
      next: (res) => {
        const response = res as {
          status: number;
          message: string;
          Data?: Slot[];
        };

        if (response.Data) {
          this.dataStore = new MatTableDataSource(response.Data);
        } else {
          console.error('Invalid server response structure');
        }
      },
      error: console.error,
    });
  }
}

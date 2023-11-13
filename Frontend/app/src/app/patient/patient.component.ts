import { Component, OnInit } from '@angular/core';
import { SlotsService } from '../services/slots.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { PatientSlotEditComponent } from '../patient-slot-edit/patient-slot-edit.component';
import { SharedService } from '../shared.service';

export interface Slot {
  id: string;
  Date: string;
  DoctorName: string;
}

@Component({
  selector: 'app-patient',
  templateUrl: './patient.component.html',
  styleUrls: ['./patient.component.scss'],
})
export class PatientComponent {
  displayedColumns: string[] = ['id', 'Date', 'DoctorName', 'action'];
  dataStore!: MatTableDataSource<Slot>;

  constructor(private _dialog:MatDialog, private slotService: SlotsService, private sharedService: SharedService) {}

  ngOnInit(): void {
    this.getSlotList();

    this.sharedService.dataUpdated$.subscribe(() => {
      this.getSlotList();
    });

  }

  navigateToPatientSlotEdit() {
    this._dialog.open(PatientSlotEditComponent);
  }

  private getSlotList(): void {
    this.slotService.GetPatientSlots().subscribe({
      next: (res) => {
        const response = res as {
          status: number;
          message: string;
          Data?: Slot[];
        };
        console.log(response);
        if (response.Data) {
          this.dataStore = new MatTableDataSource(response.Data);
        } else {
          console.error('Invalid server response structure');
        }
      },
      error: console.error,
    });
  }
  deletePatientSlot(id: string) {
    this.slotService.DeleteSlots(id).subscribe({
      next: (res) => {
        alert('slot deleted ');
        this.getSlotList();
      },
      error: console.log,
    });
  }
}

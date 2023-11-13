import { Component, OnInit } from '@angular/core';
import { SlotsService } from '../services/slots.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { DoctorEditComponent } from '../doctor-edit/doctor-edit.component';

export interface Slot {
  id: string;
  date: string;
}

@Component({
  selector: 'app-doctor',
  templateUrl: './doctor.component.html',
  styleUrls: ['./doctor.component.scss']
})
export class DoctorComponent implements OnInit {
  newSlotDate: string = '';
  displayedColumns: string[] = ['id', 'date', 'action'];
  dataSource!: MatTableDataSource<Slot>;

  constructor(private slotService: SlotsService, private _dialog: MatDialog) {}

  ngOnInit(): void {
    this.getSlotList()
  }
  openDoctorEdit(id: string){
    this._dialog.open(DoctorEditComponent,  {
      data: { id },
    } );

  }
  addSlot(){
    this.slotService.AddDoctorSlot(this.newSlotDate).subscribe({
      next: (res)=>{
        alert("slot added ")
        this.getSlotList()

      },
      error: console.log

    })
  }
  private getSlotList(): void {
    this.slotService.GetDoctorSlots().subscribe({
      next: (res) => {
        const response = res as { status: number; message: string; Data?: Slot[] };

        if (response.Data) {
          this.dataSource = new MatTableDataSource(response.Data);
        } else {
          console.error('Invalid server response structure');
        }
      },
      error: console.error
    });
  }

  deleteDoctorSlot(id:string){
    this.slotService.DeleteSlots(id).subscribe({
      next: (res)=>{
        alert("slot deleted")
        this.getSlotList()
      },
      error: console.log
    })
  }
}

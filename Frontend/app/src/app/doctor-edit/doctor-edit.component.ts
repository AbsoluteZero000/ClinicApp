import { Component, Inject, OnInit } from '@angular/core';
import { SlotsService } from '../services/slots.service';
import { MatTableDataSource } from '@angular/material/table';
import { MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { FormBuilder, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-doctor-edit',
  templateUrl: './doctor-edit.component.html',
  styleUrls: ['./doctor-edit.component.scss']
})
export class DoctorEditComponent implements OnInit {
  editForm!: FormGroup;

  constructor(
    private slotService: SlotsService,
    @Inject(MAT_DIALOG_DATA) private data: { id: string },
    private _formBuilder: FormBuilder
  ) {}

  ngOnInit() {
    this.editForm = this._formBuilder.group({
      newSlotDate: '',
    });
  }
  onFormSubmit(){
    console.log(this.editForm.value.newSlotDate);
    console.log(this.data)
    this.slotService.EditDoctorSlot(this.data.id, this.editForm.value.newSlotDate).subscribe({
      next: (res)=>{
        alert("slot added ")
      }
    })
  }
}


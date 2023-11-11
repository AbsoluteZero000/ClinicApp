import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { SignupComponent } from './signup/signup.component';
import { LoginComponent } from './login/login.component';
import { SlotsService } from './services/slots.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'app';
  constructor(private _dialog:MatDialog, private _slotService: SlotsService){}

  ngOnInit(): void {
      this.getSlotList()
  }

  openSignupForm(){
    // this._dialog.open(SignupComponent)

    var x = this._slotService.GetDoctorSlots()
    console.log(x)
    x.subscribe({
      next: (res) => {
        console.log(res)
      },
      error: console.log
    })
  }
  openLoginForm(){
    this._dialog.open(LoginComponent)
  }
  getSlotList(){
    var x = this._slotService.GetDoctorSlots()
    console.log(x)
    x.subscribe({
      next: (res) => {
        console.log(res)
      },
      error: console.log
    })
  }

}

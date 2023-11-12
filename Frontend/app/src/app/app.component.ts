import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { SignupComponent } from './signup/signup.component';
import { LoginComponent } from './login/login.component';
import { SlotsService } from './services/slots.service';
import { MatTableDataSource } from '@angular/material/table';
import { AuthService } from './services/auth.service';
import { UsersService } from './services/users.service';
import { DoctorComponent } from './doctor/doctor.component';
import { SharedService } from './shared.service';

export interface slot {
  id: string;
  date: string;
}


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent {
  title = 'app';

  constructor(private _dialog:MatDialog, private authService: AuthService, public sharedService: SharedService){}
  displayedColumns: string[] = ['id', 'date'];

  dataSource!: MatTableDataSource<slot>;


  openSignupForm(){
    this._dialog.open(SignupComponent)


  }
  openLoginForm(){

    this._dialog.open(LoginComponent)

  }
  signout(){
    this.authService.clearAuthentication()
    this.sharedService.isDoctor = false
    this.sharedService.isPatient = false
    this.sharedService.isLoggedIn = false
  }
  getSlots(){

  }
}

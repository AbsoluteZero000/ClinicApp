import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UsersService } from '../services/users.service';
import { DialogRef } from '@angular/cdk/dialog';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {

  loginForm: FormGroup;
  constructor(private _formBuilder: FormBuilder, private _userService: UsersService, private _dialogRef: DialogRef<LoginComponent>){
    this.loginForm = this._formBuilder.group({
      name:'',
      password:''
    });
  }
  hide =true
  hideRequiredControl = new FormControl(false);

  onFormSubmit(){
    if(this.loginForm.valid){
      console.log(this.loginForm.value)
      this._userService.Login(this.loginForm.value)
      // this._dialogRef.close()
    }
  }
}

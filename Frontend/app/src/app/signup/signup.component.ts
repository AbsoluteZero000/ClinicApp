import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { FloatLabelType } from '@angular/material/form-field';
import { UsersService } from '../services/users.service';
import { DialogRef } from '@angular/cdk/dialog';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss'],
})
export class SignupComponent {
  role = "doctor"

  signupForm: FormGroup;
  constructor(private _formBuilder: FormBuilder, private _userService: UsersService, private _dialogRef: DialogRef<SignupComponent>) {
    this.signupForm = this._formBuilder.group({
      name: '',
      email:'',
      password:'',
      role:''
    });

  }




  hide = true
  email = new FormControl('', [Validators.required, Validators.email]);


  hideRequiredControl = new FormControl(false);
  floatLabelControl = new FormControl('auto' as FloatLabelType);
  options = this._formBuilder.group({
    hideRequired: this.hideRequiredControl,
    floatLabel: this.floatLabelControl,
  });


  getFloatLabelValue(): FloatLabelType {
    return this.floatLabelControl.value || 'auto';
  }

  onFormSubmit(){
    if(this.signupForm.valid){
      console.log(this.signupForm.value)
      this._userService.Signup(this.signupForm.value).subscribe({
        next: (val: any) => {
          alert("you've signed up successfully");
          this._dialogRef.close();
        },
        error: (err:any)=> {
          console.error(err);
        }
      })
    }
  }
}

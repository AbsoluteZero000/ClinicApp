import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, takeUntil } from 'rxjs';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';
import { SharedService } from '../shared.service';

@Injectable({
  providedIn: 'root',
})
export class UsersService {

  constructor(
    private _http: HttpClient,
    private authService: AuthService,
    private router: Router,
    private sharedService: SharedService
  ) {}

  Signup(data: any): Observable<any> {
    let body = {
      'name': data.name,
      'email': data.email,
      'password': data.password,
      'role': data.role
    }
    console.log(body)
    return this._http.post('http://localhost:8080/signup', body);
  }

  Login(data: any) {
    const base64Credentials = btoa(`${data.name}:${data.password}`);
    const headers = new HttpHeaders({
      Authorization: `Basic ${base64Credentials}`,
    });
    console.log(base64Credentials);

    this._http
      .post('http://localhost:8080/login', {}, { headers, observe: 'response' })
      .subscribe({
        next: (res) => {
          const response = res.body as {
            status: number;
            message: string;
            Data?: any;
          };
          console.log('Login success:', res);
          if (response.Data.role == 'patient') {
            this.sharedService.isDoctor = false;
            this.sharedService.isPatient = true;
          }
          else if (response.Data.role == 'doctor') {
            this.sharedService.isDoctor = true;
            this.sharedService.isPatient = false;
          }
          else{
            this.sharedService.isDoctor = true;
            this.sharedService.isPatient = true;

          }

          this.authService.setAuthentication(base64Credentials);
          this.sharedService.isLoggedIn = true;
        },
        error: console.log,
      });
  }
}

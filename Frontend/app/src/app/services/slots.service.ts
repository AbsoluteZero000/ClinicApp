import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class SlotsService {

  constructor(private _http: HttpClient, private authService: AuthService) {}


  GetDoctorSlots(){
    const authToken =  this.authService.getAuthentication();
    console.log(this.authService.getAuthentication())
    const headers = new HttpHeaders({
      'Cookie': 'Session=' + authToken,
    });

    return this._http.get('http://localhost:8080/getdoctorslots', { headers });
  }
}


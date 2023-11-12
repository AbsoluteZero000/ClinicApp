import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SlotsService {

  constructor(private _http: HttpClient, private authService: AuthService) {}

  getFreeSlots() {
    const authToken =  this.authService.getAuthentication();
    console.log(this.authService.getAuthentication())
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,

    });

    return this._http.get('http://localhost:8080/getfreeslots', { headers });
  }

  AddPatientSlot(id: string){
    const authToken =  this.authService.getAuthentication()
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,
    })
    return this._http.post('http://localhost:8080/addpatientslot', id, {headers});
  }
  AddDoctorSlot(x: string){
    const authToken =  this.authService.getAuthentication()
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,
    })
    return this._http.post('http://localhost:8080/addslot', {"date": x}, {headers});

  }
  GetDoctorSlots(){
    const authToken =  this.authService.getAuthentication();
    console.log(this.authService.getAuthentication())
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,

    });

    return this._http.get('http://localhost:8080/getdoctorslots', { headers });
  }

  GetPatientSlots(){
    const authToken =  this.authService.getAuthentication();
    console.log(this.authService.getAuthentication())
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,

    });

    return this._http.get('http://localhost:8080/getpatientslots', { headers });
  }

  DeleteSlots(id : string): Observable<any>{

    const authToken = this.authService.getAuthentication();
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,
    });


    // Use the DELETE method
    return this._http.delete(`http://localhost:8080/deleteslot?id=${id}`, {headers});
  }

}




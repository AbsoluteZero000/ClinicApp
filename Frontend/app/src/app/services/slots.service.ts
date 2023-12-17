import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';
import { Observable } from 'rxjs';
import { SharedService } from '../shared.service';

@Injectable({
  providedIn: 'root'
})
export class SlotsService {

  constructor(private _http: HttpClient,
     private authService: AuthService,
    private sharedService: SharedService) {}

  getFreeSlots() {
    const authToken =  this.authService.getAuthentication();
    console.log(this.authService.getAuthentication())
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,

    });

    return this._http.get(this.sharedService.apiURL+'/getfreeslots', { headers });
  }

  EditDoctorSlot(id: string, date:string){
    const authToken =  this.authService.getAuthentication()
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,
    })
    console.log({"id":id,"date":date})

    return this._http.put(this.sharedService.apiURL+'/editdoctorslot', {"id":id,"date":date}, {headers});
  }
  AddPatientSlot(id: string){
    const authToken =  this.authService.getAuthentication()
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,
    })
    return this._http.post(this.sharedService.apiURL+'/addpatientslot', id, {headers});
  }
  AddDoctorSlot(x: string){
    const authToken =  this.authService.getAuthentication()
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,
    })
    return this._http.post(this.sharedService.apiURL+'/addslot', {"date": x}, {headers});

  }
  GetDoctorSlots(){
    const authToken =  this.authService.getAuthentication();
    console.log(this.authService.getAuthentication())
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,

    });

    return this._http.get(this.sharedService.apiURL+'/getdoctorslots', { headers });
  }

  GetPatientSlots(){
    const authToken =  this.authService.getAuthentication();
    console.log(this.authService.getAuthentication())
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,

    });

    return this._http.get(this.sharedService.apiURL+'/getpatientslots', { headers });
  }

  DeleteSlots(id : string): Observable<any>{

    const authToken = this.authService.getAuthentication();
    const headers = new HttpHeaders({
      Authorization: `Basic ${this.authService.getAuthentication()}`,
    });

    return this._http.delete(`http://localhost:8080/deleteslot?id=${id}`, {headers});
  }

}




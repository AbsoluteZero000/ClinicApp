import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { environment } from '../environment/environment';
@Injectable({
  providedIn: 'root'
})
export class SharedService {
  isLoggedIn = false;
  isDoctor = false;
  isPatient = false;
  apiURL = environment.apiKey;

  private dataUpdatedSource = new Subject<void>();

  dataUpdated$ = this.dataUpdatedSource.asObservable();

  triggerDataUpdate() {
    this.dataUpdatedSource.next();
  }
}

import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SharedService {
  isLoggedIn = false;
  isDoctor = false;
  isPatient = false;
  private dataUpdatedSource = new Subject<void>();

  dataUpdated$ = this.dataUpdatedSource.asObservable();

  triggerDataUpdate() {
    this.dataUpdatedSource.next();
  }
}

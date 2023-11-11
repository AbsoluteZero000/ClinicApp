import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class UsersService {
  constructor(private _http: HttpClient, private authService: AuthService) {}
  Signup(data: any): Observable<any> {
    let body = new FormData();
    body.append('name', data.name);
    body.append('email', data.email);
    body.append('password', data.password);
    body.append('role', data.role);
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
      .subscribe((response) => {
        const setCookieHeader = response.headers.get('Set-Cookie');
        console.log(setCookieHeader)
        const sessionCookieValue =
          this.extractSessionCookieValue(setCookieHeader);
          if (sessionCookieValue !== null) {
            this.authService.setAuthentication(sessionCookieValue);
          } else {
            console.error('Session cookie value is null.');
          }
      });
  }
  private extractSessionCookieValue(
    setCookieHeader: string | null
  ): string | null {
    if (!setCookieHeader) {
      return null;
    }

    const matches = setCookieHeader.match(/Session=([^;]+)/);
    return matches ? matches[1] : null;
  }
}

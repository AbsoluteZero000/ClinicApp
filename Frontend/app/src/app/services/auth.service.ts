import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private cookieService: CookieService) {}

  setAuthentication(token: string): void {
    // Save the authentication token in a cookie
    this.cookieService.set('auth_token', token);
  }

  isAuthenticated(): boolean {
    // Check if the authentication token is present
    return this.cookieService.check('auth_token');
  }

  getAuthentication(): string {
    // Retrieve the authentication token from the cookie
    return this.cookieService.get('auth_token');
  }

  clearAuthentication(): void {
    // Clear the authentication token from the cookie
    this.cookieService.delete('auth_token');
  }
}

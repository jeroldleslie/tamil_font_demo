import { Injectable } from '@angular/core';
import { Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map'
import { HttpClient } from './http-client.service';
import { environment } from '../../environments/environment';
import { WebStorageService } from './web-storage.service';

@Injectable()
export class AuthenticationService {
  
  constructor(private http: HttpClient,
            private storage:WebStorageService) {
   }
    apiBaseUrl = environment.apiBaseUrl;

    login(username: string, password: string) {

      console.log(this.apiBaseUrl);
      
      let url = this.apiBaseUrl +'/login'
      let data = JSON.stringify({ username: username, password: password })
      

        return this.http.post(url, data)
            .map((response: Response) => {
                // login successful if there's a jwt token in the response
                //console.log(JSON.stringify(response.json()))
                let resp = response.json();
                if (resp) {
                    // store user details and jwt token in local storage to keep user logged in between page refreshes
                    this.storage.setItem('currentUser', JSON.stringify(resp));
                    this.storage.setItem('currentUsername', username)
                }

                return resp;
            });
    }

    logout() {
        // remove user from local storage to log user out
        this.storage.removeItem('currentUser');
        this.storage.removeItem('currentUsername');
    }
}

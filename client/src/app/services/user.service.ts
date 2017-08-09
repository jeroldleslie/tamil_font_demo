import { Injectable } from '@angular/core';
import { HttpClient } from './http-client.service';
import { environment } from '../../environments/environment';
import { WebStorageService } from './web-storage.service';
import { User } from '../models/index';

@Injectable()
export class UserService {

  constructor(private http: HttpClient,
            private storage:WebStorageService) {
   }
    apiBaseUrl = environment.apiBaseUrl;

    create(user: User) {
        console.log(JSON.stringify(user))
        return this.http.post(this.apiBaseUrl+'/users', user);
    }

}

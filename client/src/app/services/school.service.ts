import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { HttpClient } from './http-client.service';

@Injectable()
export class SchoolService {

  constructor(private http: HttpClient) { }
  apiBaseUrl = environment.apiBaseUrl;

  getSchoolsByIds(ids:string){
    return this.http.get(this.apiBaseUrl+'/schools/getbyids?ids='+ids,null);
  }
}

import { Injectable } from '@angular/core';
import { Response, URLSearchParams } from '@angular/http';
import { HttpClient } from './http-client.service';
import { environment } from '../../environments/environment';
import 'rxjs/add/operator/map';
import { Observable } from 'rxjs';


@Injectable()
export class SearchService {

  constructor(private _http: HttpClient) { }
  apiBaseUrl = environment.apiBaseUrl;

  getResults(gpa: string, percentile: string, act: string, sat: string): Observable<Response> {
    // console.log('in service');
    let parameters = new URLSearchParams();
    parameters.set('gpa', gpa);
    parameters.set('percentile', percentile);
    parameters.set('act', act)
    parameters.set('sat', sat);
    return this._http.get(this.apiBaseUrl+'/searchschools', parameters);
  }
}

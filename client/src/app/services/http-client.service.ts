import { Injectable } from '@angular/core';
import { Http, Headers } from '@angular/http';

@Injectable()
export class HttpClient {

  constructor(private http: Http) { }

  createAuthorizationHeader(headers: Headers) {
    headers.append('Content-Type', 'application/json');
    console.log("createAuthorizationHeader")
    if (localStorage.getItem('currentUser')) {
      console.log(JSON.parse(localStorage.getItem('currentUser')).data.token)
      headers.append('Authorization', 'Bearer ' +
        JSON.parse(localStorage.getItem('currentUser')).data.token); 
    }
  }

  get(url, parameters) {
    let headers = new Headers();
    this.createAuthorizationHeader(headers);
    return this.http.get(url, {
      headers: headers,
      search: parameters,
    });
  }
  
  post(url, data) {
    let headers = new Headers();
    this.createAuthorizationHeader(headers);
    return this.http.post(url, data, {
      headers: headers
    });
  }

  delete(url){
    let headers = new Headers();
    this.createAuthorizationHeader(headers);
    return this.http.delete(url, {
      headers: headers
    });
  }

}

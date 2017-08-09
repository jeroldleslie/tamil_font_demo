import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';

@Injectable()
export class WebStorageService {
  private subject = new Subject<any>();
  constructor() { }

  setItem(key:string, value:string){
    localStorage.setItem(key, value);
    this.subject.next("s");
  }

  getItem(key:string){
    return localStorage.getItem(key);
  }

  removeItem(key:string){
    localStorage.removeItem(key);
    this.subject.next("r");
  }

  listenItemChange(){
    return this.subject.asObservable();
  }

}

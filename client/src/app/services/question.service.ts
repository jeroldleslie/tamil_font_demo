import { Injectable } from '@angular/core';
import { HttpClient } from './http-client.service';
import { environment } from '../../environments/environment';
import { WebStorageService } from './web-storage.service';
import { Question } from '../models/index';


@Injectable()
export class QuestionService {

  constructor(private http: HttpClient,
            private storage:WebStorageService) {
   }
    apiBaseUrl = environment.apiBaseUrl;

    create(question: Question) {
        console.log(JSON.stringify(question));
        return this.http.post(this.apiBaseUrl+'/question', question);
    }


    getAll(){
      return this.http.get(this.apiBaseUrl+'/question',null);
    }

    delete(id:string){
       return this.http.delete(this.apiBaseUrl+'/question/'+id);
    }
}

import { Component, OnInit,Inject } from '@angular/core';
import {MD_DIALOG_DATA, MdDialogRef} from '@angular/material';
import { AlertService, QuestionService } from '../../../services/index';

@Component({
  selector: 'app-question-form',
  templateUrl: './question-form.component.html',
  styleUrls: ['./question-form.component.css']
})
export class QuestionFormComponent implements OnInit {

  content:string;
  private config = {};

  constructor(@Inject(MD_DIALOG_DATA) public data: any,
  public dialogRef: MdDialogRef<QuestionFormComponent>,
private questionService: QuestionService,
      private alertService: AlertService) { 
        this.config = {
    contentsCss : ['assets/fonts/typeface-arima-madurai/index.css' , 
    'assets/fonts/typeface-catamaran/index.css', 'assets/fonts/fonts.css'], 
    font_names :data.font, 
    font_defaultLabel : data.font}
      }

  ngOnInit() {
  }
  onBlur(e){
  }

  onChange(e){
  }

  onFocus(e){
  }

  onReady(e){
  }
  clear(){
    this.content = "";
  }
  save(){
    let data:any = {
      question :'<div style="font-family:'+this.data.font+'">' +this.content+ '</div>'
       }
    this.questionService.create(data)
          .subscribe(
              response => {
                  this.alertService.success('Question created successful', true);
                  this.dialogRef.close();
              },
              error => {
                  this.alertService.error(error.json().message);
              });
  }

}

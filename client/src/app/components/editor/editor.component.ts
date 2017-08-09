import { Component, OnInit } from '@angular/core';
import { AlertService, QuestionService } from '../../services/index';
import { Question } from '../../models/index';
import { MdDialog , MdDialogRef} from '@angular/material';
import { QuestionFormComponent } from './question-form/question-form.component'
@Component({
  selector: 'app-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.css']
})
export class EditorComponent implements OnInit {

  content:string;
  selectedValue: string = "";
  languages = [
    {viewValue: 'English', value: 'Arial'},
    {viewValue: 'Tamil', value: 'Arima Madurai'},
  ];




  constructor(
      public dialog: MdDialog,
      private questionService: QuestionService,
      private alertService: AlertService) { 
      
      }

  /* constructor() {
    this.content = `<p><span style="font-family:Arima Madurai"><span style="font-size:20px">ஆஈஊஈஐன்றளீஊஐளீஊனீஊஐஐஈஊஐஈஊஐஈஊஐஊஈஏஊஈஏன்றளன்றள்ளீறீ</span></span></p>`
   } */

  ngOnInit() {
    this.getAll()
  }

  
  onDelete(e){
    this.getAll()
  }

  onClick(){

    if(this.selectedValue != ""){

      let dialogRef = this.dialog.open(QuestionFormComponent,{
        data: { "font" : this.selectedValue },
        height: '600px',
        width: '800px',
      });
      dialogRef.afterClosed().subscribe(result => {
        //this.selectedOption = result;
        this.getAll();
      });

      

     /*  let data:any = {
      question :'<div style="font-family:Arima Madurai">' +this.content+ '</div>'
       }
    this.questionService.create(data)
          .subscribe(
              response => {
                  this.getAll()
                  this.alertService.success('Question created successful', true);
              },
              error => {
                  this.alertService.error(error.json().message);
              }); */
    }else{
      this.alertService.error("Please select language");
    }
   

  }

  questions = [];
  getAll(){
    this.questions = [];
    this.questionService.getAll()
          .subscribe(
              response => {
                  //console.log(response.json())
                  console.log(JSON.stringify(response));
                  this.questions = response.json().data
              },
              error => {
                  this.alertService.error(error.json().message);
    });
  }


}

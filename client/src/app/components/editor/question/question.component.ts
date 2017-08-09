import { Component, EventEmitter,ElementRef, OnInit,AfterViewInit,ViewChild, Output, Input, ViewEncapsulation } from '@angular/core';
import { AlertService, QuestionService } from '../../../services/index';

@Component({
  selector: 'question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class QuestionComponent implements OnInit {

  @Input()
  question;
 
  @Input()
  index = 0;

  @Output() deleted: EventEmitter<any> = new EventEmitter();

  @ViewChild('target') target;

  private data = "";
  constructor(private questionService: QuestionService,
      private alertService: AlertService,
        elm: ElementRef) { 
          //alert(this.index);
          //this.index = elm.nativeElement.getAttribute('data-index');
          
        }

  ngOnInit() {
      let element = this.target.nativeElement;
      this.data = this.question.question+""
      element.innerHTML = this.data;
  }


  delete(){
    this.questionService.delete(this.question.id)
          .subscribe(
              response => {
                  this.alertService.success('Question deleted successful', true);
                  this.deleted.emit(null);
              },
              error => {
                  this.alertService.error(error.json().message);
              });
  }
}

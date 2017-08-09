import { Component, OnInit } from '@angular/core';
import { MdSnackBar } from '@angular/material';
import { AlertService } from '../../services/index';

@Component({
  selector: 'alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.css']
})
export class AlertComponent implements OnInit {

  constructor(private alertService: AlertService, public snackBar: MdSnackBar) {
    
   }

  ngOnInit() {
    this.alertService.getMessage().subscribe(message => { 
      if(message){
        this.openSnackBar(message.text, "");
      }
      
     });
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {
      duration: 2000,
    });
  }

  

}

import { Component, OnInit } from '@angular/core';
import {MdDialog} from '@angular/material';
import { SchoolDetailsComponent } from '../school-details/school-details.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  constructor(public dialog: MdDialog) { }

  openDialog() {
    this.dialog.open(SchoolDetailsComponent);
  }

  ngOnInit() {
  }

}

import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthenticationService , WebStorageService} from './services/index';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  routeLinks:any[];
  activeLinkIndex = 0;
  signInButtontext = "Signin / Signup";

  constructor(private router: Router, 
            private auth:AuthenticationService,
            private storage:WebStorageService) {

    this.setWelcomeLabel();          
    
    this.storage.listenItemChange().subscribe(data => { 
      this.setWelcomeLabel();      
    });

    this.routeLinks = [
    {label: 'Home', link: 'home'},
    {label: 'Schools', link: 'schools'},
    {label: 'Students', link: 'students'},
    {label: 'Landing', link: 'landing'},
    {label: 'Counselor', link: 'counselor'},
    {label: 'Search', link: 'search'},
    ];
  }

  private setWelcomeLabel(){
    if (this.storage.getItem('currentUsername')) {
      this.signInButtontext = "Hi! " +this.storage.getItem('currentUsername')
    }else{
      this.signInButtontext = "Signin / Signup";
    }
  }

  signin(){
    if(this.signInButtontext == "Signin / Signup"){
      this.router.navigate(["/login"]);
    }
  }
  logout(){
    this.auth.logout();
    this.router.navigate(["/login"]);
  }
}

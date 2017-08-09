import 'hammerjs';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { routedComponents, AppRoutingModule } from './app-routing.module';
import { HttpModule }    from '@angular/http';
import { NgModule } from '@angular/core';
import { FormsModule }   from '@angular/forms';
import { CKEditorModule } from 'ng2-ckeditor';

import { AppComponent } from './app.component';
import { MdToolbarModule,MdIconModule, 
  MdInputModule,MdButtonModule, 
  MdSnackBarModule,MdDialogModule,MdCardModule, 
MdSelectModule } from '@angular/material';

import { NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { FlexLayoutModule } from '@angular/flex-layout';

import { AuthGuard } from './gaurds/index';

import { HttpClient, SearchService, AuthenticationService, AlertService, WebStorageService, UserService, WindowRefService, SchoolService, QuestionService} from './services/index';
import { AlertComponent } from './components/alert/alert.component';
import { EditorComponent } from './components/editor/editor.component';
import { QuestionComponent } from './components/editor/question/question.component';
import { QuestionFormComponent } from './components/editor/question-form/question-form.component';

@NgModule({
  declarations: [
    AppComponent,
    routedComponents,
    AlertComponent,
    EditorComponent,
    QuestionComponent,
    QuestionFormComponent,
  ],
  imports: [
    AppRoutingModule,
    FlexLayoutModule,
    BrowserModule,
    FormsModule,
    CKEditorModule,
    BrowserAnimationsModule,
    HttpModule,
    NgbModule.forRoot(),
    MdIconModule,
    MdToolbarModule,
    MdInputModule,
    MdButtonModule,
    MdSnackBarModule,
    MdDialogModule,
    MdCardModule,
    MdSelectModule
  ],
  entryComponents: [
    QuestionFormComponent
  ],
  providers: [
    AuthGuard,
    HttpClient,
    AuthenticationService,
    SearchService,
    AlertService,
    WebStorageService,
    UserService,
    WindowRefService,
    SchoolService,
    QuestionService
    ],
  bootstrap: [AppComponent]
})


export class AppModule { }

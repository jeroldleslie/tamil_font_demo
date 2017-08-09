import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthGuard } from './gaurds/index';

import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { AlertComponent } from './components/alert/alert.component';


const routes: Routes = [
  { path: '**', component: PageNotFoundComponent }
//  { path: '**', redirectTo: '' }
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, { useHash: false }),
    ],
    exports: [
        RouterModule,
    ]
})
export class AppRoutingModule { }
export const routedComponents: any[] = [ PageNotFoundComponent, AlertComponent
];
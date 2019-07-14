import { Routes, RouterModule } from '@angular/router';
import { LoginComponent, RegistrationComponent } from './modules/adminservice';
import { HomeComponent } from './modules/home';
import {AuthGuard} from './core/guards'

const appRoutes: Routes = [
    { path: '', component: HomeComponent, canActivate: [AuthGuard] },
    { path: 'login', component: LoginComponent },
    { path: 'registration', component: RegistrationComponent },

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];

export const routing = RouterModule.forRoot(appRoutes);
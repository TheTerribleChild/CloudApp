import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppComponent } from './app.component';
import { routing } from './app.routing.module';
import { AuthGuard } from './core/guards'

import { RegistrationService } from './core/services';
import { LoginComponent, RegistrationComponent } from './modules/adminservice';
import { HomeComponent } from './modules/home'

@NgModule({
    imports: [
        BrowserModule,
        ReactiveFormsModule,
        HttpClientModule,
        routing
    ],
    declarations: [
        AppComponent,
        LoginComponent,
        RegistrationComponent,
        HomeComponent
    ],
    providers: [
        RegistrationService,
        AuthGuard
    ],
    bootstrap: [AppComponent]
})

export class AppModule { }
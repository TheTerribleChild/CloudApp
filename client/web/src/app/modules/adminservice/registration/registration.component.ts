import { Component, OnInit } from "@angular/core";
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { first } from 'rxjs/operators';

import { RegistrationService } from "src/app/core/services"
import { RegistrationState } from "./registration.component.state"

@Component({
    selector : 'registration',
    templateUrl : './registration.component.html',
})

export class RegistrationComponent implements OnInit {
    registerEmailForm: FormGroup;
    verificationCodeForm: FormGroup;
    loading = false;
    submittedEmail = false;
    state = RegistrationState.Start;
    verificationToken: string;

    constructor(
        private formBuilder: FormBuilder,
        private router: Router,
        private registrationService: RegistrationService,
        ){}

    ngOnInit() {
        this.registerEmailForm = this.formBuilder.group({
            email: ['', Validators.required]
        });
        this.verificationCodeForm = this.formBuilder.group({
            code: ['', Validators.required]
        });
        this.state = RegistrationState.Start;
    }

    get ef() {
        return this.registerEmailForm.controls;
    }

    get cf() {
        return this.verificationCodeForm.controls;
    }

    onSubmitEmail() {
        this.submittedEmail = true;
        if (this.registerEmailForm.invalid) {
            return;
        }
        this.loading = true;
        this.registrationService.submitEmail(this.ef.email.value)
            .pipe(first())
            .subscribe(
                data => {
                    console.log(data);
                    this.verificationToken = data['verification_token'];
                    console.log(this.verificationToken);
                    this.loading = false;
                    this.state = RegistrationState.VerifiedEmail;
                },
                error => {
                    console.log(error)
                    this.loading = false;
                }
            );
        console.log(this.registerEmailForm.get('email').value);    
    }

    onSubmitCode() {
        if (this.verificationCodeForm.invalid) {
            return;
        }
        this.loading = true;
        this.registrationService.verifyCode(this.verificationToken, this.cf.code.value)
            .pipe(first())
            .subscribe(
                data => {
                    this.loading = false;
                    this.state = RegistrationState.SetPassword;
                },
                error => {
                    this.loading = false;
                }
            )
    }

    isStart() {
        return this.state == RegistrationState.Start;
    }

    isVerifiedEmail() {
        return this.state == RegistrationState.VerifiedEmail;
    }

    isSetPassword() {
        return this.state == RegistrationState.SetPassword;
    }
}
import { Component, OnInit } from "@angular/core";
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { first } from 'rxjs/operators';

import { RegistrationService } from "src/app/core/services"
import { RegistrationState } from "./registration.component.state"
import { HttpClient } from '@angular/common/http';

@Component({
    selector : 'registration',
    templateUrl : './registration.component.html',
})

export class RegistrationComponent implements OnInit {
    registerEmailForm: FormGroup;
    verificationCodeForm: FormGroup;
    setPasswordForm: FormGroup;
    loading = false;
    submittedEmail = false;
    passwordMismatch = false;
    validValidationCode = true;
    state = RegistrationState.Start;
    verificationToken: string;
    setPasswordTokenId: string;

    constructor(
        private formBuilder: FormBuilder,
        private router: Router,
        private registrationService: RegistrationService,
        private http: HttpClient
        ){
            
        }

    ngOnInit() {
        this.registerEmailForm = this.formBuilder.group({
            email: ['', Validators.required]
        });
        this.verificationCodeForm = this.formBuilder.group({
            code: ['', Validators.required]
        });
        this.setPasswordForm = this.formBuilder.group({
            password: ['', [Validators.required, Validators.minLength(6)]],
            passwordConfirmation: ['', [Validators.required, Validators.minLength(6)]]
        });
        this.state = RegistrationState.Start;
    }

    get ef() {
        return this.registerEmailForm.controls;
    }

    get cf() {
        return this.verificationCodeForm.controls;
    }

    get pf() {
        return this.setPasswordForm.controls;
    }

    onSubmitEmail() {
        this.submittedEmail = true;
        if (this.registerEmailForm.invalid) {
            return;
        }
        this.loading = true;
        this.registrationService.submitEmail2();
        // this.registrationService.submitEmail(this.ef.email.value)
        //     .pipe(first())
        //     .subscribe(
        //         data => {
        //             console.log(data);
        //             this.verificationToken = data['verification_token'];
        //             console.log(this.verificationToken);
        //             this.loading = false;
        //             this.state = RegistrationState.VerifiedEmail;
        //         },
        //         error => {
        //             console.log(error)
        //             this.loading = false;
        //         }
        //     );
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
                    this.setPasswordTokenId = data['set_password_token_id'];
                    this.loading = false;
                    if(this.setPasswordTokenId){
                        this.state = RegistrationState.SetPassword;
                    } else {

                    }
                    console.log(data);
                },
                error => {
                    this.loading = false;
                    this.validValidationCode = false;
                    console.log(error);
                }
            )
    }

    onSubmitPassword() {
        if (this.setPasswordForm.invalid) {
            return;
        }
        if (this.pf.password.value != this.pf.passwordConfirmation.value){
            this.passwordMismatch = true;
            return;
        }
        this.loading = true;
        this.registrationService.setPassword(this.setPasswordTokenId, this.pf.password.value)
            .pipe(first())
            .subscribe(
                data => {
                    this.router.navigate(['/login']);
                    console.log(data);
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

    isValidVerificationCode() {
        return this.validValidationCode;
    }
}
import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { RegisterEmailMessage, VerifyCodeTokenMessage } from 'src/app/shared/models/adminservice';
import { environment } from 'src/environments/environment';
import { first } from 'rxjs/operators';

@Injectable()
export class RegistrationService {
    constructor(private http: HttpClient) { }

    submitEmail(email: string) {
        var message = new RegisterEmailMessage();
        message.email = email;
        return this.http.post(`${environment.apiUrl}/admin/v1/register/user`, message);
    }

    verifyCode(verificationToken: string, verificationCode: string) {
        var message = new VerifyCodeTokenMessage();
        message.verification_code = verificationCode;
        message.verification_token = verificationToken;
        return this.http.post(`${environment.apiUrl}/admin/v1/register/verify`, message);
    }
}
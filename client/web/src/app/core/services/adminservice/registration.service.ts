import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { RegisterEmailMessage, VerifyCodeTokenMessage, SetPasswordWithTokenMessage } from 'src/app/shared/models/adminservice';
import { environment } from 'src/environments/environment';

@Injectable()
export class RegistrationService {
    constructor(private http: HttpClient) { }

    // submitEmail(email: string) {
    //     var message = new RegisterEmailMessage();
    //     message.email = email;
    //     return this.http.post(`${environment.apiUrl}/admin/v1/register/user`, message);
    // }

    submitEmail2() {
        return this.http.post(`${environment.apiUrl}/admin/v1/register/user`, {emailONE:"yolo"}).subscribe(
            data => {
                console.log("HERE " + data);
            }
        );
    }

    verifyCode(verificationToken: string, verificationCode: string) {
        var message = new VerifyCodeTokenMessage();
        message.verification_code = verificationCode;
        message.verification_token = verificationToken;
        return this.http.post(`${environment.apiUrl}/admin/v1/register/verify`, message);
    }

    setPassword(setPasswordTokenId: string, newPassword: string) {
        var message = new SetPasswordWithTokenMessage();
        message.new_password = newPassword;
        message.token_id = setPasswordTokenId;
        return this.http.post(`${environment.apiUrl}/admin/v1/users/password`, message);
    }

    getUser() {
        return this.http.get(`http://localhost:8080/api/admin/v1/users/abc`);
    }
}
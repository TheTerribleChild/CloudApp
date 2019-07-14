export class RegisterEmailMessage {
    email: string;
}

export class VerifyCodeTokenMessage {
    verification_token: string;
    verification_code: string;
}

export class VerificationTokenResponse {
    verification_token: string;
}
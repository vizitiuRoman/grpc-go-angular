import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { ToastController } from '@ionic/angular';

import { AuthService } from '@services/auth.service';

@Component({
    selector: 'app-auth',
    templateUrl: './auth.page.html',
    styleUrls: ['./auth.page.scss'],
})
export class AuthPage {
    public form: FormGroup;
    public signupView!: boolean;

    constructor(
        private formBuilder: FormBuilder,
        private authService: AuthService,
        private toastCtrl: ToastController
    ) {
        this.form = this.formBuilder.group({
            email: [
                'roma@mail.ru',
                [Validators.minLength(4), Validators.required],
            ],
            password: [
                'qweqweqwe',
                [Validators.minLength(8), Validators.required],
            ],
        });
    }

    public onSubmit(type: 'register' | 'login'): void {
        if (this.form.valid) {
            const payload = {
                email: this.form.controls.email.value,
                password: this.form.controls.password.value,
            };
            console.log(type);
            const submit =
                type === 'login'
                    ? this.authService.auth(payload)
                    : this.authService.register(payload);

            submit.subscribe(
                () => {},
                (err) => {
                    console.log(err);
                    this.toastCtrl
                        .create({
                            message: err.message,
                            position: 'bottom',
                            duration: 3000,
                        })
                        .then((toast) => toast.present());
                }
            );
        }
    }

    toggleSignUpView(): void {
        this.signupView = !this.signupView;
    }
}

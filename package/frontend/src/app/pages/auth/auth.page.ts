import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { ToastController } from '@ionic/angular';

import { AuthService } from '@services/auth.service';

@Component({
    selector: 'app-auth',
    templateUrl: './auth.page.html',
    styleUrls: ['./auth.page.scss'],
})
export class AuthPage {
    public form: FormGroup;

    constructor(
        private formBuilder: FormBuilder,
        private authService: AuthService,
        private toastCtrl: ToastController,
        private router: Router,
    ) {
        this.form = this.formBuilder.group({
            email: ['roman@mail.ru', [Validators.minLength(4), Validators.required]],
            password: ['romka228', [Validators.minLength(8), Validators.required]],
        });
    }

    public onSubmit(): void {
        if (this.form.valid) {
            this.authService.auth({
                email: this.form.controls.email.value,
                password: this.form.controls.password.value,
            })
                .subscribe(
                    () => {
                        this.form.reset();
                        this.router.navigateByUrl('/home');
                    },
                    (err) => {
                        this.form.reset();
                        this.toastCtrl.create({
                            message: err.message,
                            position: 'bottom',
                            duration: 3000,
                        });
                    }
                );
        }
    }
}

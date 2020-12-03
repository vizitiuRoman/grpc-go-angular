import { Component } from '@angular/core';

import { Platform, ToastController } from '@ionic/angular';
import { SplashScreen } from '@ionic-native/splash-screen/ngx';
import { StatusBar } from '@ionic-native/status-bar/ngx';

import { Observable } from 'rxjs';

import { jwtAuthError$ } from '@grpc/helpers/grpc-jwt';
import { AuthService } from '@services/auth.service';

@Component({
    selector: 'app-root',
    templateUrl: 'app.component.html',
    styleUrls: ['app.component.scss'],
})
export class AppComponent {
    public isAuth$: Observable<boolean> = this.authService.isLoggedIn();

    constructor(
        private platform: Platform,
        private splashScreen: SplashScreen,
        private statusBar: StatusBar,
        private authService: AuthService,
        private toastCtrl: ToastController
    ) {
        this.initializeApp();
    }

    initializeApp(): void {
        this.platform.ready().then(() => {
            this.statusBar.styleDefault();
            this.splashScreen.hide();

            const updateAuth = this.authService.updateAuth();
            if (updateAuth instanceof Observable) {
                updateAuth.subscribe(
                    () => {},
                    (err) => {
                        this.toastCtrl
                            .create({
                                message: 'Error',
                                position: 'bottom',
                                duration: 3000,
                            })
                            .then((toast) => toast.present());
                        this.authService.logout();
                    }
                );
            }

            jwtAuthError$.asObservable().subscribe(() => {
                this.logout();
            });
        });
    }

    public logout(): void {
        this.authService.logout();
    }
}

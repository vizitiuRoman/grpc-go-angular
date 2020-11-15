import { Component, OnDestroy, OnInit } from '@angular/core';

import { ToastController } from '@ionic/angular';

import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';

import { UserGrpcService } from '@grpc/services/user/user.service';
import { User } from '@grpc/grpc-proto/user/user_pb';
import { getUserIdFromJWT } from '@grpc/helpers/grpc-jwt-id';
import { StorageService } from '@services/storage.service';
import { ACCESS_TOKEN } from '@utils/constants';

@Component({
    selector: 'app-home',
    templateUrl: './home.page.html',
    styleUrls: ['./home.page.scss'],
})
export class HomePage implements OnInit, OnDestroy {

    public user: User.AsObject;
    private ngDestroy$ = new Subject<boolean>();

    constructor(
        private storageService: StorageService,
        private userGrpcService: UserGrpcService,
        private toastCtrl: ToastController,
    ) {
    }

    ngOnDestroy(): void {
        this.ngDestroy$.next(true);
        this.ngDestroy$.complete();
    }

    ngOnInit(): void {
        const id = getUserIdFromJWT(
            this.storageService.get<string>(ACCESS_TOKEN)
        );
        this.userGrpcService.getUser({ id })
            .pipe(takeUntil(this.ngDestroy$))
            .subscribe(
                (user) => {
                    this.user = user;
                },
                (err) => {
                    this.toastCtrl.create({
                        message: err.message,
                        position: 'bottom',
                        duration: 3000,
                    })
                        .then((toast) => toast.present());
                }
            );
    }


}

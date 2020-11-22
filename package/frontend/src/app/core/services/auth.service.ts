import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

import { interval, Observable, ReplaySubject, Subject } from 'rxjs';
import { map, switchMap, takeUntil } from 'rxjs/operators';

import { AuthReq, AuthRes, RegisterReq, UpdateAuthRes } from '@grpc/grpc-proto/auth/auth_pb';
import { AuthGrpcService } from '@grpc/services/auth/auth.service';
import { StorageService } from '@services/storage.service';
import { ACCESS_TOKEN, REFRESH_TOKEN } from '@utils/constants';

@Injectable({
    providedIn: 'root'
})
export class AuthService {

    private ngOnDestroy$ = new Subject<void>();
    private loggedInSubject$ = new ReplaySubject<boolean>(1);

    constructor(
        private authGrpcService: AuthGrpcService,
        private router: Router,
        private storageService: StorageService,
    ) {
    }

    private updateToken(): void {
        const aToken = this.storageService.get<string>(ACCESS_TOKEN).split('.')[1];
        const rToken = this.storageService.get<string>(REFRESH_TOKEN);

        const jwt = JSON.parse(atob(aToken));
        const now = Date.now() / 1000;
        const period = Math.ceil(jwt.exp - now - 60) * 1000;

        interval(period)
            .pipe(
                switchMap(() => this.authGrpcService.updateAuth({ refreshtoken: rToken })),
                takeUntil(this.ngOnDestroy$),
            )
            .subscribe(
                (res) => {
                    this.storageService.set(ACCESS_TOKEN, res.token);
                    this.storageService.set(REFRESH_TOKEN, res.refreshtoken);
                },
            );
    }

    public isLoggedIn(): Observable<boolean> {
        return this.loggedInSubject$.asObservable();
    }

    public auth(data: AuthReq.AsObject): Observable<AuthRes.AsObject> {
        return this.authGrpcService.auth(data)
            .pipe(
                map(
                    (res) => {
                        this.storageService.set(ACCESS_TOKEN, res.token);
                        this.storageService.set(REFRESH_TOKEN, res.refreshtoken);
                        this.loggedInSubject$.next(true);
                        this.updateToken();
                        this.router.navigateByUrl('/home');
                        return res;
                    }
                ),
            );
    }

    public register(data: RegisterReq.AsObject): Observable<AuthRes.AsObject> {
        return this.authGrpcService.register(data)
            .pipe(
                map(
                    (res) => {
                        this.storageService.set(ACCESS_TOKEN, res.token);
                        this.storageService.set(REFRESH_TOKEN, res.refreshtoken);
                        this.loggedInSubject$.next(true);
                        this.updateToken();
                        this.router.navigateByUrl('/home');
                        return res;
                    }
                ),
            );
    }

    public updateAuth(): Observable<UpdateAuthRes.AsObject> | void {
        const aToken = this.storageService.get<string>(ACCESS_TOKEN);
        const rToken = this.storageService.get<string>(REFRESH_TOKEN);
        if (aToken && rToken) {
            return this.authGrpcService.updateAuth({ refreshtoken: rToken })
                .pipe(
                    map(
                        (res) => {
                            this.storageService.set(ACCESS_TOKEN, res.token);
                            this.storageService.set(REFRESH_TOKEN, res.refreshtoken);
                            this.loggedInSubject$.next(true);
                            this.updateToken();
                            return res;
                        },
                    )
                );
        }
        return this.logout();
    }

    public logout(): void {
        if (this.storageService.get<string>(ACCESS_TOKEN)) {
            this.authGrpcService.logout().subscribe();
        }
        this.storageService.remove(ACCESS_TOKEN);
        this.storageService.remove(REFRESH_TOKEN);
        this.loggedInSubject$.next(false);
        this.ngOnDestroy$.next();
        this.router.navigateByUrl('/auth');
    }
}

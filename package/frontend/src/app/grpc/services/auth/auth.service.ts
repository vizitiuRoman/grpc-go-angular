import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

import { Metadata } from 'grpc-web';

import { grpcUnary } from '@grpc/helpers/grpc-unary';
import { grpcJwtMetadata } from '@grpc/helpers/grpc-metadata';
import { AuthServicePromiseClient } from '@grpc/grpc-proto/auth/auth_grpc_web_pb';
import { AuthReq, AuthRes, RegisterReq, Stub, UpdateAuthReq, UpdateAuthRes } from '@grpc/grpc-proto/auth/auth_pb';

@Injectable({
    providedIn: 'root',
})
export class AuthGrpcService {

    constructor(private client: AuthServicePromiseClient) {
    }

    public auth(data: AuthReq.AsObject): Observable<AuthRes.AsObject> {
        const req = new AuthReq();

        req.setEmail(data.email);
        req.setPassword(data.password);

        return grpcUnary<AuthRes.AsObject>(this.client.auth(req));
    }

    public register(data: AuthReq.AsObject): Observable<AuthRes.AsObject> {
        const req = new RegisterReq();

        req.setEmail(data.email);
        req.setPassword(data.password);

        return grpcUnary<AuthRes.AsObject>(this.client.register(req));
    }

    public updateAuth(data: UpdateAuthReq.AsObject): Observable<UpdateAuthRes.AsObject> {
        const req = new UpdateAuthReq();
        const meta: Metadata = grpcJwtMetadata();

        req.setRefreshtoken(data.refreshtoken);

        return grpcUnary<AuthRes.AsObject>(this.client.updateAuth(req, meta));
    }

    public logout(): Observable<void> {
        const req = new Stub();
        const meta: Metadata = grpcJwtMetadata();

        return grpcUnary<void>(this.client.logout(req, meta));
    }
}

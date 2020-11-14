import { Injectable } from '@angular/core';
import { UserServicePromiseClient } from '@grpc/grpc-proto/user/user_grpc_web_pb';
import { CreateUserReq, UpdateUserReq, User, UserReq, UserRes } from '@grpc/grpc-proto/user/user_pb';

import { Observable } from 'rxjs';

import { Metadata } from 'grpc-web';

import { grpcUnary } from '@grpc/helpers/grpc-unary';
import { grpcJwtMetadata } from '@grpc/helpers/grpc-metadata';

@Injectable({
    providedIn: 'root',
})
export class UserGrpcService {

    constructor(private client: UserServicePromiseClient) {
    }

    public createUser(data: CreateUserReq.AsObject): Observable<UserRes.AsObject> {
        const req = new CreateUserReq();

        req.setName(data.name);
        req.setEmail(data.email);
        req.setPassword(data.password);

        return grpcUnary<UserRes.AsObject>(this.client.createUser(req));
    }

    public updateUser(data: UpdateUserReq.AsObject): Observable<UserRes.AsObject> {
        const req = new UpdateUserReq();
        const meta: Metadata = grpcJwtMetadata();

        req.setName(data.name);
        req.setEmail(data.email);

        return grpcUnary<UserRes.AsObject>(this.client.updateUser(req, meta));
    }

    public deleteUser(data: UserReq.AsObject): Observable<UserRes.AsObject> {
        const req = new UserReq();

        req.setId(data.id);

        return grpcUnary<UserRes.AsObject>(this.client.deleteUser(req));
    }

    public getUser(data: UserReq.AsObject): Observable<User.AsObject> {
        const req = new UserReq();

        req.setId(data.id);

        return grpcUnary<User.AsObject>(this.client.getUser(req));
    }

}

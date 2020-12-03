import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { AuthServicePromiseClient } from '@grpc/grpc-proto/auth/auth_grpc_web_pb';
import { UserServicePromiseClient } from '@grpc/grpc-proto/user/user_grpc_web_pb';

const GRPC_CLIENTS = [AuthServicePromiseClient, UserServicePromiseClient];

@NgModule({
    imports: [CommonModule],
    providers: GRPC_CLIENTS.map((service) => {
        return {
            provide: service,
            useFactory: () => new service('http://localhost:443', null, null),
        };
    }),
})
export class GrpcModule {}

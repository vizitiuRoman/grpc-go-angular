import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

import { MovieServicePromiseClient } from '@grpc/grpc-proto/movie/movie_grpc_web_pb';
import { MovieRes } from '@grpc/grpc-proto/movie/movie_pb';
import { AuthReq } from '@grpc/grpc-proto/auth/auth_pb';
import { grpcUnary } from '@grpc/helpers/grpc-unary';

@Injectable({
    providedIn: 'root',
})
export class MovieGrpcService {
    constructor(private client: MovieServicePromiseClient) {}

    public getMovies(): Observable<MovieRes.AsObject> {
        const req = new AuthReq();

        return grpcUnary<MovieRes.AsObject>(this.client.getMovies(req));
    }
}

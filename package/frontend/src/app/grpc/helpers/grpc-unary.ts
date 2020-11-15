import { Status, StatusCode } from 'grpc-web';

import { Observable, from, throwError } from 'rxjs';
import { map, catchError } from 'rxjs/operators';

import { jwtAuthError$ } from '@grpc/helpers/grpc-jwt';

import * as protobuf from 'google-protobuf';

export function grpcUnary<T>(promise): Observable<T> {
    return from(promise).pipe(
        map((response: protobuf.Message) => response.toObject()),
        catchError((error: Status) => {
            if (error.code === StatusCode.UNAUTHENTICATED) {
                jwtAuthError$.next();
            }

            return throwError(error);
        }),
    );
}

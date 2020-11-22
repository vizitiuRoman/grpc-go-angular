import { Metadata } from 'grpc-web';

import { ACCESS_TOKEN } from '@utils/constants';

export function grpcJwtMetadata(token: string = ''): Metadata {
    return {
        Authorization: token || JSON.parse(localStorage.getItem(ACCESS_TOKEN) as string),
    };
}

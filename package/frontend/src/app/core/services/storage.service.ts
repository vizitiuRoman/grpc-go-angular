import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root',
})
export class StorageService {
    public set<T>(key: string, value: T): void {
        localStorage.setItem(key, JSON.stringify(value));
    }

    public get<T>(key: string): T {
        return JSON.parse(localStorage.getItem(key) as string) as T;
    }

    public clear(): void {
        localStorage.clear();
    }

    public remove(key: string): void {
        localStorage.removeItem(key);
    }
}

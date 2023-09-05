/* eslint-disable no-console */
import { Injectable } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class PersistanceService {
  public set(key: string, data: unknown): void {
    try {
      localStorage.setItem(key, JSON.stringify(data));
    } catch (err: unknown) {
      console.error('Error saving to local storage');
    }
  }

  public get(key: string): unknown {
    try {
      const localStorageItem: string | null = localStorage.getItem(key);
      return localStorageItem ? JSON.parse(localStorageItem) : null;
    } catch (err: unknown) {
      console.error('Error getting from local storage', err);
      return null;
    }
  }
}

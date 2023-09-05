import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable, map } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({ providedIn: 'root' })
export class PopularTagsService {
  private readonly http: HttpClient = inject(HttpClient);
  private readonly baseUrl: string = environment.baseApiUrl;

  public getPopularTags$(): Observable<string[]> {
    return this.http.get<{ tags: string[] }>(`${this.baseUrl}/tags`).pipe(map(({ tags }): string[] => tags));
  }
}

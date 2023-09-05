import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { GetFeedResponse } from '@core/models/get-feed-response.model';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({ providedIn: 'root' })
export class FeedService {
  private readonly http: HttpClient = inject(HttpClient);
  private readonly baseApiUrl: string = environment.baseApiUrl;

  public getFeed$(url: string): Observable<GetFeedResponse> {
    return this.http.get<GetFeedResponse>(`${this.baseApiUrl}/${url}`);
  }
}

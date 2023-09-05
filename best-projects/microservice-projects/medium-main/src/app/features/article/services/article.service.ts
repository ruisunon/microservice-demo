import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { ArticlePayload } from '@article/models/article-payload.model';
import { ArticleRequest } from '@article/models/article-request.model';
import { Article } from '@core/models/article.model';
import { Observable, map } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({ providedIn: 'root' })
export class ArticleService {
  private readonly http: HttpClient = inject(HttpClient);
  private readonly baseUrl: string = environment.baseApiUrl;

  public getArticle$(slug: string): Observable<Article> {
    return this.http.get<ArticleRequest>(`${this.baseUrl}/articles/${slug}`).pipe(map(({ article }): Article => article));
  }

  public deleteArticle$(slug: string): Observable<{}> {
    return this.http.delete(`${this.baseUrl}/articles/${slug}`);
  }

  public createArticle$(articlePayload: ArticlePayload): Observable<Article> {
    return this.http
      .post<ArticleRequest>(`${this.baseUrl}/articles`, { article: articlePayload })
      .pipe(map(({ article }): Article => article));
  }

  public updateArticle$(slug: string, articlePayload: ArticlePayload): Observable<Article> {
    return this.http
      .put<ArticleRequest>(`${this.baseUrl}/articles/${slug}`, { article: articlePayload })
      .pipe(map(({ article }): Article => article));
  }
}

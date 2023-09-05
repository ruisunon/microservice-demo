import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { ArticleRequest } from '@article/models/article-request.model';
import { Article } from '@core/models/article.model';
import { getArticle } from '@core/utils/get-article';
import { getFavoritesUrl } from '@core/utils/get-favorites-url';
import { Observable, map } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class AddToFavoritesService {
  private readonly http: HttpClient = inject(HttpClient);

  public addToFavorites$(slug: string): Observable<Article> {
    return this.http.post<ArticleRequest>(getFavoritesUrl(slug), {}).pipe(map(getArticle));
  }

  public removeFromFavorites$(slug: string): Observable<Article> {
    return this.http.delete<ArticleRequest>(getFavoritesUrl(slug)).pipe(map(getArticle));
  }
}

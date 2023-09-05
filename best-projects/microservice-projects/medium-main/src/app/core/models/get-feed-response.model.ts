import { Article } from '@core/models/article.model';

export interface GetFeedResponse {
  articles: Article[];
  articlesCount: number;
}

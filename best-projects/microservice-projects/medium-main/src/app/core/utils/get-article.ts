import { ArticleRequest } from '@article/models/article-request.model';
import { Article } from '@core/models/article.model';

export const getArticle = (response: ArticleRequest): Article => {
  return response.article;
};

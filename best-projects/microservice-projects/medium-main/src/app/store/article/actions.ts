import { ArticlePayload } from '@article/models/article-payload.model';
import { Article } from '@core/models/article.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { createAction, props } from '@ngrx/store';

export const getArticle = createAction('[Article] Get article', props<{ slug: string }>());
export const getArticleSuccess = createAction('[Article] Get article success', props<{ article: Article }>());
export const getArticleFailure = createAction('[Article] Get article failure');

export const deleteArticle = createAction('[Article] Delete article', props<{ slug: string }>());
export const deleteArticleSuccess = createAction('[Article] Delete article success');
export const deleteArticleFailure = createAction('[Article] Delete article failure');

export const createArticle = createAction('[Article] Create article', props<{ articlePayload: ArticlePayload }>());
export const createArticleSuccess = createAction('[Article] Create article success', props<{ article: Article }>());
export const createArticleFailure = createAction('[Article] Create article failure', props<{ errors: BackendErrors }>());

export const updateArticle = createAction('[Article] Update article', props<{ articlePayload: ArticlePayload; slug: string }>());
export const updateArticleSuccess = createAction('[Article] Update article success', props<{ article: Article }>());
export const updateArticleFailure = createAction('[Article] Update article failure', props<{ errors: BackendErrors }>());

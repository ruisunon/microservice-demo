import { Article } from '@core/models/article.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { routerNavigatedAction } from '@ngrx/router-store';
import { createReducer, on } from '@ngrx/store';
import { ArticleActions } from '@store/article';

export const FeatureKey = 'article';

export interface State {
  isLoading: boolean;
  error: string | null;
  article: Article | null;
  createArticleErrors: BackendErrors | null;
  isCreateArticleSubmitting: boolean;
}

const initialState: State = {
  isLoading: false,
  error: null,
  article: null,
  createArticleErrors: null,
  isCreateArticleSubmitting: false,
};

export const reducer = createReducer(
  initialState,

  on(ArticleActions.getArticle, (state): State => {
    return { ...state, isLoading: true };
  }),
  on(ArticleActions.getArticleSuccess, (state, { article }): State => {
    return { ...state, isLoading: false, article };
  }),
  on(ArticleActions.getArticleFailure, (state): State => {
    return { ...state, isLoading: false };
  }),

  on(ArticleActions.createArticle, (state): State => {
    return { ...state, isCreateArticleSubmitting: true };
  }),
  on(ArticleActions.createArticleSuccess, (state): State => {
    return { ...state, isCreateArticleSubmitting: false };
  }),
  on(ArticleActions.createArticleFailure, (state, { errors }): State => {
    return { ...state, isCreateArticleSubmitting: false, createArticleErrors: errors };
  }),

  // TODO: refactor article state and update action
  on(ArticleActions.updateArticle, (state): State => {
    return { ...state, isCreateArticleSubmitting: true };
  }),
  on(ArticleActions.updateArticleSuccess, (state, { article }): State => {
    return { ...state, isCreateArticleSubmitting: false, article };
  }),
  on(ArticleActions.updateArticleFailure, (state, { errors }): State => {
    return { ...state, isCreateArticleSubmitting: false, createArticleErrors: errors };
  }),

  on(routerNavigatedAction, () => initialState)
);

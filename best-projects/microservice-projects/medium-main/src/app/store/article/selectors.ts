import { Article } from '@core/models/article.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { createFeatureSelector, createSelector } from '@ngrx/store';
import { FeatureKey, State as ArticleState } from '@store/article/reducer';

const selectArticleState = createFeatureSelector<ArticleState>(FeatureKey);

export const isLoading = createSelector(selectArticleState, ({ isLoading }: ArticleState): boolean => isLoading);
export const error = createSelector(selectArticleState, ({ error }: ArticleState): string | null => error);
export const article = createSelector(selectArticleState, ({ article }: ArticleState): Article | null => article);

export const createArticleErrors = createSelector(
  selectArticleState,
  ({ createArticleErrors }: ArticleState): BackendErrors | null => createArticleErrors
);
export const isCreateArticleSubmitting = createSelector(
  selectArticleState,
  ({ isCreateArticleSubmitting }: ArticleState): boolean => isCreateArticleSubmitting
);

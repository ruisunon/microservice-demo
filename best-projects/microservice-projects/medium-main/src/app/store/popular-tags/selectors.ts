import { createFeatureSelector, createSelector } from '@ngrx/store';
import { State as PopularTagsState } from '@store/popular-tags';
import { FeatureKey } from '@store/popular-tags/reducer';

const selectFeedState = createFeatureSelector<PopularTagsState>(FeatureKey);

export const error = createSelector(selectFeedState, ({ error }: PopularTagsState): string | null => error);
export const isLoading = createSelector(selectFeedState, ({ isLoading }: PopularTagsState): boolean => isLoading);
export const popularTags = createSelector(selectFeedState, ({ popularTags }: PopularTagsState): string[] | null => popularTags);

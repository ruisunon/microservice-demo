import { GetFeedResponse } from '@core/models/get-feed-response.model';
import { createFeatureSelector, createSelector } from '@ngrx/store';
import { FeatureKey, State as FeedState } from '@store/feed';

const selectFeedState = createFeatureSelector<FeedState>(FeatureKey);

export const isLoading = createSelector(selectFeedState, ({ isLoading }: FeedState): boolean => isLoading);
export const errors = createSelector(selectFeedState, ({ errors }: FeedState): string | null => errors);
export const feedData = createSelector(selectFeedState, ({ feed }: FeedState): GetFeedResponse | null => feed);

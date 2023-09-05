import { GetFeedResponse } from '@core/models/get-feed-response.model';
import { routerNavigatedAction } from '@ngrx/router-store';
import { createReducer, on } from '@ngrx/store';
import { FeedActions } from '@store/feed';

export const FeatureKey = 'feed';

export interface State {
  feed: GetFeedResponse | null;
  isLoading: boolean;
  errors: string | null;
}

const initialState: State = {
  feed: null,
  isLoading: false,
  errors: null,
};

export const reducer = createReducer(
  initialState,

  // get feed
  on(FeedActions.getFeed, (state): State => {
    return { ...state, isLoading: true };
  }),
  on(FeedActions.getFeedSuccess, (state, { feed }): State => {
    return { ...state, isLoading: false, feed };
  }),
  on(FeedActions.getFeedFailure, (state): State => {
    return { ...state, isLoading: false };
  }),

  on(routerNavigatedAction, () => initialState)
);

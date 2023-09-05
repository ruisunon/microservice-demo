import { createReducer, on } from '@ngrx/store';
import { PopularTagsActions } from '@store/popular-tags';

export const FeatureKey = 'popular-tags';

export interface State {
  isLoading: boolean;
  error: string | null;
  popularTags: string[] | null;
}

const initialState: State = {
  isLoading: false,
  error: null,
  popularTags: null,
};

export const reducer = createReducer(
  initialState,

  // get popular tags
  on(PopularTagsActions.getPopularTags, (state): State => {
    return { ...state, isLoading: true };
  }),
  on(PopularTagsActions.getPopularTagsSuccess, (state, { popularTags }): State => {
    return { ...state, isLoading: false, popularTags };
  }),
  on(PopularTagsActions.getPopularTagsFailure, (state): State => {
    return { ...state, isLoading: false };
  })
);

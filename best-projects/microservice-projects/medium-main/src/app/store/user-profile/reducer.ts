import { routerNavigatedAction } from '@ngrx/router-store';
import { createReducer, on } from '@ngrx/store';
import { UserProfileActions } from '@store/user-profile';
import { UserProfile } from '@user-profile/models/user-profile.model';

export const FeatureKey = 'user-profile';

export interface State {
  userProfile: UserProfile | null;
  isLoading: boolean;
  error: string | null;
}

const initialState: State = {
  userProfile: null,
  isLoading: false,
  error: null,
};

export const reducer = createReducer(
  initialState,

  // get user profile
  on(UserProfileActions.getUserProfile, (state): State => {
    return { ...state, isLoading: true };
  }),
  on(UserProfileActions.getUserProfileSuccess, (state, { userProfile }): State => {
    return { ...state, isLoading: false, userProfile };
  }),
  on(UserProfileActions.getUserProfileFailure, (state): State => {
    return { ...state, isLoading: false };
  }),

  on(routerNavigatedAction, (): State => initialState)
);

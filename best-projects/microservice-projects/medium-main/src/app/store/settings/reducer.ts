import { BackendErrors } from '@core/models/backend-errors.model';
import { routerNavigatedAction } from '@ngrx/router-store';
import { createReducer, on } from '@ngrx/store';
import { AuthActions } from '@store/auth';

export const FeatureKey = 'settings';

export interface State {
  isSubmitting: boolean;
  errors: BackendErrors | null;
}

const initialState: State = {
  isSubmitting: false,
  errors: null,
};

export const reducer = createReducer(
  initialState,

  // update current user
  on(AuthActions.updateCurrentUser, (state) => {
    return { ...state, isSubmitting: true };
  }),
  on(AuthActions.updateCurrentUserSuccess, (state) => {
    return { ...state, isSubmitting: false };
  }),
  on(AuthActions.updateCurrentUserFailure, (_, { errors }) => {
    return { isSubmitting: false, errors };
  }),

  on(routerNavigatedAction, () => initialState)
);

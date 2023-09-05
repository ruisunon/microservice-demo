import { CurrentUser } from '@auth/models/current-user.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { routerNavigatedAction } from '@ngrx/router-store';
import { createReducer, on } from '@ngrx/store';
import { AuthActions } from '@store/auth';

export const FeatureKey = 'auth';

export interface State {
  isLoading: boolean;
  isSubmitting: boolean;
  currentUser: CurrentUser | null | undefined;
  errors: BackendErrors | null;
}

const initialState: State = {
  isLoading: false,
  isSubmitting: false,
  currentUser: undefined,
  errors: null,
};

export const reducer = createReducer(
  initialState,

  // register
  on(AuthActions.register, (state): State => {
    return { ...state, isSubmitting: true };
  }),
  on(AuthActions.registerSuccess, (state, { currentUser }): State => {
    return { ...state, isSubmitting: false, currentUser };
  }),
  on(AuthActions.registerFailure, (state, { errors }): State => {
    return { ...state, isSubmitting: false, errors };
  }),

  // login
  on(AuthActions.login, (state): State => {
    return { ...state, isSubmitting: true };
  }),
  on(AuthActions.loginSuccess, (state, { currentUser }): State => {
    return { ...state, isSubmitting: false, currentUser };
  }),
  on(AuthActions.loginFailure, (state, { errors }): State => {
    return { ...state, isSubmitting: false, errors };
  }),

  // reset errors on every route change
  on(routerNavigatedAction, (state): State => {
    return { ...state, errors: null };
  }),

  // get current user
  on(AuthActions.getCurrentUser, (state): State => {
    return { ...state, isLoading: true };
  }),
  on(AuthActions.getCurrentUserSuccess, (state, { currentUser }): State => {
    return { ...state, isLoading: false, currentUser };
  }),
  on(AuthActions.getCurrentUserFailure, (state): State => {
    return { ...state, isLoading: false, currentUser: null };
  }),

  // update current user
  on(AuthActions.updateCurrentUserSuccess, (state, { currentUser }): State => {
    return { ...state, currentUser };
  }),
  on(AuthActions.updateCurrentUserFailure, (state, { errors }): State => {
    return { ...state, errors };
  }),

  // on logout
  on(AuthActions.logout, (state): State => {
    return { ...state, ...initialState, currentUser: null };
  })
);

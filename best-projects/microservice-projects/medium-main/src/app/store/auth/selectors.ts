import { CurrentUser } from '@auth/models/current-user.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { createFeatureSelector, createSelector } from '@ngrx/store';
import { State as AuthState, FeatureKey } from '@store/auth';

const selectAuthState = createFeatureSelector<AuthState>(FeatureKey);

export const isLoading = createSelector(selectAuthState, ({ isLoading }: AuthState): boolean => isLoading);
export const isSubmitting = createSelector(selectAuthState, ({ isSubmitting }: AuthState): boolean => isSubmitting);
export const currentUser = createSelector(selectAuthState, ({ currentUser }: AuthState): CurrentUser | null | undefined => currentUser);
export const errors = createSelector(selectAuthState, ({ errors }: AuthState): BackendErrors | null => errors);

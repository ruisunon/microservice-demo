import { CurrentUserRequest } from '@auth/models/current-user-request.model';
import { CurrentUser } from '@auth/models/current-user.model';
import { LoginRequest } from '@auth/models/login-request.model';
import { RegisterRequest } from '@auth/models/register-request.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { createAction, props } from '@ngrx/store';

export const register = createAction('[Auth] Register', props<{ request: RegisterRequest }>());
export const registerSuccess = createAction('[Auth] Register success', props<{ currentUser: CurrentUser }>());
export const registerFailure = createAction('[Auth] Register failure', props<{ errors: BackendErrors }>());

export const login = createAction('[Auth] Login', props<{ request: LoginRequest }>());
export const loginSuccess = createAction('[Auth] Login success', props<{ currentUser: CurrentUser }>());
export const loginFailure = createAction('[Auth] Login failure', props<{ errors: BackendErrors }>());

export const getCurrentUser = createAction('[Auth] Get current user');
export const getCurrentUserSuccess = createAction('[Auth] Get current user success', props<{ currentUser: CurrentUser }>());
export const getCurrentUserFailure = createAction('[Auth] Get current user failure');

export const updateCurrentUser = createAction('[Auth] Update current user', props<{ currentUserRequest: CurrentUserRequest }>());
export const updateCurrentUserSuccess = createAction('[Auth] Update current user success', props<{ currentUser: CurrentUser }>());
export const updateCurrentUserFailure = createAction('[Auth] Update current user failure', props<{ errors: BackendErrors }>());

export const logout = createAction('[Auth] Logout');

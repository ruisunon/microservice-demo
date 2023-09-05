import { BackendErrors } from '@core/models/backend-errors.model';
import { createAction, props } from '@ngrx/store';
import { UserProfile } from '@user-profile/models/user-profile.model';

export const getUserProfile = createAction('[User profile] Get user profile', props<{ slug: string }>());
export const getUserProfileSuccess = createAction('[User profile] Get user profile success', props<{ userProfile: UserProfile }>());
export const getUserProfileFailure = createAction('[User profile] Get user profile failure');

import { createFeatureSelector, createSelector } from '@ngrx/store';
import { FeatureKey, State as UserProfileState } from '@store/user-profile';
import { UserProfile } from '@user-profile/models/user-profile.model';

const selectUserProfileState = createFeatureSelector<UserProfileState>(FeatureKey);

export const isLoading = createSelector(selectUserProfileState, ({ isLoading }: UserProfileState): boolean => isLoading);
export const userProfile = createSelector(selectUserProfileState, ({ userProfile }: UserProfileState): UserProfile | null => userProfile);
export const error = createSelector(selectUserProfileState, ({ error }: UserProfileState): string | null => error);

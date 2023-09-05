import { BackendErrors } from '@core/models/backend-errors.model';
import { createFeatureSelector, createSelector } from '@ngrx/store';
import { State as SettingsState, FeatureKey } from '@store/settings';

const selectSettingsState = createFeatureSelector<SettingsState>(FeatureKey);

export const errors = createSelector(selectSettingsState, ({ errors }: SettingsState): BackendErrors | null => errors);
export const isSubmitting = createSelector(selectSettingsState, ({ isSubmitting }: SettingsState): boolean => isSubmitting);

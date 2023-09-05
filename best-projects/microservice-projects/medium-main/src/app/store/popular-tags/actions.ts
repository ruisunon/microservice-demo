import { createAction, props } from '@ngrx/store';

export const getPopularTags = createAction('[Tags] Get popular tags');
export const getPopularTagsSuccess = createAction('[Tags] Get popular tags success', props<{ popularTags: string[] }>());
export const getPopularTagsFailure = createAction('[Tags] Get popular tags failure');

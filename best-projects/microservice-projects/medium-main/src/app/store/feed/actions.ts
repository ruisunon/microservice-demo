import { GetFeedResponse } from '@core/models/get-feed-response.model';
import { createAction, props } from '@ngrx/store';

export const getFeed = createAction('[Feed] Get feed', props<{ url: string }>());
export const getFeedSuccess = createAction('[Feed] Get feed success', props<{ feed: GetFeedResponse }>());
export const getFeedFailure = createAction('[Feed] Get feed failure');

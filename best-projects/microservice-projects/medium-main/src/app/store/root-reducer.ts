import { InjectionToken } from '@angular/core';
import { RouterState, routerReducer } from '@ngrx/router-store';
import { Action, ActionReducerMap } from '@ngrx/store';
import * as fromArticle from '@store/article';
import * as fromAuth from '@store/auth';
import * as fromFeed from '@store/feed';
import * as fromPopularTags from '@store/popular-tags';
import * as fromSettings from '@store/settings';
import * as fromUserProfile from '@store/user-profile';

export interface AppState {
  router: RouterState;
  [fromAuth.FeatureKey]: fromAuth.State;
  [fromFeed.FeatureKey]: fromFeed.State;
  [fromPopularTags.FeatureKey]: fromPopularTags.State;
  [fromArticle.FeatureKey]: fromArticle.State;
  [fromSettings.FeatureKey]: fromSettings.State;
  [fromUserProfile.FeatureKey]: fromUserProfile.State;
}

export const ROOT_REDUCER_TOKEN = 'Root reducers';

export const ROOT_REDUCER = new InjectionToken<ActionReducerMap<AppState>>(ROOT_REDUCER_TOKEN, {
  factory: (): ActionReducerMap<AppState, Action> => ({
    router: routerReducer,
    [fromAuth.FeatureKey]: fromAuth.reducer,
    [fromFeed.FeatureKey]: fromFeed.reducer,
    [fromPopularTags.FeatureKey]: fromPopularTags.reducer,
    [fromArticle.FeatureKey]: fromArticle.reducer,
    [fromSettings.FeatureKey]: fromSettings.reducer,
    [fromUserProfile.FeatureKey]: fromUserProfile.reducer,
  }),
});

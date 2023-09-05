import { Article } from '@core/models/article.model';
import { createAction, props } from '@ngrx/store';

export const addToFavorites = createAction('[Favorites] Add to favorites', props<{ isFavorited: boolean; slug: string }>());
export const addToFavoritesSuccess = createAction('[Favorites] Add to favorites success', props<{ article: Article }>());
export const addToFavoritesFailure = createAction('[Favorites] Add to favorites failure');

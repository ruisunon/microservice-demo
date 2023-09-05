import { inject } from '@angular/core';
import { Article } from '@core/models/article.model';
import { AddToFavoritesService } from '@feed/services/add-to-favorites.service';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { AddToFavoritesActions } from '@store/add-to-favorites';
import { Observable, exhaustMap, map, catchError, of } from 'rxjs';

export const addToFavoritesEffect = createEffect(
  (actions$ = inject(Actions), addToFavoritesService = inject(AddToFavoritesService)) => {
    return actions$.pipe(
      ofType(AddToFavoritesActions.addToFavorites),

      exhaustMap(({ isFavorited, slug }) => {
        const article$: Observable<Article> = isFavorited
          ? addToFavoritesService.removeFromFavorites$(slug)
          : addToFavoritesService.addToFavorites$(slug);

        return article$.pipe(
          map((article: Article) => {
            return AddToFavoritesActions.addToFavoritesSuccess({ article });
          }),
          catchError(() => {
            return of(AddToFavoritesActions.addToFavoritesFailure());
          })
        );
      })
    );
  },
  { functional: true }
);

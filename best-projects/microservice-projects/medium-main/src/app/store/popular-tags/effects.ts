import { inject } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { PopularTagsService } from '@shared/components/popular-tags/services/popular-tags.service';
import { PopularTagsActions } from '@store/popular-tags';
import { catchError, exhaustMap, map, of } from 'rxjs';

export const getPopularTags = createEffect(
  (actions$ = inject(Actions), popularTagsService = inject(PopularTagsService)) => {
    return actions$.pipe(
      ofType(PopularTagsActions.getPopularTags),
      exhaustMap(() => {
        return popularTagsService.getPopularTags$().pipe(
          map((popularTags) => {
            return PopularTagsActions.getPopularTagsSuccess({ popularTags });
          }),
          catchError(() => {
            return of(PopularTagsActions.getPopularTagsFailure());
          })
        );
      })
    );
  },
  { functional: true }
);

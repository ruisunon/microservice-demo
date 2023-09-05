import { inject } from '@angular/core';
import { GetFeedResponse } from '@core/models/get-feed-response.model';
import { FeedService } from '@feed/services/feed.service';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { FeedActions } from '@store/feed';
import { catchError, exhaustMap, map, of } from 'rxjs';

export const getFeed = createEffect(
  (actions$ = inject(Actions), feedService = inject(FeedService)) => {
    return actions$.pipe(
      ofType(FeedActions.getFeed),
      exhaustMap(({ url }) => {
        return feedService.getFeed$(url).pipe(
          map((feed: GetFeedResponse) => {
            return FeedActions.getFeedSuccess({ feed });
          }),
          catchError(() => {
            return of(FeedActions.getFeedFailure());
          })
        );
      })
    );
  },
  { functional: true }
);

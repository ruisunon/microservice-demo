import { inject } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { UserProfileActions } from '@store/user-profile';
import { UserProfile } from '@user-profile/models/user-profile.model';
import { UserProfileService } from '@user-profile/services/user-profile.service';
import { exhaustMap, map, catchError, of } from 'rxjs';

export const getUserProfile = createEffect(
  (actions$ = inject(Actions), userProfileService = inject(UserProfileService)) => {
    return actions$.pipe(
      ofType(UserProfileActions.getUserProfile),
      exhaustMap(({ slug }) => {
        return userProfileService.getUserProfile$(slug).pipe(
          map((userProfile: UserProfile) => {
            return UserProfileActions.getUserProfileSuccess({ userProfile });
          }),
          catchError(() => {
            return of(UserProfileActions.getUserProfileFailure());
          })
        );
      })
    );
  },
  { functional: true }
);

import { HttpErrorResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { CurrentUser } from '@auth/models/current-user.model';
import { LoginRequest } from '@auth/models/login-request.model';
import { RegisterRequest } from '@auth/models/register-request.model';
import { AuthService } from '@auth/services/auth.service';
import { AccessToken } from '@core/constants/access-token';
import { BackendErrors } from '@core/models/backend-errors.model';
import { PersistanceService } from '@core/services/persistance.service';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { AuthActions } from '@store/auth';
import { catchError, exhaustMap, map, of, tap } from 'rxjs';

export const registerEffect = createEffect(
  (actions$ = inject(Actions), authService = inject(AuthService), persistanceService = inject(PersistanceService)) => {
    return actions$.pipe(
      ofType(AuthActions.register),
      exhaustMap(({ request }) => {
        return authService.register$(<RegisterRequest>request).pipe(
          map((currentUser: CurrentUser) => {
            persistanceService.set(AccessToken, currentUser.token);

            return AuthActions.registerSuccess({ currentUser });
          }),
          catchError((err: HttpErrorResponse) => {
            const errors = err.error.errors as BackendErrors;
            return of(AuthActions.registerFailure({ errors }));
          })
        );
      })
    );
  },
  { functional: true }
);

export const redirectAfterRegisterSuccess = createEffect(
  (actions$ = inject(Actions), router = inject(Router)) => {
    return actions$.pipe(
      ofType(AuthActions.registerSuccess),
      tap((): void => {
        router.navigateByUrl('/');
      })
    );
  },
  { functional: true, dispatch: false }
);

export const loginEffect = createEffect(
  (actions$ = inject(Actions), authService = inject(AuthService), persistanceService = inject(PersistanceService)) => {
    return actions$.pipe(
      ofType(AuthActions.login),
      exhaustMap(({ request }) => {
        return authService.login$(<LoginRequest>request).pipe(
          map((currentUser: CurrentUser) => {
            persistanceService.set(AccessToken, currentUser.token);
            return AuthActions.loginSuccess({ currentUser });
          }),
          catchError((err: HttpErrorResponse) => {
            const errors = err.error.errors as BackendErrors;
            return of(AuthActions.loginFailure({ errors }));
          })
        );
      })
    );
  },
  { functional: true }
);

export const redirectAfterLoginSuccess = createEffect(
  (actions$ = inject(Actions), router = inject(Router)) => {
    return actions$.pipe(
      ofType(AuthActions.loginSuccess),
      tap((): void => {
        router.navigateByUrl('/');
      })
    );
  },
  { functional: true, dispatch: false }
);

export const getCurrentUser = createEffect(
  (actions$ = inject(Actions), authService = inject(AuthService), persistanceService = inject(PersistanceService)) => {
    return actions$.pipe(
      ofType(AuthActions.getCurrentUser),
      exhaustMap(() => {
        const token = persistanceService.get(AccessToken);
        if (!token) return of(AuthActions.getCurrentUserFailure());

        return authService.getCurrentUser$().pipe(
          map((currentUser: CurrentUser) => {
            return AuthActions.getCurrentUserSuccess({ currentUser });
          }),
          catchError(() => {
            return of(AuthActions.getCurrentUserFailure());
          })
        );
      })
    );
  },
  { functional: true }
);

export const updateCurrentUser = createEffect(
  (actions$ = inject(Actions), authService = inject(AuthService)) => {
    return actions$.pipe(
      ofType(AuthActions.updateCurrentUser),
      exhaustMap(({ currentUserRequest }) => {
        return authService.updateCurrentUser$(currentUserRequest).pipe(
          map((currentUser: CurrentUser) => {
            return AuthActions.updateCurrentUserSuccess({ currentUser });
          }),
          catchError((err: HttpErrorResponse) => {
            const errors = err.error.errors as BackendErrors;
            return of(AuthActions.updateCurrentUserFailure({ errors }));
          })
        );
      })
    );
  },
  { functional: true }
);

export const logout = createEffect(
  (actions$ = inject(Actions), router = inject(Router), persistanceService = inject(PersistanceService)) => {
    return actions$.pipe(
      ofType(AuthActions.logout),
      tap((): void => {
        persistanceService.set(AccessToken, '');
        router.navigateByUrl('/');
      })
    );
  },
  { functional: true, dispatch: false }
);

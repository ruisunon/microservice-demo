import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { CurrentUser } from '@auth/models/current-user.model';
import { Store } from '@ngrx/store';
import { AuthSelectors } from '@store/auth';
import { Observable, map, tap } from 'rxjs';

export const authGuard = (): Observable<boolean> => {
  const router: Router = inject(Router);

  return inject(Store)
    .select(AuthSelectors.currentUser)
    .pipe(
      map((currentUser: CurrentUser | null | undefined): boolean => !!currentUser),
      tap((isCurrentUser: boolean): void => {
        !isCurrentUser && router.navigateByUrl('/');
      })
    );
};

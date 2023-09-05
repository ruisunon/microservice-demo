import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { CurrentUser } from '@auth/models/current-user.model';
import { Store } from '@ngrx/store';
import { AuthSelectors } from '@store/auth';
import { UserProfileSelectors } from '@store/user-profile';
import { GetUserProfileResponse } from '@user-profile/models/get-user-profile-response.model';
import { UserProfileDataSet } from '@user-profile/models/user-profile-data-set.model';
import { UserProfile } from '@user-profile/models/user-profile.model';
import { Observable, combineLatest, filter, map } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({ providedIn: 'root' })
export class UserProfileService {
  private readonly http: HttpClient = inject(HttpClient);
  private readonly store: Store = inject(Store);
  private readonly baseUrl: string = environment.baseApiUrl;

  public getUserProfile$(slug: string): Observable<UserProfile> {
    return this.http.get<GetUserProfileResponse>(`${this.baseUrl}/profiles/${slug}`).pipe(map(({ profile }): UserProfile => profile));
  }

  public getUserProfileDataSet$(): Observable<UserProfileDataSet> {
    return combineLatest({
      userProfile: this.store.select(UserProfileSelectors.userProfile),
      isLoading: this.store.select(UserProfileSelectors.isLoading),
      error: this.store.select(UserProfileSelectors.error),
      isCurrentUserProfile: this.getIsCurrentUserProfile$(),
    });
  }

  private getIsCurrentUserProfile$(): Observable<boolean> {
    return combineLatest({
      currentUser: this.store
        .select(AuthSelectors.currentUser)
        .pipe(filter((currentUser: CurrentUser | null | undefined): currentUser is CurrentUser => Boolean(currentUser))),
      userProfile: this.store
        .select(UserProfileSelectors.userProfile)
        .pipe(filter((userProfile: UserProfile | null): userProfile is UserProfile => Boolean(userProfile))),
    }).pipe(
      map(({ currentUser, userProfile }: { currentUser: CurrentUser; userProfile: UserProfile }) => {
        return currentUser.username === userProfile.username;
      })
    );
  }
}

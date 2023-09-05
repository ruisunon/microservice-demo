import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, inject } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { DestroyComponent } from '@core/abstracts/destroy/destroy.component';
import { FeedComponent } from '@feed/components/feed/feed.component';
import { Store } from '@ngrx/store';
import { UserProfileActions } from '@store/user-profile';
import { UserProfileDataSet } from '@user-profile/models/user-profile-data-set.model';
import { UserProfileService } from '@user-profile/services/user-profile.service';
import { Observable, takeUntil } from 'rxjs';

const UserProfileImports: Array<any> = [CommonModule, FeedComponent];

@Component({
  selector: 'app-user-profile',
  standalone: true,
  imports: UserProfileImports,
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class UserProfileComponent extends DestroyComponent implements OnInit {
  private readonly store: Store = inject(Store);
  private readonly router: Router = inject(Router);
  private readonly route: ActivatedRoute = inject(ActivatedRoute);
  private readonly userProfileService: UserProfileService = inject(UserProfileService);

  private slug: string = '';

  public userProfileDataSet$: Observable<UserProfileDataSet> = this.userProfileService.getUserProfileDataSet$();

  public ngOnInit(): void {
    this.route.params.pipe(takeUntil(this.destroy$)).subscribe({
      next: (params: Params): void => {
        this.slug = params['slug'];
        this.loadUserProfile(this.slug);
      },
    });
  }

  private loadUserProfile(slug: string): void {
    this.store.dispatch(UserProfileActions.getUserProfile({ slug }));
  }

  public get apiUrl(): string {
    const isFavorites: boolean = this.router.url.includes('favorites');
    return isFavorites ? `articles?favorited=${this.slug}` : `articles?author=${this.slug}`;
  }
}

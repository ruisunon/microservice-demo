import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, inject } from '@angular/core';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { ErrorMessageComponent } from '@shared/components/error-message/error-message.component';
import { PopularTagsActions, PopularTagsSelectors } from '@store/popular-tags';
import { Observable } from 'rxjs';

const PopularTagsImports: Array<any> = [CommonModule, ErrorMessageComponent, MatProgressSpinnerModule, RouterLink];

@Component({
  selector: 'app-popular-tags',
  standalone: true,
  imports: PopularTagsImports,
  templateUrl: './popular-tags.component.html',
  styleUrls: ['./popular-tags.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PopularTagsComponent implements OnInit {
  private readonly store: Store = inject(Store);

  public readonly error$: Observable<string | null> = this.store.select(PopularTagsSelectors.error);
  public readonly isLoading$: Observable<boolean> = this.store.select(PopularTagsSelectors.isLoading);
  public readonly popularTags$: Observable<string[] | null> = this.store.select(PopularTagsSelectors.popularTags);

  public ngOnInit(): void {
    this.store.dispatch(PopularTagsActions.getPopularTags());
  }
}

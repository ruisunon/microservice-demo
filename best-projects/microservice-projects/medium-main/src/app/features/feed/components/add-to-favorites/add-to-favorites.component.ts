import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input, Self, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { AddToFavoritesService } from '@feed/services/add-to-favorites.service';
import { Store } from '@ngrx/store';
import { AddToFavoritesActions } from '@store/add-to-favorites';

const AddToFavoritesImports: Array<any> = [CommonModule, MatButtonModule, MatIconModule];
const AddToFavoritesProviders: Array<any> = [AddToFavoritesService];

@Component({
  selector: 'app-add-to-favorites',
  standalone: true,
  imports: AddToFavoritesImports,
  providers: AddToFavoritesProviders,
  template: `
    <button (click)="handleLike()" [color]="isFavorited ? 'accent' : 'primary'" mat-icon-button>
      <mat-icon>favorite</mat-icon>
      <span>{{ favoritesCount }}</span>
    </button>
  `,
  styleUrls: ['./add-to-favorites.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AddToFavoritesComponent {
  @Input() public isFavorited: boolean = false;
  @Input() public favoritesCount: number = 0;
  @Input() public articleSlug: string = '';

  private readonly store: Store = inject(Store);

  public handleLike(): void {
    this.store.dispatch(AddToFavoritesActions.addToFavorites({ isFavorited: this.isFavorited, slug: this.articleSlug }));

    if (this.isFavorited) {
      this.favoritesCount = this.favoritesCount - 1;
    } else this.favoritesCount + 1;

    this.isFavorited = !this.isFavorited;
  }
}

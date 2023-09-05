import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input, inject } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { CurrentUser } from '@auth/models/current-user.model';
import { Store } from '@ngrx/store';
import { AuthSelectors } from '@store/auth';
import { Observable } from 'rxjs';

const FeedTogglerImports: Array<any> = [CommonModule, RouterLink, RouterLinkActive];

@Component({
  selector: 'app-feed-toggler',
  standalone: true,
  imports: FeedTogglerImports,
  templateUrl: './feed-toggler.component.html',
  styleUrls: ['./feed-toggler.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class FeedTogglerComponent {
  @Input() public tagName?: string;

  public currentUser$: Observable<CurrentUser | null | undefined> = inject(Store).select(AuthSelectors.currentUser);
}

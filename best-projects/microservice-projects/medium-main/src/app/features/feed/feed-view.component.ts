import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

const FeedContainerImports: Array<any> = [RouterOutlet];

@Component({
  selector: 'app-feed-view',
  standalone: true,
  imports: FeedContainerImports,
  template: `
    <section>
      <router-outlet />
    </section>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class FeedViewComponent {}

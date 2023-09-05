import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { FeedContainerComponent } from '@feed/components/feed-container/feed-container.component';

const YourFeedImports: Array<any> = [CommonModule, FeedContainerComponent];

@Component({
  selector: 'app-your-feed',
  standalone: true,
  imports: YourFeedImports,
  template: `<app-feed-container [apiUrl]="apiUrl" />`,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class YourFeedComponent {
  public readonly apiUrl: string = 'articles';
}

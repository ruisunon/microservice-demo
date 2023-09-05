import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { FeedContainerComponent } from '@feed/components/feed-container/feed-container.component';

const GlobalFeedImports: Array<any> = [CommonModule, FeedContainerComponent];

@Component({
  selector: 'app-global-feed',
  standalone: true,
  imports: GlobalFeedImports,
  template: `<app-feed-container [apiUrl]="apiUrl" />`,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class GlobalFeedComponent {
  public apiUrl: string = 'articles';
}

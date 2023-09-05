import { ChangeDetectionStrategy, Component, Input } from '@angular/core';
import { FeedComponent } from '@feed/components/feed/feed.component';
import { FeedTogglerComponent } from '@feed/components/feed-toggler/feed-toggler.component';
import { PopularTagsComponent } from '@shared/components/popular-tags/popular-tags.component';
import { BannerComponent } from '@ui/banner/banner.component';

const FeedContainerImports: Array<any> = [BannerComponent, FeedTogglerComponent, PopularTagsComponent, FeedComponent];

@Component({
  selector: 'app-feed-container',
  standalone: true,
  imports: FeedContainerImports,
  templateUrl: './feed-container.component.html',
  styleUrls: ['./feed-container.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class FeedContainerComponent {
  @Input({ required: true }) public apiUrl: string = '';
  @Input() public tagName: string = '';
}

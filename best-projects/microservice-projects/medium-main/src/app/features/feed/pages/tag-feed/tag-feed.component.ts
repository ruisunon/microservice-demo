import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, OnInit, inject } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { DestroyComponent } from '@core/abstracts/destroy/destroy.component';
import { FeedContainerComponent } from '@feed/components/feed-container/feed-container.component';
import { takeUntil } from 'rxjs';

const TagFeedImports: Array<any> = [CommonModule, FeedContainerComponent];

@Component({
  selector: 'app-tag-feed',
  standalone: true,
  imports: TagFeedImports,
  template: `<app-feed-container [apiUrl]="apiUrl" [tagName]="tagName" />`,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class TagFeedComponent extends DestroyComponent implements OnInit {
  private readonly route: ActivatedRoute = inject(ActivatedRoute);

  public apiUrl: string = '';
  public tagName: string = '';

  public ngOnInit(): void {
    this.route.params.pipe(takeUntil(this.destroy$)).subscribe({
      next: (params: Params): void => {
        this.tagName = params['slug'];
        this.apiUrl = `articles?tag=${this.tagName}`;
      },
    });
  }
}

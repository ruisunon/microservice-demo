import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

const ArticleContainerImports: Array<any> = [RouterOutlet];

@Component({
  selector: 'app-article-view',
  standalone: true,
  imports: ArticleContainerImports,
  template: `
    <section class="container-max-w-sm">
      <router-outlet />
    </section>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class ArticleViewComponent {}

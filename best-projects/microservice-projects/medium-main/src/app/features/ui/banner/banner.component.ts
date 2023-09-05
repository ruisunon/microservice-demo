import { Component } from '@angular/core';

@Component({
  selector: 'app-banner',
  standalone: true,
  template: `
    <div class="banner">
      <div class="banner__inner">
        <h1 class="banner__title">Medium</h1>
        <div>A place to share knowledge</div>
      </div>
    </div>
  `,
  styleUrls: ['./banner.component.scss'],
})
export class BannerComponent {}

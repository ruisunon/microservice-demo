import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

const UserProfileViewImports: Array<any> = [CommonModule, RouterOutlet];

@Component({
  selector: 'app-user-profile-view',
  standalone: true,
  imports: UserProfileViewImports,
  template: `
    <section>
      <router-outlet />
    </section>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class UserProfileViewComponent {}

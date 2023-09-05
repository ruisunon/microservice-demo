import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

const AuthImports: Array<any> = [RouterOutlet];

@Component({
  selector: 'app-auth-view',
  standalone: true,
  imports: AuthImports,
  template: `
    <section>
      <router-outlet />
    </section>
  `,
})
export default class AuthViewComponent {}

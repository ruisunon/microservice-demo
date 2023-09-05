import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { AuthFormComponent } from '@auth/components/auth-form/auth-form.component';
import { AuthFormMode } from '@auth/enums/auth-form-mode.enum';
import { AuthFormPayload } from '@auth/models/auth-form-payload.model';
import { LoginForm } from '@auth/models/form.model';
import { LoginRequest } from '@auth/models/login-request.model';
import { AuthFormService } from '@auth/services/auth-form.service';
import { BackendErrors } from '@core/models/backend-errors.model';
import { Store } from '@ngrx/store';
import { AuthActions, AuthSelectors } from '@store/auth';
import { Observable } from 'rxjs';

const LoginImports: Array<any> = [CommonModule, AuthFormComponent, RouterModule];
const LoginProviders: Array<any> = [AuthFormService];
@Component({
  selector: 'app-login',
  standalone: true,
  imports: LoginImports,
  providers: LoginProviders,
  template: `
    <app-auth-form
      [form]="loginForm"
      [mode]="authFormMode.LOGIN"
      [errors]="loginErrors$ | async"
      [isSubmitting]="(isSubmitting$ | async)!"
      (formSubmit)="onFormSubmit($event)" />
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class LoginComponent {
  private store: Store = inject(Store);

  public isLoading$: Observable<boolean> = this.store.select(AuthSelectors.isLoading);
  public isSubmitting$: Observable<boolean> = this.store.select(AuthSelectors.isSubmitting);
  public loginErrors$: Observable<BackendErrors | null> = this.store.select(AuthSelectors.errors);

  public loginForm: FormGroup<LoginForm> = inject(AuthFormService).getLoginForm();
  public authFormMode = AuthFormMode;

  public onFormSubmit(user: AuthFormPayload): void {
    this.store.dispatch(AuthActions.login({ request: { user } as LoginRequest }));
  }
}

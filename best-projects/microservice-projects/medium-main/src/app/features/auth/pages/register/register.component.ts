import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { AuthFormComponent } from '@auth/components/auth-form/auth-form.component';
import { AuthFormMode } from '@auth/enums/auth-form-mode.enum';
import { AuthFormPayload } from '@auth/models/auth-form-payload.model';
import { RegisterForm } from '@auth/models/form.model';
import { RegisterRequest } from '@auth/models/register-request.model';
import { AuthFormService } from '@auth/services/auth-form.service';
import { BackendErrors } from '@core/models/backend-errors.model';
import { Store } from '@ngrx/store';
import { AuthActions, AuthSelectors } from '@store/auth';
import { Observable } from 'rxjs';

const RegisterImports: Array<any> = [CommonModule, AuthFormComponent, RouterModule];
const RegisterProviders: Array<any> = [AuthFormService];
@Component({
  selector: 'app-register',
  standalone: true,
  imports: RegisterImports,
  providers: RegisterProviders,
  template: `
    <app-auth-form
      [form]="registerForm"
      [mode]="authFormMode.REGISTER"
      [errors]="registerErrors$ | async"
      [isSubmitting]="(isSubmitting$ | async)!"
      (formSubmit)="onRegisterFormSubmit($event)" />
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class RegisterComponent {
  private readonly store: Store = inject(Store);

  public isLoading$: Observable<boolean> = this.store.select(AuthSelectors.isLoading);
  public isSubmitting$: Observable<boolean> = this.store.select(AuthSelectors.isSubmitting);
  public registerErrors$: Observable<BackendErrors | null> = this.store.select(AuthSelectors.errors);

  public readonly registerForm: FormGroup<RegisterForm> = inject(AuthFormService).getRegisterForm();
  public readonly authFormMode = AuthFormMode;

  public onRegisterFormSubmit(user: AuthFormPayload): void {
    this.store.dispatch(AuthActions.register({ request: { user } as RegisterRequest }));
  }
}

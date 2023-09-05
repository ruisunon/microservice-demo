import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { RouterLink } from '@angular/router';
import { ErrorMessagesComponent } from '@auth/components/error-messages/error-messages.component';
import { AuthFormMode } from '@auth/enums/auth-form-mode.enum';
import { AuthFormPayload } from '@auth/models/auth-form-payload.model';
import { LoginForm, RegisterForm } from '@auth/models/form.model';
import { AuthFormService } from '@auth/services/auth-form.service';
import { BackendErrors } from '@core/models/backend-errors.model';

const AuthFormImports: Array<any> = [
  CommonModule,
  ReactiveFormsModule,
  MatFormFieldModule,
  MatInputModule,
  MatButtonModule,
  ErrorMessagesComponent,
  RouterLink,
];
const AuthFormProviders: Array<any> = [AuthFormService];

@Component({
  selector: 'app-auth-form',
  standalone: true,
  imports: AuthFormImports,
  providers: AuthFormProviders,
  templateUrl: './auth-form.component.html',
  styleUrls: ['./auth-form.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AuthFormComponent {
  @Input() public form!: FormGroup<RegisterForm> | FormGroup<LoginForm>;
  @Input() public mode!: AuthFormMode;
  @Input() public isSubmitting: boolean = false;
  @Input() set errors(errors: BackendErrors | null) {
    this.errorMessages = this.setErrorMessages(errors);
  }

  @Output() public formSubmit: EventEmitter<AuthFormPayload> = new EventEmitter<AuthFormPayload>();

  get email(): FormControl<string> {
    return this.form.controls.email;
  }

  get password(): FormControl<string> {
    return this.form.controls.password;
  }

  public AuthFormMode = AuthFormMode;
  public errorMessages: string[] = [];

  public onSubmit(): void {
    if (this.form.invalid) {
      this.form.markAsDirty();
      return;
    }

    this.formSubmit.emit(this.form.getRawValue());
  }

  private setErrorMessages(errors: BackendErrors | null): string[] {
    if (errors === null) return [];

    return Object.keys(errors).map((name: string): string => {
      const messages: string = errors[name].join(' ');
      return `${name} ${messages}`;
    });
  }
}

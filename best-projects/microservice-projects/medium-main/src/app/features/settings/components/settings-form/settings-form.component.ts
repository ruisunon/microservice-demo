import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnInit, Output, Self, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { ErrorMessagesComponent } from '@auth/components/error-messages/error-messages.component';
import { CurrentUserRequest } from '@auth/models/current-user-request.model';
import { CurrentUser } from '@auth/models/current-user.model';
import { DestroyComponent } from '@core/abstracts/destroy/destroy.component';
import { BackendErrors } from '@core/models/backend-errors.model';
import { Store } from '@ngrx/store';
import { SettingsForm } from '@settings/models/settings-form.model';
import { SettingsFormService } from '@settings/services/settings-form.service';
import { AuthSelectors } from '@store/auth';
import { filter, takeUntil } from 'rxjs';

const SettingsFormImports: Array<any> = [
  CommonModule,
  ReactiveFormsModule,
  MatFormFieldModule,
  ErrorMessagesComponent,
  MatInputModule,
  MatButtonModule,
];
const SettingsFormProviders: Array<any> = [SettingsFormService];

@Component({
  selector: 'app-settings-form',
  standalone: true,
  imports: SettingsFormImports,
  providers: SettingsFormProviders,
  templateUrl: './settings-form.component.html',
  styleUrls: ['./settings-form.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SettingsFormComponent extends DestroyComponent implements OnInit {
  @Input() public errorMessages: BackendErrors | null = null;

  @Output() public formSubmit: EventEmitter<CurrentUserRequest> = new EventEmitter<CurrentUserRequest>();

  private readonly store: Store = inject(Store);

  private currentUser!: CurrentUser;
  public readonly form: FormGroup<SettingsForm> = this.settingsFormService.getSettingsForm();

  constructor(@Self() private readonly settingsFormService: SettingsFormService) {
    super();
  }

  public ngOnInit(): void {
    this.store
      .select(AuthSelectors.currentUser)
      .pipe(filter(Boolean), takeUntil(this.destroy$))
      .subscribe({
        next: (currentUser: CurrentUser): void => {
          this.currentUser = currentUser;

          this.form.patchValue({
            image: this.currentUser.image ?? '',
            username: this.currentUser.username,
            bio: this.currentUser.bio ?? '',
            email: this.currentUser.email,
            password: '',
          });
        },
      });
  }

  public onSubmit(): void {
    if (this.form.invalid) {
      this.form.markAsDirty();
      return;
    }

    this.formSubmit.emit({ user: { ...this.currentUser, ...this.form.getRawValue() } });
  }

  public get username() {
    return this.form.get('username') as FormControl<string>;
  }

  public get email() {
    return this.form.get('email') as FormControl<string>;
  }

  public get password() {
    return this.form.get('password') as FormControl<string>;
  }
}

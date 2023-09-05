import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { CurrentUserRequest } from '@auth/models/current-user-request.model';
import { Store } from '@ngrx/store';
import { SettingsFormComponent } from '@settings/components/settings-form/settings-form.component';
import { AuthActions } from '@store/auth';

const SettingsImports: Array<any> = [CommonModule, SettingsFormComponent];

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: SettingsImports,
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class SettingsComponent {
  private readonly store: Store = inject(Store);

  public onSettingsFormSubmit(currentUserRequest: CurrentUserRequest): void {
    this.store.dispatch(AuthActions.updateCurrentUser({ currentUserRequest }));
  }
}

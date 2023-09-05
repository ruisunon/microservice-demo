import { Injectable, inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SettingsForm } from '@settings/models/settings-form.model';

@Injectable()
export class SettingsFormService {
  private readonly formBuilder: FormBuilder = inject(FormBuilder);

  public getSettingsForm(): FormGroup<SettingsForm> {
    return this.formBuilder.nonNullable.group({
      image: '',
      username: ['', [Validators.required]],
      bio: '',
      email: ['', [Validators.email, Validators.required]],
      password: ['', [Validators.minLength(8), Validators.required]],
    });
  }
}

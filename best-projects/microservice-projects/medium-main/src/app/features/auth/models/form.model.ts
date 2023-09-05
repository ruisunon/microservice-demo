import { FormControl } from '@angular/forms';

export type RegisterForm = {
  username: FormControl<string>;
  email: FormControl<string>;
  password: FormControl<string>;
};

export type LoginForm = Omit<RegisterForm, 'username'>;

import { User } from '@auth/models/user.model';

export interface LoginRequest {
  user: Pick<User, 'email' | 'password'>;
}

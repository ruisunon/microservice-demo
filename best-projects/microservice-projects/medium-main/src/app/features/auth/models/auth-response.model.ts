import { CurrentUser } from '@auth/models/current-user.model';

export interface AuthResponse {
  user: CurrentUser;
}

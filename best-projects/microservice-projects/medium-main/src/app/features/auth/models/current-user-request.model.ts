import { CurrentUser } from '@auth/models/current-user.model';

export interface CurrentUserRequest {
  user: CurrentUser & { password: string };
}

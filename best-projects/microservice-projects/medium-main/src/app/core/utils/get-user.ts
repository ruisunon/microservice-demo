import { AuthResponse } from '@auth/models/auth-response.model';
import { CurrentUser } from '@auth/models/current-user.model';

export const getUser = ({ user }: AuthResponse): CurrentUser => user;

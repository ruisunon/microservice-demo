import { UserProfile } from '@user-profile/models/user-profile.model';

export interface UserProfileDataSet {
  isLoading: boolean;
  error: string | null;
  userProfile: UserProfile | null;
  isCurrentUserProfile: boolean;
}

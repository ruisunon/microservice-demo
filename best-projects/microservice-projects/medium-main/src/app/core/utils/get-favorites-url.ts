import { environment } from 'src/environments/environment.development';

export const getFavoritesUrl = (slug: string): string => {
  return `${environment.baseApiUrl}/articles/${slug}/favorite`;
};

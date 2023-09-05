import { Routes } from '@angular/router';
import { RouteFragment } from '@core/enums/route-fragment.enum';
import { Route } from '@core/enums/route.enum';
import { authGuard } from '@core/guards/auth.guard';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('@feed/feed-view.component'),
    children: [
      {
        path: '',
        loadComponent: () => import('@feed/pages/your-feed/your-feed.component'),
      },
      {
        path: Route.FEED,
        loadComponent: () => import('@feed/pages/your-feed/your-feed.component'),
      },
      {
        path: `${Route.TAGS}/${RouteFragment.SLUG}`,
        loadComponent: () => import('@feed/pages/tag-feed/tag-feed.component'),
      },
    ],
  },
  {
    path: Route.AUTHENTICATION,
    loadComponent: () => import('@auth/auth-view.component'),
    children: [
      {
        path: RouteFragment.REGISTER,
        loadComponent: () => import('@auth/pages/register/register.component'),
      },
      {
        path: RouteFragment.LOGIN,
        loadComponent: () => import('@auth/pages/login/login.component'),
      },
    ],
  },
  {
    path: Route.ARTICLES,
    loadComponent: () => import('@article/article-view.component'),
    children: [
      {
        path: RouteFragment.NEW,
        loadComponent: () => import('@article/pages/create-article/create-article.component'),
      },
      {
        path: RouteFragment.SLUG,
        loadComponent: () => import('@article/pages/article/article.component'),
      },
    ],
  },
  {
    path: Route.SETTINGS,
    loadComponent: () => import('@settings/settings-view.component'),
    children: [
      {
        path: '',
        canActivate: [authGuard],
        loadComponent: () => import('@settings/pages/settings/settings.component'),
      },
    ],
  },
  {
    path: Route.PROFILES,
    loadComponent: () => import('@user-profile/user-profile-view.component'),
    children: [
      {
        path: RouteFragment.SLUG,
        loadComponent: () => import('@user-profile/pages/user-profile/user-profile.component'),
      },
      {
        path: `${RouteFragment.SLUG}/${RouteFragment.FAVORITES}`,
        loadComponent: () => import('@user-profile/pages/user-profile/user-profile.component'),
      },
    ],
  },
];

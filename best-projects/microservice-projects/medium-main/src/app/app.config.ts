import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { ApplicationConfig, isDevMode } from '@angular/core';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideRouter } from '@angular/router';
import { authInterceptor } from '@core/interceptors/auth.interceptor';
import { provideEffects } from '@ngrx/effects';
import { provideRouterStore } from '@ngrx/router-store';
import { provideStore } from '@ngrx/store';
import { provideStoreDevtools } from '@ngrx/store-devtools';
import * as addToFavoritesEffects from '@store/add-to-favorites/effects';
import * as articleEffects from '@store/article/effects';
import * as authEffects from '@store/auth/effects';
import * as feedEffects from '@store/feed/effects';
import * as popularTagsEffects from '@store/popular-tags/effects';
import { ROOT_REDUCER } from '@store/root-reducer';
import * as userProfileEffects from '@store/user-profile/effects';
import { routes } from 'src/app/app.routes';

const StoreEffects: Array<any> = [authEffects, feedEffects, popularTagsEffects, articleEffects, addToFavoritesEffects, userProfileEffects];
const Interceptors: Array<any> = [authInterceptor];

export const appConfig: ApplicationConfig = {
  providers: [
    provideAnimations(),
    provideRouter(routes),
    provideHttpClient(withInterceptors(Interceptors)),

    // NgRx
    provideRouterStore(),
    provideStore(ROOT_REDUCER),
    provideEffects(StoreEffects),
    provideStoreDevtools({
      maxAge: 25,
      logOnly: !isDevMode(),
      autoPause: true,
      trace: false,
      traceLimit: 75,
    }),
  ],
};

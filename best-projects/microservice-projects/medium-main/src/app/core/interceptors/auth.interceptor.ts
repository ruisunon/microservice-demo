import { HttpHandlerFn, HttpInterceptorFn, HttpRequest } from '@angular/common/http';
import { inject } from '@angular/core';
import { AccessToken } from '@core/constants/access-token';
import { PersistanceService } from '@core/services/persistance.service';

export const authInterceptor: HttpInterceptorFn = (request: HttpRequest<unknown>, next: HttpHandlerFn) => {
  const token = inject(PersistanceService).get(AccessToken);

  request = request.clone({
    setHeaders: { Authorization: token ? `Token ${token}` : '' },
  });

  return next(request);
};

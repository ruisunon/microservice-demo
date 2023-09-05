import { HttpErrorResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import * as articleMessage from '@article/constants/article-messages';
import { ArticleService } from '@article/services/article.service';
import { Route } from '@core/enums/route.enum';
import { Article } from '@core/models/article.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { SnackbarService } from '@shared/components/snackbar/services/snackbar.service';
import { ArticleActions } from '@store/article';
import { catchError, exhaustMap, map, of, tap } from 'rxjs';

export const getArticle = createEffect(
  (actions$ = inject(Actions), articleService = inject(ArticleService), snackbar = inject(SnackbarService)) => {
    return actions$.pipe(
      ofType(ArticleActions.getArticle),
      exhaustMap(({ slug }) =>
        articleService.getArticle$(slug).pipe(
          map((article: Article) => {
            return ArticleActions.getArticleSuccess({ article });
          }),
          catchError(() => {
            snackbar.openSnackbar({ message: articleMessage.GetArticlesError, type: 'error' });
            return of(ArticleActions.getArticleFailure());
          })
        )
      )
    );
  },
  { functional: true }
);

export const deleteArticle = createEffect(
  (actions$ = inject(Actions), articleService = inject(ArticleService), snackbar = inject(SnackbarService)) => {
    return actions$.pipe(
      ofType(ArticleActions.deleteArticle),
      exhaustMap(({ slug }) =>
        articleService.deleteArticle$(slug).pipe(
          map(() => {
            snackbar.openSnackbar({ message: articleMessage.DeleteArticleSuccess, type: 'success' });
            return ArticleActions.deleteArticleSuccess();
          }),
          catchError(() => {
            return of(ArticleActions.deleteArticleFailure());
          })
        )
      )
    );
  },
  { functional: true }
);

export const redirectAfterDeleteArticle = createEffect(
  (actions$ = inject(Actions), router = inject(Router)) => {
    return actions$.pipe(
      ofType(ArticleActions.deleteArticle),
      tap(() => {
        router.navigateByUrl('/');
      })
    );
  },
  { functional: true, dispatch: false }
);

export const createArticle = createEffect(
  (actions$ = inject(Actions), articleService = inject(ArticleService), snackbar = inject(SnackbarService)) => {
    return actions$.pipe(
      ofType(ArticleActions.createArticle),
      exhaustMap(({ articlePayload }) => {
        return articleService.createArticle$(articlePayload).pipe(
          map((article: Article) => {
            snackbar.openSnackbar({ message: articleMessage.CreateArticleSuccess, type: 'success' });
            return ArticleActions.createArticleSuccess({ article });
          }),
          catchError((err: HttpErrorResponse) => {
            const errors = err.error.errors as BackendErrors;

            return of(ArticleActions.createArticleFailure({ errors }));
          })
        );
      })
    );
  },
  { functional: true }
);

export const redirectAfterCreateArticle = createEffect(
  (actions$ = inject(Actions), router = inject(Router)) => {
    return actions$.pipe(
      ofType(ArticleActions.createArticleSuccess),
      tap(({ article }) => {
        router.navigate([Route.ARTICLES, article.slug]);
      })
    );
  },
  { functional: true, dispatch: false }
);

export const updateArticle = createEffect(
  (actions$ = inject(Actions), articleService = inject(ArticleService), snackbar = inject(SnackbarService)) => {
    return actions$.pipe(
      ofType(ArticleActions.updateArticle),
      exhaustMap(({ articlePayload, slug }) => {
        return articleService.updateArticle$(slug, articlePayload).pipe(
          map((article: Article) => {
            snackbar.openSnackbar({ message: articleMessage.UpdateArticleSuccess, type: 'success' });
            return ArticleActions.updateArticleSuccess({ article });
          }),
          catchError((err: HttpErrorResponse) => {
            const errors = err.error.errors as BackendErrors;

            return of(ArticleActions.updateArticleFailure({ errors }));
          })
        );
      })
    );
  },
  { functional: true }
);

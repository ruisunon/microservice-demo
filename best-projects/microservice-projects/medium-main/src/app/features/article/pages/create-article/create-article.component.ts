import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import ArticleFormComponent from '@article/components/article-form/article-form.component';
import { ArticleFormMode } from '@article/enums/article-form-mode.enum';
import { ArticlePayload } from '@article/models/article-payload.model';
import { BackendErrors } from '@core/models/backend-errors.model';
import { Store } from '@ngrx/store';
import { ArticleActions, ArticleSelectors } from '@store/article';
import { Observable } from 'rxjs';

const CreateArticleImports: Array<any> = [ArticleFormComponent, MatDialogModule];
const CreateArticleProviders: Array<any> = [
  { provide: MAT_DIALOG_DATA, useValue: {} },
  { provide: MatDialogRef, useValue: {} },
];

@Component({
  selector: 'app-create-article',
  standalone: true,
  imports: CreateArticleImports,
  providers: CreateArticleProviders,
  template: ` <app-article-form (formSubmit)="onFormSubmit($event)" [mode]="articleFormMode.CREATE" /> `,
  styleUrls: ['./create-article.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class CreateArticleComponent {
  private readonly store: Store = inject(Store);
  public readonly articleFormMode = ArticleFormMode;

  public readonly isSubmitting$: Observable<boolean> = this.store.select(ArticleSelectors.isCreateArticleSubmitting);
  public readonly validationErrors$: Observable<BackendErrors | null> = this.store.select(ArticleSelectors.createArticleErrors);

  public onFormSubmit(articlePayload: ArticlePayload): void {
    this.store.dispatch(ArticleActions.createArticle({ articlePayload }));
  }
}

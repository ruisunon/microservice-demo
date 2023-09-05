import { ArticleFormMode } from '@article/enums/article-form-mode.enum';
import { ArticlePayload } from '@article/models/article-payload.model';

export interface ArticleFormData {
  mode: ArticleFormMode;
  formValues: ArticlePayload;
}

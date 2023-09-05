import { FormControl } from '@angular/forms';

export interface ArticleForm {
  title: FormControl<string>;
  description: FormControl<string>;
  body: FormControl<string>;
  tagList: FormControl<string>;
}

import { Injectable, inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ArticleForm } from '@article/models/article-form.model';

@Injectable()
export class ArticleFormService {
  private readonly fb: FormBuilder = inject(FormBuilder);

  public getArticleForm(): FormGroup<ArticleForm> {
    return this.fb.nonNullable.group({
      title: ['', [Validators.required]],
      description: ['', [Validators.required]],
      body: ['', [Validators.required]],
      tagList: ['', [Validators.required]],
    });
  }
}

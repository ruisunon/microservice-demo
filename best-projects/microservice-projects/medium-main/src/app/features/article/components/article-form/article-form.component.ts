import { TextFieldModule } from '@angular/cdk/text-field';
import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, EventEmitter, Inject, Input, OnInit, Output, inject } from '@angular/core';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { ArticleFormMode } from '@article/enums/article-form-mode.enum';
import { ArticleFormData } from '@article/models/article-form-data.model';
import { ArticleForm } from '@article/models/article-form.model';
import { ArticlePayload } from '@article/models/article-payload.model';
import { ArticleFormService } from '@article/services/article-form.service';
import { BackendErrors } from '@core/models/backend-errors.model';

const ArticleFormImports: Array<any> = [
  CommonModule,
  ReactiveFormsModule,
  MatFormFieldModule,
  MatInputModule,
  MatIconModule,
  MatButtonModule,
  TextFieldModule,
];
const ArticleFormProviders: Array<any> = [ArticleFormService];

@Component({
  selector: 'app-article-form',
  standalone: true,
  imports: ArticleFormImports,
  providers: ArticleFormProviders,
  templateUrl: './article-form.component.html',
  styleUrls: ['./article-form.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export default class ArticleFormComponent implements OnInit {
  @Input() public mode: ArticleFormMode = ArticleFormMode.CREATE;
  @Input() public isSubmitting: boolean = false;
  @Input() public errors: BackendErrors | null = null;

  @Output() public formSubmit: EventEmitter<ArticlePayload> = new EventEmitter<ArticlePayload>();

  public form: FormGroup<ArticleForm> = inject(ArticleFormService).getArticleForm();

  constructor(@Inject(MAT_DIALOG_DATA) public formData: ArticleFormData, public readonly dialogRef: MatDialogRef<ArticleFormComponent>) {
    if (this.formData) this.mode = this.formData.mode;
  }

  public ngOnInit(): void {
    this.mode === ArticleFormMode.UPDATE && this.patchFormValues();
  }

  public onSubmit(): void {
    if (this.form.invalid) {
      this.form.markAsDirty();
      return;
    }

    const articleFormValues: ArticlePayload = { ...this.form.getRawValue(), tagList: this.form.value.tagList!.split(' ') };

    this.mode === ArticleFormMode.CREATE && this.formSubmit.emit(articleFormValues);
    this.mode === ArticleFormMode.UPDATE && this.dialogRef.close(articleFormValues);
  }

  private patchFormValues(): void {
    this.form.patchValue({
      title: this.formData.formValues.title,
      description: this.formData.formValues.description,
      body: this.formData.formValues.body,
      tagList: this.formData.formValues.tagList.join(' '),
    });
  }
}

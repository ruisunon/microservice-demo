import { ComponentFixture, TestBed } from '@angular/core/testing';
import ArticleFormComponent from '@article/components/article-form/article-form.component';

describe('ArticleFormComponent', () => {
  let component: ArticleFormComponent;
  let fixture: ComponentFixture<ArticleFormComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ArticleFormComponent],
    });
    fixture = TestBed.createComponent(ArticleFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
